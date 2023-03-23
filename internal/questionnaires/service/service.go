package service

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company"

	"github.com/google/uuid"
	companyEntities "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/jira"
	jiraEntities "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/jira/entities"

	s3client "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/aws_client/s3"
	s3Entities "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/aws_client/s3/entities"
	frameworksClient "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/frameworks/client"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/questionnaires/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/questionnaires/repository"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/salesforce"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/nullable"
	"go.uber.org/zap"
)

type Service interface {
	GetCategories(ctx context.Context, companyUuid, userUuid *uuid.UUID) ([]*entities.CategoryStats, error)
	GetQuestionnairesByCategory(ctx context.Context, companyUuid, userUuid *uuid.UUID, category string) ([]*entities.Questionnaires, error)
	PostAnswerEndpoint(ctx context.Context, companyUUID, userUUID *uuid.UUID, answers []*entities.AnswerBody) error
	CreateCompanyQuestionnaires(ctx context.Context, companyUUID, userUUID *uuid.UUID) error
	AddEngineerFeedback(ctx context.Context, req *entities.EngineerFeedbackRequest) error
	SubmitQuestionnaires(ctx context.Context, req *entities.SubmitQuestionnairesRequest) error
	AddAnswerWithEvidence(ctx context.Context, req *entities.AnswerWithEvidenceRequest) error
	UpdateAnswerWithEvidence(ctx context.Context, req *entities.UpdateAnswerWithEvidenceRequest) error
	DownloadQuestionnairesEvidence(ctx context.Context, req *entities.DownloadEvidenceRequest) ([]byte, error)
	DeleteQuestionnairesEvidence(ctx context.Context, req *entities.DeleteEvidenceRequest) error
}

type service struct {
	repo             repository.Repository
	logger           *zap.SugaredLogger
	frameworksClient frameworksClient.Client
	sfClient         salesforce.Client
	s3client         s3client.Client
	jiraClient       jira.Client
	companyClient    company.Client
}

// New service for user.
func New(repo repository.Repository,
	logger *zap.SugaredLogger,
	sfClient salesforce.Client,
	frameworksClient frameworksClient.Client,
	s3client s3client.Client,
	jiraClient jira.Client,
	companyClient company.Client) Service {
	svc := &service{
		repo:             repo,
		logger:           logger,
		frameworksClient: frameworksClient,
		sfClient:         sfClient,
		jiraClient:       jiraClient,
		s3client:         s3client,
		companyClient:    companyClient,
	}

	return svc
}

func (s *service) CreateCompanyQuestionnaires(ctx context.Context, companyUUID, userUUID *uuid.UUID) error {
	return s.repo.CreateCompanyQuestionnaires(ctx, companyUUID, userUUID)
}

func (s *service) GetCategories(ctx context.Context, companyUuid, userUuid *uuid.UUID) ([]*entities.CategoryStats, error) {
	//TODO this section need to be removed on salesforce webhook is implemented.
	// Because the company framework link will be created via webhook
	frameworks, err := s.repo.GetFrameworks(ctx, companyUuid, userUuid)
	if err != nil {
		s.logger.Warn("error getting  frameworks", err)
	} else {
		if len(frameworks) == 0 {
			err = s.companyClient.FetchCompanySubscriptionFromSF(ctx, companyUuid)
			if err != nil {
				return nil, err
			}
		}
	}

	categories, err := s.repo.GetCategories(ctx, companyUuid, userUuid)
	if err != nil {
		return nil, err
	}
	var categoryStats []*entities.CategoryStats
	for _, c := range categories {
		cs := &entities.CategoryStats{
			Category: c.Category,
			Stats: &entities.Stats{
				Total:     c.Total,
				Completed: c.Completed,
			},
		}

		categoryStats = append(categoryStats, cs)
	}

	return categoryStats, nil
}

func (s *service) GetQuestionnairesByCategory(ctx context.Context, companyUuid, userUuid *uuid.UUID, category string) ([]*entities.Questionnaires, error) {
	return s.repo.GetQuestionnairesByCategory(ctx, companyUuid, userUuid, category)
}

func (s *service) PostAnswerEndpoint(ctx context.Context, companyUUID, userUUID *uuid.UUID, answers []*entities.AnswerBody) error {

	var questionUuids []uuid.UUID

	for _, a := range answers {
		questionUuids = append(questionUuids, a.QuestionnairesUuid)
	}

	qs, err := s.repo.GetQuestionnairesByUuids(ctx, questionUuids)
	if err != nil {
		return err
	}

	questionnaireMap := make(map[uuid.UUID]*entities.Questionnaires)
	for _, q := range qs {
		questionnaireMap[q.QuestionnairesUuid] = q
	}

	// validate answers
	for _, a := range answers {
		question := questionnaireMap[a.QuestionnairesUuid]
		err = s.validateAnswer(a, questionnaireMap[a.QuestionnairesUuid])
		if err != nil {
			return err
		}

		if len(question.Options) == 0 {
			a.Answer.Options = nil
		}
	}

	return s.repo.PostAnswer(ctx, companyUUID, userUUID, answers)
}

func (s *service) AddEngineerFeedback(ctx context.Context, req *entities.EngineerFeedbackRequest) error {
	err := s.repo.UpdateQuestionnaireAnswerByUUID(ctx, &req.AnswerUuid, map[string]interface{}{
		"feedback":   req.Feedback,
		"updated_at": nullable.NewNullTime(time.Now()),
		"updated_by": req.UserUuid,
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *service) AddAnswerWithEvidence(ctx context.Context, req *entities.AnswerWithEvidenceRequest) error {
	var files entities.EvidenceFiles

	objects := make([]s3Entities.BatchUploadObject, 0)

	if len(req.Files) > 0 {
		for _, f := range req.Files {
			fileUUID := uuid.New()
			s3key := s.s3key(
				req.CompanyUuid.String(),
				req.QuestionnairesUuid.String(),
				fileUUID.String(),
				f.Filename,
			)
			file := entities.EvidenceFile{
				Id:       fileUUID,
				FileName: f.Filename,
				S3Key:    s3key,
			}

			objects = append(objects, s3Entities.BatchUploadObject{
				FileHeader: f,
				Key:        s3key,
				FileName:   f.Filename,
			})
			files = append(files, file)
		}

		err := s.s3client.UploadFilesFromFileHeaders(ctx, objects)
		if err != nil {
			return err
		}
	}

	err := s.repo.AddAnswer(ctx, &req.CompanyUuid, &req.UserUuid, &req.QuestionnairesUuid, &files, &req.Answer)

	if err != nil {
		return err
	}

	return nil
}

func (s *service) SubmitQuestionnaires(ctx context.Context, req *entities.SubmitQuestionnairesRequest) error {
	updateStatusRequest := companyEntities.UpdateCompanyRequest{
		CompanyUuid: req.CompanyUuid,
		Data: map[string]interface{}{
			"updated_by": req.UserUuid,
			"gap_status": req.Status.Status,
		},
	}
	//TODO use company client to update company
	err := s.repo.UpdateCompany(ctx, &updateStatusRequest)
	if err != nil {
		return err
	}

	err = s.jiraClient.AddComment(ctx, &req.CompanyUuid, &jiraEntities.Comment{
		Body: jiraEntities.Body{
			Type:    "doc",
			Version: 1,
			Content: []jiraEntities.ContentMain{
				{
					Type: "paragraph",
					Content: []jiraEntities.ContentInner{
						{
							Text: "Gap Analysis is completed by the customer.",
							Type: "text",
							Marks: []jiraEntities.Marks{
								{
									Type: "strong",
								},
							},
						},
					},
				},
			},
		},
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *service) UpdateAnswerWithEvidence(ctx context.Context, req *entities.UpdateAnswerWithEvidenceRequest) error {

	var files entities.EvidenceFiles

	objects := make([]s3Entities.BatchUploadObject, 0)

	qa, err := s.repo.GetQuestionnaireAnswer(ctx, &req.CompanyUuid, &req.QuestionnairesUuid, &req.AnswerUuid)
	if err != nil {
		return err
	}

	// Frontend only send files when new files need to be uploaded.
	// The logic is same as of adding new files to bucket
	if len(req.Files) > 0 {
		for _, f := range req.Files {
			fileUUID := uuid.New()
			s3key := s.s3key(
				req.CompanyUuid.String(),
				req.QuestionnairesUuid.String(),
				fileUUID.String(),
				f.Filename,
			)
			file := entities.EvidenceFile{
				Id:       fileUUID,
				FileName: f.Filename,
				S3Key:    s3key,
			}

			objects = append(objects, s3Entities.BatchUploadObject{
				FileHeader: f,
				Key:        s3key,
				FileName:   f.Filename,
			})
			files = append(files, file)
		}

		err := s.s3client.UploadFilesFromFileHeaders(ctx, objects)
		if err != nil {
			return err
		}

		// add any newly uploaded files to current json blob
		files = append(files, qa.Files...)
	} else {
		files = qa.Files
	}

	err = s.repo.UpdateQuestionnaireAnswerByUUID(ctx, &req.AnswerUuid, map[string]interface{}{
		"comment":    req.Answer.Comment,
		"files":      files,
		"updated_at": nullable.NewNullTime(time.Now()),
		"updated_by": req.UserUuid,
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *service) DownloadQuestionnairesEvidence(ctx context.Context, req *entities.DownloadEvidenceRequest) ([]byte, error) {
	var fileS3Path string

	// find full s3 path for the file to download
	qa, err := s.repo.GetQuestionnaireAnswer(ctx, &req.CompanyUuid, &req.QuestionnairesUuid, &req.AnswerUuid)
	if err != nil {
		return nil, err
	}

	for _, file := range qa.Files {
		if file.Id.String() == req.FileId {
			// file.S3Key has filename attached to it
			fileS3Path = file.S3Key
		}
	}

	fileBytes, err := s.s3client.DownloadFile(ctx, fileS3Path)
	if err != nil {
		return nil, err
	}

	return fileBytes, nil
}

func (s *service) DeleteQuestionnairesEvidence(ctx context.Context, req *entities.DeleteEvidenceRequest) error {
	var updatedFiles entities.EvidenceFiles
	var deletedFileS3Path string

	qa, err := s.repo.GetQuestionnaireAnswer(ctx, &req.CompanyUuid, &req.QuestionnairesUuid, &req.AnswerUuid)
	if err != nil {
		return err
	}

	if qa != nil {
		for _, file := range qa.Files {
			if file.Id.String() != req.FileId {
				untouchedFile := entities.EvidenceFile{
					Id:       file.Id,
					FileName: file.FileName,
					S3Key:    file.S3Key,
				}
				updatedFiles = append(updatedFiles, untouchedFile)
			} else {
				deletedFileS3Path = file.S3Key
			}
		}
	}

	// delete from s3
	err = s.s3client.DeleteFile(ctx, deletedFileS3Path)
	if err != nil {
		return err
	}

	// update files json
	err = s.repo.UpdateQuestionnaireAnswerByUUID(ctx, &req.AnswerUuid, map[string]interface{}{
		"files": updatedFiles,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *service) validateAnswer(answer *entities.AnswerBody, question *entities.Questionnaires) error {
	switch question.CommentType {
	case "optional":
		switch question.OptionType {
		case "radio":
			if len(answer.Answer.Options) > 1 {
				return errors.New("multiple options are not allowed")
			}
		case "multi":
		}
	case "none":
		if strings.TrimSpace(answer.Answer.Comment) != "" {
			return errors.New("comment is not allowed")
		}
		switch question.OptionType {
		case "radio":
			if len(answer.Answer.Options) > 1 {
				return errors.New("multiple options are not allowed")
			}
		case "multi":
		}
	case "mandatory":
		if strings.TrimSpace(answer.Answer.Comment) == "" {
			return errors.New("comment should not be empty")
		}
		switch question.OptionType {
		case "radio":
			if len(answer.Answer.Options) > 1 {
				return errors.New("multiple options are not allowed")
			}
		case "multi":
		}
	}

	return nil
}

// S3 Key
//
//	{company_uuid}/gap_analysis_evidences/{questionnaire_uuid}/{file_id}/file.png
func (s *service) s3key(companyUuid, questionnaire_id, file_id, filename string) string {
	return fmt.Sprintf("%s/gap_analysis_evidences/%s/%s/%s", companyUuid, questionnaire_id, file_id, filename)
}
