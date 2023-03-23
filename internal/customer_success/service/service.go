package service

import (
	"context"
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
	s3client "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/aws_client/s3"
	s3Entities "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/aws_client/s3/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company"
	companyEntity "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/customer_success/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/customer_success/repository"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/jira"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/salesforce"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/salesforce/config"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/nullable"
	"github.com/pkg/errors"
	"golang.org/x/exp/slices"
)

type Service interface {
	GetSubscriptions(ctx context.Context, req *entities.GetSubscriptionsRequest) (*entities.GetSubscriptionsResponse, error)
	GetSubscriptionPlans(ctx context.Context, req *entities.GetSubscriptionPlansRequest) (*entities.GetSubscriptionPlansResponse, error)
	GetConsumedHours(ctx context.Context, companyUuid, userUuid *uuid.UUID) ([]*entities.GetConsumedHoursResponse, error)
	UploadReports(ctx context.Context, req *entities.UploadReportsRequest) error
	DownloadReport(ctx context.Context, req *entities.DownloadReportRequest) ([]byte, error)
	DeleteReport(ctx context.Context, req *entities.DeleteReportRequest) error
	GetConsultingHours(ctx context.Context, companyUuid, userUuid *uuid.UUID) ([]*entities.GetConsultingHoursResponse, error)
	GetServiceReview(ctx context.Context, req *entities.GetServiceReviewsRequest) ([]*entities.GetServiceReviewResponse, error)
	UpdateServiceReviewStatus(ctx context.Context, req *entities.UpdateServiceReviewsStatusRequest) (*entities.Evidence, error)

	UploadServiceReport(ctx context.Context, req *entities.UploadEvidencesRequest) error
	DeleteEvidenceFile(ctx context.Context, req *entities.DeleteEvidenceFileRequest) error
	AddEvidenceFiles(ctx context.Context, req *entities.AddEvidenceFilesRequest) error
	DownloadEvidenceReport(ctx context.Context, req *entities.DownloadEvidenceReportRequest) ([]byte, error)
}

// New service for user.
func New(repo repository.Repository, jiraClient jira.Client, salesforceClient salesforce.Client, companyClient company.Client, s3client s3client.Client, config config.Config) Service {
	svc := &service{repo, jiraClient, salesforceClient, companyClient, s3client, config}

	return svc
}

type service struct {
	repo          repository.Repository
	jiraClient    jira.Client
	sfClient      salesforce.Client
	companyClient company.Client
	s3client      s3client.Client
	sfConfig      config.Config
}

func (s *service) DownloadEvidenceReport(ctx context.Context, req *entities.DownloadEvidenceReportRequest) ([]byte, error) {
	parts := strings.Split(req.ReportName, ".")
	fileId := parts[0]
	extension := parts[1]
	s3FileName := fileId + "." + extension
	key := s.s3key(req.CompanyUuid.String(), req.ServiceUuid.String(), s3FileName)

	fileBytes, err := s.s3client.DownloadFile(ctx, key)
	if err != nil {
		return nil, err
	}

	return fileBytes, nil
}

func (s *service) AddEvidenceFiles(ctx context.Context, req *entities.AddEvidenceFilesRequest) error {
	objects := make([]s3Entities.BatchUploadObject, 0)
	serviceReportDataList := make([]companyEntity.ServiceEvidenceData, 0)
	for _, f := range req.Files {
		fileId := uuid.New()
		extension := filepath.Ext(f.Filename)
		s3fileName := fileId.String() + extension

		objects = append(objects, s3Entities.BatchUploadObject{
			FileHeader: f,
			Key:        s.s3key(req.CompanyUuid.String(), req.ServiceUuid.String(), s3fileName),
			FileName:   s3fileName,
		})
		serviceReportDataList = append(serviceReportDataList, companyEntity.ServiceEvidenceData{
			FileName:  f.Filename,
			FileUuid:  fileId,
			FileExt:   extension,
			CreatedAt: time.Now(),
		})

	}

	err := s.s3client.UploadFilesFromFileHeaders(ctx, objects)
	if err != nil {
		return err
	}

	sEvidence, err := s.repo.GetEvidenceByEvidenceId(ctx, &req.EvidenceUuid)
	if err != nil {
		return err
	}
	filesData := append(sEvidence.Data, serviceReportDataList...)

	err = s.repo.UpdateEvidence(ctx, &req.EvidenceUuid, map[string]interface{}{
		"data": filesData,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *service) DeleteEvidenceFile(ctx context.Context, req *entities.DeleteEvidenceFileRequest) error {
	sEvidence, err := s.repo.GetEvidenceByEvidenceId(ctx, &req.EvidenceUuid)
	if err != nil {
		return err
	}

	filesData := make([]companyEntity.ServiceEvidenceData, 0)
	for _, e := range sEvidence.Data {
		if e.FileUuid == req.FileUuid {
			extension := filepath.Ext(e.FileName)
			s3fileName := e.FileUuid.String() + extension
			key := s.s3key(req.CompanyUuid.String(), req.ServiceUuid.String(), s3fileName)
			err = s.s3client.DeleteFile(ctx, key)
			if err != nil {
				return err
			}
		} else {
			filesData = append(filesData, e)
		}
	}

	b, _ := json.Marshal(filesData)
	err = s.repo.UpdateEvidence(ctx, &req.EvidenceUuid, map[string]interface{}{
		"data": string(b),
	})
	if err != nil {
		return err
	}

	if len(sEvidence.Data) == 1 {
		err = s.repo.DeleteEvidence(ctx, &req.EvidenceUuid)
		if err != nil {
			return err
		}

		return nil
	}

	return nil
}

func (s *service) UploadServiceReport(ctx context.Context, req *entities.UploadEvidencesRequest) error {
	objects := make([]s3Entities.BatchUploadObject, 0)
	serviceReportDataList := make([]companyEntity.ServiceEvidenceData, 0)
	for _, f := range req.Files {
		fileId := uuid.New()
		extension := filepath.Ext(f.Filename)
		s3fileName := fileId.String() + extension

		objects = append(objects, s3Entities.BatchUploadObject{
			FileHeader: f,
			Key:        s.s3key(req.CompanyUuid.String(), req.ServiceUuid.String(), s3fileName),
			FileName:   s3fileName,
		})
		serviceReportDataList = append(serviceReportDataList, companyEntity.ServiceEvidenceData{
			FileName:  f.Filename,
			FileUuid:  fileId,
			FileExt:   extension,
			CreatedAt: time.Now(),
		})

	}

	err := s.s3client.UploadFilesFromFileHeaders(ctx, objects)
	if err != nil {
		return err
	}

	serviceReports := companyEntity.ServiceEvidence{
		ServiceEvidencesUuid:     uuid.New(),
		CompanySubscriptionsUuid: req.ServiceUuid,
		Status:                   entities.ServiceReportsStatusTypeRequiredAcknowledgement,
		Data:                     serviceReportDataList,
	}
	err = s.repo.CreateEvidences(ctx, serviceReports)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetConsultingHours(ctx context.Context, companyUuid, userUuid *uuid.UUID) ([]*entities.GetConsultingHoursResponse, error) {
	var chSubs []companyEntity.CompanySubscription
	chSubs, err := s.companyClient.GetConsultingHoursSubscriptions(ctx, companyUuid)
	if err != nil {
		return nil, err
	}

	if len(chSubs) == 0 {
		err = s.companyClient.FetchCompanySubscriptionFromSF(ctx, companyUuid)
		if err != nil {
			return nil, err
		}

		chSubs, err = s.companyClient.GetConsultingHoursSubscriptions(ctx, companyUuid)
		if err != nil {
			return nil, err
		}
	}

	sfSess, err := s.sfClient.NewSession(ctx)
	if err != nil {
		return nil, err
	}

	var subIds []string
	for _, sub := range chSubs {
		subIds = append(subIds, sub.SfSubscriptionID)
	}

	response := make([]*entities.GetConsultingHoursResponse, 0)
	if len(subIds) > 0 {
		subscriptions, err := s.sfClient.GetSubscriptionsByIDs(ctx, sfSess, subIds)
		if err != nil {
			return nil, err
		}

		for _, sub := range subscriptions {
			response = append(response, &entities.GetConsultingHoursResponse{
				SubscriptionID: sub.ID,
				Name:           sub.SBQQProductNameC,
				HoursConsumed:  sub.QuantityInHourUsedC,
				HoursTotal:     sub.QuantityBillableHoursC,
				From:           sub.SBQQStartDateC,
				To:             sub.SBQQEndDateC,
			})
		}
	}

	return response, nil
}

func (s *service) DownloadReport(ctx context.Context, req *entities.DownloadReportRequest) ([]byte, error) {
	key := s.s3key(req.CompanyUuid.String(), req.ServiceName, req.ReportName)
	fileBytes, err := s.s3client.DownloadFile(ctx, key)
	if err != nil {
		return nil, err
	}

	return fileBytes, nil
}

func (s *service) UploadReports(ctx context.Context, req *entities.UploadReportsRequest) error {
	objects := make([]s3Entities.BatchUploadObject, 0)
	for _, f := range req.Files {
		objects = append(objects, s3Entities.BatchUploadObject{
			FileHeader: f,
			Key:        s.s3key(req.CompanyUuid.String(), req.ServiceName, f.Filename),
			FileName:   f.Filename,
		})

	}

	err := s.s3client.UploadFilesFromFileHeaders(ctx, objects)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) DeleteReport(ctx context.Context, req *entities.DeleteReportRequest) error {
	key := s.s3key(req.CompanyUuid.String(), req.ServiceName, req.ReportName)
	err := s.s3client.DeleteFile(ctx, key)
	if err != nil {
		return err
	}

	var subUuid uuid.UUID
	subUuid, err = uuid.Parse(req.ServiceName)
	if err != nil {
		cs, err := s.companyClient.GetSubscriptionByName(ctx, &req.CompanyUuid, req.ServiceName)
		if err != nil {
			return err
		}
		subUuid = cs.CompanySubscriptionsUuid
	}

	err = s.repo.DeleteEvidenceFile(ctx, &subUuid, req.ReportName)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetSubscriptions(ctx context.Context, req *entities.GetSubscriptionsRequest) (*entities.GetSubscriptionsResponse, error) {
	sfSess, err := s.sfClient.NewSession(ctx)
	if err != nil {
		return nil, err
	}

	company, err := s.companyClient.FindByUUID(ctx, &companyEntity.GetCompanyByIdRequest{
		CompanyUuid: req.CompanyUuid,
	})
	if err != nil {
		return nil, err
	}

	subscriptions, err := s.sfClient.GetSubscriptionsByAccountID(ctx, sfSess, company.ExternalId)
	if err != nil {
		return nil, err
	}

	response := &entities.GetSubscriptionsResponse{}
	response.Services = make([]entities.Service, 0)

	mainSubsProducts := strings.Split(s.sfConfig.MainSubscriptions, ",")
	serviceSubsProducts := strings.Split(s.sfConfig.ServiceSubscriptions, ",")

	for _, sub := range subscriptions {
		if slices.Contains(mainSubsProducts, sub.SBQQProductC) {
			response.SubscriptionName = sub.SBQQProductNameC
			response.StartDate = sub.SBQQStartDateC
			response.EndDate = sub.SBQQEndDateC
		}

		if slices.Contains(serviceSubsProducts, sub.SBQQProductC) {
			response.Services = append(response.Services, entities.Service{
				Name:      sub.SBQQProductNameC,
				StartDate: sub.SBQQStartDateC,
				EndDate:   sub.SBQQEndDateC,
			})
		}
	}

	return response, nil
}

func (s *service) GetSubscriptionPlans(ctx context.Context, req *entities.GetSubscriptionPlansRequest) (*entities.GetSubscriptionPlansResponse, error) {
	subs, err := s.companyClient.GetAllSubscriptionsByCompany(ctx, &req.CompanyUuid)
	if err != nil {
		return nil, err
	}

	if len(subs) == 0 {
		err = s.companyClient.FetchCompanySubscriptionFromSF(ctx, &req.CompanyUuid)
		if err != nil {
			return nil, err
		}

		subs, err = s.companyClient.GetAllSubscriptionsByCompany(ctx, &req.CompanyUuid)
		if err != nil {
			return nil, err
		}
	}

	for _, sub := range subs {
		productIds := strings.Split(s.sfClient.GetConfig().MainSubscriptions, ",")
		if slices.Contains(productIds, sub.SfProductID) {
			return &entities.GetSubscriptionPlansResponse{
				PlanId:    sub.CompanySubscriptionsUuid,
				PlanType:  sub.Name,
				Status:    sub.Status,
				StartDate: nullable.NewNullTime(sub.StartDate),
				EndDate:   nullable.NewNullTime(sub.EndDate),
			}, nil
		}
	}

	return nil, nil
}
func (s *service) GetConsumedHours(ctx context.Context, companyUuid, userUuid *uuid.UUID) ([]*entities.GetConsumedHoursResponse, error) {
	company, err := s.getCompanyByUuid(ctx, companyUuid)
	if err != nil {
		return nil, err
	}

	if company.JiraEpicId == "" {
		return nil, errors.New(fmt.Sprintf("JIRA Epic ID not found for the company '%s'", company.Name))
	}

	issues, err := s.jiraClient.GetIssuesByEpicId(ctx, company.JiraEpicId)
	if err != nil {
		return nil, err
	}

	response := make([]*entities.GetConsumedHoursResponse, 0)
	for _, i := range issues {
		ch := &entities.GetConsumedHoursResponse{
			TicketNumber: i.Key,
			Type:         i.Fields.Summary,
			AssignedTo:   i.Fields.Assignee.DisplayName,
		}

		var timeSpentInSec int
		for _, wl := range i.Fields.Worklog.Worklogs {
			timeSpentInSec += wl.TimeSpentSeconds
		}

		duration := time.Duration(timeSpentInSec) * time.Second

		ch.TimeConsumed = duration.String()

		for _, clHistory := range i.Changelog.Histories {
			for _, chItem := range clHistory.Items {
				if chItem.Field == "status" && chItem.Fieldtype == "jira" && chItem.ToString == entities.ChangelogStatus {
					t, err := time.Parse("2006-01-02T15:04:05.999-0700", clHistory.Created)
					if err != nil {
						return nil, err
					}

					ch.CloseDate = nullable.NewNullTime(t)
				}
			}
		}

		response = append(response, ch)

	}

	return response, nil
}

func (s *service) GetServiceReview(ctx context.Context, req *entities.GetServiceReviewsRequest) ([]*entities.GetServiceReviewResponse, error) {
	subs, err := s.companyClient.GetServiceReviewSubscriptions(ctx, &req.CompanyUuid)
	if err != nil {
		return nil, err
	}

	response := make([]*entities.GetServiceReviewResponse, 0)
	for _, sub := range subs {
		evidences := make([]entities.Evidence, 0)
		srResp := &entities.GetServiceReviewResponse{}
		srResp.ServiceReviewUuid = sub.CompanySubscriptionsUuid
		srResp.ServiceName = sub.Name
		for _, e := range sub.ServiceEvidence {
			evidence := entities.Evidence{}
			evidence.EvidenceId = e.ServiceEvidencesUuid
			evidence.CompletedOn = e.CompletedOn
			evidence.AcknowledgedAt = e.AcknowledgedAt
			evidence.AcknowledgedBy = strings.TrimSpace(e.AcknowledgedUser.FirstName + " " + e.AcknowledgedUser.LastName)
			evidence.Status = e.Status
			evidence.Data = make([]entities.FileData, 0)
			for _, d := range e.Data {
				evidence.Data = append(evidence.Data, entities.FileData{
					FileId:   d.FileUuid,
					FileName: d.FileName,
					FilePath: fmt.Sprintf("/companies/%s/services-review/%s/evidence/%s/reports/%s%s",
						req.CompanyUuid.String(), sub.CompanySubscriptionsUuid.String(), e.ServiceEvidencesUuid.String(), d.FileUuid.String(), d.FileExt),
				})
			}

			evidences = append(evidences, evidence)
		}
		srResp.Evidence = evidences

		response = append(response, srResp)
	}

	return response, nil
}
func (s *service) UpdateServiceReviewStatus(ctx context.Context, req *entities.UpdateServiceReviewsStatusRequest) (*entities.Evidence, error) {
	err := s.repo.UpdateEvidence(ctx, &req.EvidenceUuid, map[string]interface{}{
		"acknowledged_at": time.Now().UTC(),
		"acknowledged_by": req.UserUuid,
		"status":          req.PatchRequestBody.Status,
	})
	if err != nil {
		return nil, err
	}

	se, err := s.repo.GetEvidenceByEvidenceId(ctx, &req.EvidenceUuid)
	if err != nil {
		return nil, err
	}

	response := &entities.Evidence{}
	response.EvidenceId = se.ServiceEvidencesUuid
	response.CompletedOn = se.CompletedOn
	response.AcknowledgedAt = se.AcknowledgedAt
	response.AcknowledgedBy = strings.TrimSpace(se.AcknowledgedUser.FirstName + " " + se.AcknowledgedUser.LastName)
	response.Status = se.Status
	response.Data = make([]entities.FileData, 0)

	for _, d := range se.Data {
		response.Data = append(response.Data, entities.FileData{
			FileId:   d.FileUuid,
			FileName: d.FileName,
		})
	}

	return response, nil
}

func (s *service) s3key(companyUuid, serviceName, filename string) string {
	return fmt.Sprintf("%s/%s/%s", companyUuid, serviceName, filename)
}

func (s *service) getCompanyByUuid(ctx context.Context, companyUuid *uuid.UUID) (*companyEntity.Company, error) {
	return s.companyClient.FindByUUID(ctx, &companyEntity.GetCompanyByIdRequest{
		CompanyUuid: *companyUuid,
	})
}
