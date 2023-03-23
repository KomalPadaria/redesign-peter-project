package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_tech_info/applications/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_tech_info/applications/repository"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/jira"
	jiraEntities "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/jira/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/onboarding"
	onboardingEntities "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/onboarding/entities"
)

type Service interface {
	// Application
	CreateApplication(ctx context.Context, companyUuid, userUuid *uuid.UUID, application *entities.TechInfoApplication) (*entities.TechInfoApplicationCreateUpdateResponse, error)
	UpdateApplication(ctx context.Context, companyUuid, userUuid, applicationUuid *uuid.UUID, application *entities.TechInfoApplication) (*entities.TechInfoApplicationCreateUpdateResponse, error)
	UpdateApplicationPatch(ctx context.Context, companyUuid, userUUID, applicationUuid *uuid.UUID, req *entities.UpdateApplicationPatchRequestBody) (*entities.UpdateApplicationPatchResponse, error)
	GetApplicationById(ctx context.Context, companyUuid, userUuid, applicationUuid *uuid.UUID) (*entities.TechInfoApplication, error)
	GetAllApplications(ctx context.Context, req *entities.GetAllApplicationsRequest) ([]*entities.TechInfoApplication, error)
	DeleteApplication(ctx context.Context, companyUuid, userUuid, applicationUuid *uuid.UUID) error
	// Environment
	CreateApplicationEnv(ctx context.Context, companyUuid, userUuid, applicationUuid *uuid.UUID, req *entities.ApplicationEnvCreateRequestBody) (*entities.ApplicationEnv, error)
	UpdateApplicationEnv(ctx context.Context, companyUuid, userUuid, applicationUuid, appEnvUuid *uuid.UUID, req *entities.ApplicationEnvUpdateRequestBody) (*entities.ApplicationEnv, error)
	UpdateApplicationEnvPatch(ctx context.Context, companyUuid, userUUID, applicationUuid, appEnvUuid *uuid.UUID, req *entities.UpdateApplicationEnvPatchRequestBody) (*entities.UpdateApplicationEnvPatchResponse, error)
	DeleteApplicationEnv(ctx context.Context, companyUuid, userUuid, applicationUuid, appEnvUuid *uuid.UUID) error
}

type service struct {
	repo             repository.Repository
	onboardingClient onboarding.Client
	jiraClient       jira.Client
}

func (s *service) DeleteApplicationEnv(ctx context.Context, companyUuid, userUuid, applicationUuid, appEnvUuid *uuid.UUID) error {
	appEnv, err := s.repo.GetApplicationEnvById(ctx, appEnvUuid)
	if err != nil {
		return err
	}
	_, err = s.repo.DeleteApplicationEnv(ctx, appEnvUuid)
	if err != nil {
		return err
	}

	comment := s.getApplicationEnvJiraComment(appEnv, "Deleted application environment: (Deleted application environment details are below)")

	err = s.jiraClient.AddComment(ctx, companyUuid, comment)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) UpdateApplicationEnvPatch(ctx context.Context, companyUuid, userUUID, applicationUuid, appEnvUuid *uuid.UUID, req *entities.UpdateApplicationEnvPatchRequestBody) (*entities.UpdateApplicationEnvPatchResponse, error) {
	err := s.repo.UpdateApplicationEnvPatch(ctx, userUUID, appEnvUuid, req)
	if err != nil {
		return nil, err
	}

	return &entities.UpdateApplicationEnvPatchResponse{
		ApplicationEnvUuid: *appEnvUuid,
		Status:             req.Status,
	}, nil
}

func (s *service) UpdateApplicationEnv(ctx context.Context, companyUuid, userUuid, applicationUuid, appEnvUuid *uuid.UUID, req *entities.ApplicationEnvUpdateRequestBody) (*entities.ApplicationEnv, error) {
	oldAppEnv, err := s.repo.GetApplicationEnvById(ctx, appEnvUuid)
	if err != nil {
		return nil, err
	}

	applicationEnv := &entities.ApplicationEnv{
		ApplicationEnvUuid:      *appEnvUuid,
		CompanyUuid:             *companyUuid,
		TechInfoApplicationUuid: *applicationUuid,
		Type:                    req.Type,
		Name:                    req.Name,
		Description:             req.Description,
		Url:                     req.Url,
		HostingProviderType:     req.HostingProviderType,
		HostingProvider:         req.HostingProvider,
		MfaType:                 req.MfaType,
		Mfa:                     req.Mfa,
		IDsIpsType:              req.IDsIpsType,
		IDsIps:                  req.IDsIps,
		Status:                  req.Status,
		UpdatedBy:               *userUuid,
	}

	applicationEnv.CreatedAt = oldAppEnv.CreatedAt
	applicationEnv.CreatedBy = oldAppEnv.CreatedBy

	appEnv, err := s.repo.UpdateApplicationEnv(ctx, applicationEnv)
	if err != nil {
		return nil, err
	}
	comment := s.getApplicationEnvJiraComment(appEnv, "Updated application environment: (updated application environment details are below)")

	err = s.jiraClient.AddComment(ctx, companyUuid, comment)
	if err != nil {
		return nil, err
	}

	return appEnv, nil
}

func (s *service) CreateApplicationEnv(ctx context.Context, companyUuid, userUuid, applicationUuid *uuid.UUID, req *entities.ApplicationEnvCreateRequestBody) (*entities.ApplicationEnv, error) {
	applicationEnv := &entities.ApplicationEnv{
		CompanyUuid:             *companyUuid,
		TechInfoApplicationUuid: *applicationUuid,
		Type:                    req.Type,
		Name:                    req.Name,
		Description:             req.Description,
		Url:                     req.Url,
		HostingProviderType:     req.HostingProviderType,
		HostingProvider:         req.HostingProvider,
		MfaType:                 req.MfaType,
		Mfa:                     req.Mfa,
		IDsIpsType:              req.IDsIpsType,
		IDsIps:                  req.IDsIps,
		Status:                  "Active",
		CreatedBy:               *userUuid,
	}

	appEnv, err := s.repo.CreateApplicationEnv(ctx, applicationEnv)
	if err != nil {
		return nil, err
	}
	comment := s.getApplicationEnvJiraComment(appEnv, "New application environment is created:")

	err = s.jiraClient.AddComment(ctx, companyUuid, comment)
	if err != nil {
		return nil, err
	}

	return appEnv, nil
}

func (s *service) CreateApplication(ctx context.Context, companyUuid, userUuid *uuid.UUID, application *entities.TechInfoApplication) (*entities.TechInfoApplicationCreateUpdateResponse, error) {
	application.CompanyUuid = *companyUuid
	application.CreatedBy = *userUuid

	app, err := s.repo.CreateApplication(ctx, application)
	if err != nil {
		return nil, err
	}

	comment := s.getApplicationJiraComment(app, "New application is created:")

	err = s.jiraClient.AddComment(ctx, companyUuid, comment)
	if err != nil {
		return nil, err
	}

	s.onboardingClient.UpdateOnboardingStatus(ctx, onboardingEntities.SetTechnicalInformation, *companyUuid, *userUuid)

	return &entities.TechInfoApplicationCreateUpdateResponse{
		TechInfoApplicationUuid: app.TechInfoApplicationUuid,
		Name:                    app.Name,
		Type:                    app.Type,
		CreatedAt:               app.CreatedAt,
		UpdatedAt:               app.UpdatedAt,
		CreatedBy:               app.CreatedBy,
		UpdatedBy:               app.UpdatedBy,
	}, nil
}

func (s *service) UpdateApplication(ctx context.Context, companyUuid, userUuid, applicationUuid *uuid.UUID, application *entities.TechInfoApplication) (*entities.TechInfoApplicationCreateUpdateResponse, error) {
	application.CompanyUuid = *companyUuid
	application.UpdatedBy = *userUuid
	application.TechInfoApplicationUuid = *applicationUuid

	app, err := s.repo.UpdateApplication(ctx, application)
	if err != nil {
		return nil, err
	}

	comment := s.getApplicationJiraComment(app, "Updated application: (updated application details are below)")

	err = s.jiraClient.AddComment(ctx, companyUuid, comment)
	if err != nil {
		return nil, err
	}

	return &entities.TechInfoApplicationCreateUpdateResponse{
		TechInfoApplicationUuid: app.TechInfoApplicationUuid,
		Name:                    app.Name,
		Type:                    app.Type,
		CreatedAt:               app.CreatedAt,
		UpdatedAt:               app.UpdatedAt,
		CreatedBy:               app.CreatedBy,
		UpdatedBy:               app.UpdatedBy,
	}, nil
}

func (s *service) UpdateApplicationPatch(ctx context.Context, companyUuid, userUUID, applicationUuid *uuid.UUID, req *entities.UpdateApplicationPatchRequestBody) (*entities.UpdateApplicationPatchResponse, error) {
	err := s.repo.UpdateApplicationPatch(ctx, userUUID, applicationUuid, req)
	if err != nil {
		return nil, err
	}

	return &entities.UpdateApplicationPatchResponse{
		TechInfoApplicationUuid: *applicationUuid,
		Status:                  req.Status,
	}, nil
}

func (s *service) GetApplicationById(ctx context.Context, companyUuid, userUuid, applicationUuid *uuid.UUID) (*entities.TechInfoApplication, error) {
	return s.repo.GetApplicationById(ctx, applicationUuid)
}

func (s *service) GetAllApplications(ctx context.Context, req *entities.GetAllApplicationsRequest) ([]*entities.TechInfoApplication, error) {
	return s.repo.GetAllApplications(ctx, &req.CompanyUuid, req.Keyword)
}

func (s *service) DeleteApplication(ctx context.Context, companyUuid, userUuid, applicationUuid *uuid.UUID) error {
	app, err := s.repo.GetApplicationById(ctx, applicationUuid)
	if err != nil {
		return err
	}

	_, err = s.repo.DeleteApplication(ctx, applicationUuid)
	if err != nil {
		return err
	}

	if app != nil {
		comment := s.getApplicationJiraComment(app, "Deleted application: (Deleted application details are below)")

		err = s.jiraClient.AddComment(ctx, companyUuid, comment)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *service) getApplicationEnvJiraComment(appEnv *entities.ApplicationEnv, title string) *jiraEntities.Comment {
	return &jiraEntities.Comment{
		Body: jiraEntities.Body{
			Type:    "doc",
			Version: 1,
			Content: []jiraEntities.ContentMain{
				{
					Type: "paragraph",
					Content: []jiraEntities.ContentInner{
						{
							Text: title,
							Type: "text",
							Marks: []jiraEntities.Marks{
								{
									Type: "strong",
								},
							},
						},
					},
				},
				{
					Type: "paragraph",
					Content: []jiraEntities.ContentInner{
						{
							Text: "Type: ",
							Type: "text",
							Marks: []jiraEntities.Marks{
								{
									Type: "strong",
								},
							},
						},
						{
							Text: appEnv.Type,
							Type: "text",
						},
					},
				},
				{
					Type: "paragraph",
					Content: []jiraEntities.ContentInner{
						{
							Text: "Name: ",
							Type: "text",
							Marks: []jiraEntities.Marks{
								{
									Type: "strong",
								},
							},
						},
						{
							Text: appEnv.Name,
							Type: "text",
						},
					},
				},
				{
					Type: "paragraph",
					Content: []jiraEntities.ContentInner{
						{
							Text: "URL: ",
							Type: "text",
							Marks: []jiraEntities.Marks{
								{
									Type: "strong",
								},
							},
						},
						{
							Text: appEnv.Url,
							Type: "text",
						},
					},
				},
				{
					Type: "paragraph",
					Content: []jiraEntities.ContentInner{
						{
							Text: "Hosting Provider Type: ",
							Type: "text",
							Marks: []jiraEntities.Marks{
								{
									Type: "strong",
								},
							},
						},
						{
							Text: appEnv.HostingProviderType,
							Type: "text",
						},
					},
				},
				{
					Type: "paragraph",
					Content: []jiraEntities.ContentInner{
						{
							Text: "Hosting Provider: ",
							Type: "text",
							Marks: []jiraEntities.Marks{
								{
									Type: "strong",
								},
							},
						},
						{
							Text: appEnv.HostingProvider,
							Type: "text",
						},
					},
				},
				{
					Type: "paragraph",
					Content: []jiraEntities.ContentInner{
						{
							Text: "Mfa Type: ",
							Type: "text",
							Marks: []jiraEntities.Marks{
								{
									Type: "strong",
								},
							},
						},
						{
							Text: appEnv.MfaType,
							Type: "text",
						},
					},
				},
				{
					Type: "paragraph",
					Content: []jiraEntities.ContentInner{
						{
							Text: "Mfa: ",
							Type: "text",
							Marks: []jiraEntities.Marks{
								{
									Type: "strong",
								},
							},
						},
						{
							Text: appEnv.Mfa,
							Type: "text",
						},
					},
				},
				{
					Type: "paragraph",
					Content: []jiraEntities.ContentInner{
						{
							Text: "Ids Ips Type: ",
							Type: "text",
							Marks: []jiraEntities.Marks{
								{
									Type: "strong",
								},
							},
						},
						{
							Text: appEnv.IDsIpsType,
							Type: "text",
						},
					},
				},
				{
					Type: "paragraph",
					Content: []jiraEntities.ContentInner{
						{
							Text: "Ids Ips: ",
							Type: "text",
							Marks: []jiraEntities.Marks{
								{
									Type: "strong",
								},
							},
						},
						{
							Text: appEnv.IDsIps,
							Type: "text",
						},
					},
				},
				{
					Type: "paragraph",
					Content: []jiraEntities.ContentInner{
						{
							Text: "Status: ",
							Type: "text",
							Marks: []jiraEntities.Marks{
								{
									Type: "strong",
								},
							},
						},
						{
							Text: appEnv.Status,
							Type: "text",
						},
					},
				},
			},
		},
	}
}

func (s *service) getApplicationJiraComment(application *entities.TechInfoApplication, title string) *jiraEntities.Comment {
	return &jiraEntities.Comment{
		Body: jiraEntities.Body{
			Type:    "doc",
			Version: 1,
			Content: []jiraEntities.ContentMain{
				{
					Type: "paragraph",
					Content: []jiraEntities.ContentInner{
						{
							Text: title,
							Type: "text",
							Marks: []jiraEntities.Marks{
								{
									Type: "strong",
								},
							},
						},
					},
				},
				{
					Type: "paragraph",
					Content: []jiraEntities.ContentInner{
						{
							Text: "Name: ",
							Type: "text",
							Marks: []jiraEntities.Marks{
								{
									Type: "strong",
								},
							},
						},
						{
							Text: application.Name,
							Type: "text",
						},
					},
				},
				{
					Type: "paragraph",
					Content: []jiraEntities.ContentInner{
						{
							Text: "Type: ",
							Type: "text",
							Marks: []jiraEntities.Marks{
								{
									Type: "strong",
								},
							},
						},
						{
							Text: application.Type,
							Type: "text",
						},
					},
				},
			},
		},
	}
}

func New(repo repository.Repository, onboardingClient onboarding.Client, jiraClient jira.Client) Service {
	return &service{
		repo:             repo,
		onboardingClient: onboardingClient,
		jiraClient:       jiraClient,
	}
}
