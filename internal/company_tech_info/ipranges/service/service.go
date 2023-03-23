package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_tech_info/ipranges/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_tech_info/ipranges/repository"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/jira"
	jiraEntities "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/jira/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/onboarding"
	onboardingEntities "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/onboarding/entities"
)

type Service interface {
	GetAllTechInfoIpRange(ctx context.Context, companyUuid, userUuid *uuid.UUID) ([]*entities.TechInfoIpRanges, error)
	CreateTechInfoIpRange(ctx context.Context, companyUuid, userUuid *uuid.UUID, ipRange *entities.TechInfoIpRanges) (*entities.TechInfoIpRanges, error)
	UpdateTechInfoIpRange(ctx context.Context, companyUuid, userUuid, ipRangeUuid *uuid.UUID, ipRange *entities.TechInfoIpRanges) (*entities.TechInfoIpRanges, error)
	UpdateTechInfoIpRangePatch(ctx context.Context, companyUuid, userUUID, ipRangeUuid *uuid.UUID, req *entities.UpdateTechInfoIpRangePatchRequestBody) (*entities.UpdateTechInfoIpRangePatchResponse, error)
	DeleteTechInfoIpRange(ctx context.Context, companyUuid, userUuid, ipRange *uuid.UUID) error
	UpdateTechInfoIpRangeStatusByFacilities(ctx context.Context, userUUID *uuid.UUID, facilityUuids []uuid.UUID, status string) error
}

type service struct {
	repo             repository.Repository
	onboardingClient onboarding.Client
	jiraClient       jira.Client
}

func (s *service) UpdateTechInfoIpRangeStatusByFacilities(ctx context.Context, userUUID *uuid.UUID, facilityUuids []uuid.UUID, status string) error {
	return s.repo.UpdateTechInfoIpRangeStatusByFacilities(ctx, userUUID, facilityUuids, status)
}

func (s *service) DeleteTechInfoIpRange(ctx context.Context, companyUuid, userUuid, ipRange *uuid.UUID) error {

	techiprange, err := s.repo.GetTechInfoIpRangeByUuid(ctx, ipRange)
	if err != nil {
		return err
	}

	_, err = s.repo.DeleteTechInfoIpRange(ctx, ipRange)
	if err != nil {
		return err
	}

	comment := s.getIpRangeJiraComment(techiprange, "Deleted IP Range: (Deleted IP Range details are below)")

	err = s.jiraClient.AddComment(ctx, companyUuid, comment)
	if err != nil {
		return err
	}
	return err
}

func (s *service) UpdateTechInfoIpRangePatch(ctx context.Context, companyUuid, userUUID, ipRangeUuid *uuid.UUID, req *entities.UpdateTechInfoIpRangePatchRequestBody) (*entities.UpdateTechInfoIpRangePatchResponse, error) {
	err := s.repo.UpdateTechInfoIpRangePatch(ctx, userUUID, ipRangeUuid, req)
	if err != nil {
		return nil, err
	}

	return &entities.UpdateTechInfoIpRangePatchResponse{
		TechInfoIpRangeUuid: *ipRangeUuid,
		Status:              req.Status,
	}, nil
}

func (s *service) UpdateTechInfoIpRange(ctx context.Context, companyUuid, userUuid, ipRangeUuid *uuid.UUID, ipRange *entities.TechInfoIpRanges) (*entities.TechInfoIpRanges, error) {
	ipRange.CompanyUuid = *companyUuid
	ipRange.UpdatedBy = *userUuid
	ipRange.TechInfoIpRangeUuid = *ipRangeUuid

	techiprange, err := s.repo.UpdateTechInfoIpRange(ctx, ipRange)
	if err != nil {
		return nil, err
	}

	comment := s.getIpRangeJiraComment(techiprange, "Updated IP Range: (updated IP Range details are below)")

	err = s.jiraClient.AddComment(ctx, companyUuid, comment)
	if err != nil {
		return nil, err
	}

	return techiprange, nil
}

func (s *service) CreateTechInfoIpRange(ctx context.Context, companyUuid, userUuid *uuid.UUID, ipRange *entities.TechInfoIpRanges) (*entities.TechInfoIpRanges, error) {
	ipRange.CompanyUuid = *companyUuid
	ipRange.CreatedBy = *userUuid

	techiprange, err := s.repo.CreateTechInfoIpRange(ctx, ipRange)
	if err != nil {
		return nil, err
	}

	comment := s.getIpRangeJiraComment(techiprange, "New IP Range is created:")

	err = s.jiraClient.AddComment(ctx, companyUuid, comment)
	if err != nil {
		return nil, err
	}

	s.onboardingClient.UpdateOnboardingStatus(ctx, onboardingEntities.SetTechnicalInformation, *companyUuid, *userUuid)

	return techiprange, nil
}

func (s *service) GetAllTechInfoIpRange(ctx context.Context, companyUuid, userUuid *uuid.UUID) ([]*entities.TechInfoIpRanges, error) {
	return s.repo.GetAllTechInfoIpRange(ctx, companyUuid)
}

func (s *service) getIpRangeJiraComment(ipRange *entities.TechInfoIpRanges, title string) *jiraEntities.Comment {
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
							Text: "IP Address: ",
							Type: "text",
							Marks: []jiraEntities.Marks{
								{
									Type: "strong",
								},
							},
						},
						{
							Text: ipRange.IpAddress,
							Type: "text",
						},
					},
				},
				{
					Type: "paragraph",
					Content: []jiraEntities.ContentInner{
						{
							Text: "Prefix: ",
							Type: "text",
							Marks: []jiraEntities.Marks{
								{
									Type: "strong",
								},
							},
						},
						{
							Text: fmt.Sprintf("%d", ipRange.IpSize),
							Type: "text",
						},
					},
				},
				{
					Type: "paragraph",
					Content: []jiraEntities.ContentInner{
						{
							Text: "Is External: ",
							Type: "text",
							Marks: []jiraEntities.Marks{
								{
									Type: "strong",
								},
							},
						},
						{
							Text: fmt.Sprintf("%t", ipRange.IsExternal),
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
							Text: ipRange.Status,
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
							Text: ipRange.CompanyFacility.Name,
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
							Text: ipRange.CompanyFacility.CompanyAddressString(),
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
