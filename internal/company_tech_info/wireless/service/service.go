package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_tech_info/wireless/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_tech_info/wireless/repository"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/jira"
	jiraEntities "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/jira/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/onboarding"
	onboardingEntities "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/onboarding/entities"
)

type Service interface {
	CreateTechInfoWireless(ctx context.Context, companyUuid, userUuid *uuid.UUID, wireless *entities.TechInfoWireless) (*entities.TechInfoWireless, error)
	UpdateTechInfoWireless(ctx context.Context, companyUuid, userUuid, wirelessUuid *uuid.UUID, wireless *entities.TechInfoWireless) (*entities.TechInfoWireless, error)
	GetTechInfoWirelessById(ctx context.Context, companyUuid, userUuid, wirelessUuid *uuid.UUID) (*entities.TechInfoWireless, error)
	GetAllTechInfoWirelesss(ctx context.Context, companyUuid, userUuid *uuid.UUID) ([]*entities.TechInfoWireless, error)
	DeleteTechInfoWireless(ctx context.Context, companyUuid, userUuid, wirelessUuid *uuid.UUID) error
	UpdateTechInfoWirelessPatch(ctx context.Context, companyUuid, userUUID, wirelessUuid *uuid.UUID, req *entities.UpdateTechInfoWirelessPatchRequestBody) (*entities.UpdateTechInfoWirelessPatchResponse, error)
	UpdateTechInfoWirelessStatusByFacilities(ctx context.Context, userUUID *uuid.UUID, facilityUuids []uuid.UUID, status string) error
}

type service struct {
	repo             repository.Repository
	onboardingClient onboarding.Client
	jiraClient       jira.Client
}

func (s *service) UpdateTechInfoWirelessStatusByFacilities(ctx context.Context, userUUID *uuid.UUID, facilityUuids []uuid.UUID, status string) error {
	return s.repo.UpdateTechInfoWirelessStatusByFacilities(ctx, userUUID, facilityUuids, status)
}

func (s *service) UpdateTechInfoWirelessPatch(ctx context.Context, companyUuid, userUUID, wirelessUuid *uuid.UUID, req *entities.UpdateTechInfoWirelessPatchRequestBody) (*entities.UpdateTechInfoWirelessPatchResponse, error) {
	err := s.repo.UpdateTechInfoWirelessPatch(ctx, userUUID, wirelessUuid, req)
	if err != nil {
		return nil, err
	}

	return &entities.UpdateTechInfoWirelessPatchResponse{
		TechInfoWirelessUuid: *wirelessUuid,
		Status:               req.Status,
	}, nil
}

func (s *service) CreateTechInfoWireless(ctx context.Context, companyUuid, userUuid *uuid.UUID, wireless *entities.TechInfoWireless) (*entities.TechInfoWireless, error) {
	wireless.CompanyUuid = *companyUuid
	wireless.CreatedBy = *userUuid

	techinfo, err := s.repo.CreateTechInfoWireless(ctx, wireless)
	if err != nil {
		return nil, err
	}

	comment := s.getWirelessJiraComment(techinfo, "New wireless is created:")

	err = s.jiraClient.AddComment(ctx, companyUuid, comment)
	if err != nil {
		return nil, err
	}

	s.onboardingClient.UpdateOnboardingStatus(ctx, onboardingEntities.SetTechnicalInformation, *companyUuid, *userUuid)

	return techinfo, nil
}

func (s *service) UpdateTechInfoWireless(ctx context.Context, companyUuid, userUuid, wirelessUuid *uuid.UUID, wireless *entities.TechInfoWireless) (*entities.TechInfoWireless, error) {
	wireless.CompanyUuid = *companyUuid
	wireless.UpdatedBy = *userUuid
	wireless.TechInfoWirelessUuid = *wirelessUuid

	tiWireless, err := s.repo.UpdateTechInfoWireless(ctx, wireless)
	if err != nil {
		return nil, err
	}
	comment := s.getWirelessJiraComment(tiWireless, "Updated wireless: (updated wireless details are below)")

	err = s.jiraClient.AddComment(ctx, companyUuid, comment)
	if err != nil {
		return nil, err
	}
	return tiWireless, nil
}

func (s *service) GetTechInfoWirelessById(ctx context.Context, companyUuid, userUuid, wirelessUuid *uuid.UUID) (*entities.TechInfoWireless, error) {
	return s.repo.GetTechInfoWirelessById(ctx, wirelessUuid)
}

func (s *service) GetAllTechInfoWirelesss(ctx context.Context, companyUuid, userUuid *uuid.UUID) ([]*entities.TechInfoWireless, error) {
	return s.repo.GetAllTechInfoWirelesss(ctx, companyUuid)
}

func (s *service) DeleteTechInfoWireless(ctx context.Context, companyUuid, userUuid, wirelessUuid *uuid.UUID) error {
	wireless, err := s.repo.GetTechInfoWirelessById(ctx, wirelessUuid)
	if err != nil {
		return err
	}
	_, err = s.repo.DeleteTechInfoWireless(ctx, wirelessUuid)
	if err != nil {
		return err
	}

	comment := s.getWirelessJiraComment(wireless, "Deleted wireless: (Deleted wireless details are below)")

	err = s.jiraClient.AddComment(ctx, companyUuid, comment)
	if err != nil {
		return err
	}

	return err
}

func (s *service) getWirelessJiraComment(wireless *entities.TechInfoWireless, title string) *jiraEntities.Comment {
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
							Text: "SSID: ",
							Type: "text",
							Marks: []jiraEntities.Marks{
								{
									Type: "strong",
								},
							},
						},
						{
							Text: wireless.Ssid,
							Type: "text",
						},
					},
				},
				{
					Type: "paragraph",
					Content: []jiraEntities.ContentInner{
						{
							Text: "Protocol Type: ",
							Type: "text",
							Marks: []jiraEntities.Marks{
								{
									Type: "strong",
								},
							},
						},
						{
							Text: wireless.ProtocolType,
							Type: "text",
						},
					},
				},
				{
					Type: "paragraph",
					Content: []jiraEntities.ContentInner{
						{
							Text: "Protocol: ",
							Type: "text",
							Marks: []jiraEntities.Marks{
								{
									Type: "strong",
								},
							},
						},
						{
							Text: wireless.Protocol,
							Type: "text",
						},
					},
				},
				{
					Type: "paragraph",
					Content: []jiraEntities.ContentInner{
						{
							Text: "Security Type: ",
							Type: "text",
							Marks: []jiraEntities.Marks{
								{
									Type: "strong",
								},
							},
						},
						{
							Text: wireless.SecurityType,
							Type: "text",
						},
					},
				},
				{
					Type: "paragraph",
					Content: []jiraEntities.ContentInner{
						{
							Text: "Security: ",
							Type: "text",
							Marks: []jiraEntities.Marks{
								{
									Type: "strong",
								},
							},
						},
						{
							Text: wireless.Security,
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
							Text: wireless.Status,
							Type: "text",
						},
					},
				},
				{
					Type: "paragraph",
					Content: []jiraEntities.ContentInner{
						{
							Text: "Facility Name: ",
							Type: "text",
							Marks: []jiraEntities.Marks{
								{
									Type: "strong",
								},
							},
						},
						{
							Text: wireless.CompanyFacility.Name,
							Type: "text",
						},
					},
				},
				{
					Type: "paragraph",
					Content: []jiraEntities.ContentInner{
						{
							Text: "Facility Address: ",
							Type: "text",
							Marks: []jiraEntities.Marks{
								{
									Type: "strong",
								},
							},
						},
						{
							Text: wireless.CompanyFacility.CompanyAddressString(),
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
