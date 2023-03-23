package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	companyEntities "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company/entities"
	frameworkEntity "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/frameworks/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/questionnaires/entities"
	"gorm.io/gorm"
)

// Repository for websites.
type Repository interface {
	CreateCompanyQuestionnaires(ctx context.Context, companyUUID, userUUID *uuid.UUID) error
	GetQuestionnairesByCategory(ctx context.Context, companyUuid, userUuid *uuid.UUID, category string) ([]*entities.Questionnaires, error)
	GetCategories(ctx context.Context, companyUuid, userUuid *uuid.UUID) ([]*entities.Category, error)
	PostAnswer(ctx context.Context, companyUUID, userUUID *uuid.UUID, answers []*entities.AnswerBody) error
	GetQuestionnairesByUuids(ctx context.Context, uuids []uuid.UUID) ([]*entities.Questionnaires, error)
	GetFrameworks(ctx context.Context, companyUuid, userUuid *uuid.UUID) ([]*frameworkEntity.Framework, error)
	GetCompanyByUUID(ctx context.Context, companyUuid *uuid.UUID) (*companyEntities.Company, error)
	AddAnswer(ctx context.Context, companyUUID, userUUID, questionnairesUUID *uuid.UUID, evidences *entities.EvidenceFiles, answer *entities.Answer) error
	UpdateQuestionnaireAnswerByUUID(ctx context.Context, answerUUID *uuid.UUID, data map[string]interface{}) error
	GetQuestionnaireAnswer(ctx context.Context, companyUUID, questionnaireUUID, answerUUID *uuid.UUID) (*entities.QuestionnaireAnswers, error)

	UpdateCompany(ctx context.Context, updateData *companyEntities.UpdateCompanyRequest) error
	GetAllSubscriptionsByCompany(ctx context.Context, companyUuid *uuid.UUID) ([]companyEntities.CompanySubscription, error)
}

// New repository for websites.
func New(gormDB *gorm.DB, sqlDB *sql.DB) Repository {
	repo := &sqlRepository{gormDB, sqlDB}

	return repo
}
