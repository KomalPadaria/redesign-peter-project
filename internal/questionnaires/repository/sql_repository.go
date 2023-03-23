package repository

import (
	"context"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/lib/pq"
	companyEntities "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company/entities"
	frameworkEntity "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/frameworks/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/questionnaires/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/db"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/nullable"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type sqlRepository struct {
	gormDB *gorm.DB
	sqlDB  *sql.DB
}

func (s *sqlRepository) GetAllSubscriptionsByCompany(ctx context.Context, companyUuid *uuid.UUID) ([]companyEntities.CompanySubscription, error) {
	var subs []companyEntities.CompanySubscription
	result := s.gormDB.WithContext(ctx).Where("company_uuid = ?", companyUuid).Find(&subs)
	if result.Error != nil {
		return nil, result.Error
	}

	return subs, nil
}

func (s *sqlRepository) GetQuestionnairesByUuids(ctx context.Context, uuids []uuid.UUID) ([]*entities.Questionnaires, error) {
	var qs []*entities.Questionnaires
	err := s.gormDB.WithContext(ctx).Where("questionnaires_uuid IN ?", uuids).Preload("Options").Find(&qs).Error
	if err != nil {
		return nil, err
	}

	return qs, nil
}

func (s *sqlRepository) PostAnswer(ctx context.Context, companyUUID, userUUID *uuid.UUID, answers []*entities.AnswerBody) error {
	err := s.gormDB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for _, a := range answers {
			err := s.removeQuestionnaireAnswersIfExistsByQuestionnairesUuid(ctx, a.QuestionnairesUuid, *companyUUID)
			if err != nil {
				return err
			}

			qaUuid := uuid.New()
			qa := entities.CreateQuestionnaireAnswers{
				QuestionnaireAnswersUuid: qaUuid,
				CompanyUuid:              *companyUUID,
				QuestionnairesUuid:       a.QuestionnairesUuid,
				Comment:                  a.Answer.Comment,
				CreatedAt:                nullable.NewNullTime(time.Now()),
				CreatedBy:                *userUUID,
			}

			err = tx.Create(&qa).Error
			if err != nil {
				return err
			}
			var qaos []entities.QuestionnaireAnswersOptions
			for _, o := range a.Answer.Options {
				qao := entities.QuestionnaireAnswersOptions{
					QuestionnaireOptionsUuid: o,
					QuestionnaireAnswersUuid: qaUuid,
				}
				qaos = append(qaos, qao)

			}

			err = tx.CreateInBatches(qaos, 10).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *sqlRepository) GetCategories(ctx context.Context, companyUuid, userUuid *uuid.UUID) ([]*entities.Category, error) {
	var cs []*entities.Category

	err := s.gormDB.WithContext(ctx).Table("frameworks_questionnaires fq").
		Select("q.category,count(q.questionnaires_uuid) as total,count(qa.questionnaire_answers_uuid) as completed").
		Joins("left join questionnaires q on q.questionnaires_uuid = fq.questionnaires_uuid").
		Joins("left join company_frameworks cf on cf.frameworks_uuid = fq.frameworks_uuid").
		Joins("left join questionnaire_answers qa on qa.questionnaires_uuid = q.questionnaires_uuid and qa.company_uuid = ?", companyUuid).
		Where("cf.company_uuid = ?", companyUuid).
		Group("q.category").
		Order("q.category").
		Find(&cs).Error

	if err != nil {
		return nil, err
	}

	return cs, nil
}

func (s *sqlRepository) GetQuestionnairesByCategory(ctx context.Context, companyUuid, userUuid *uuid.UUID, category string) ([]*entities.Questionnaires, error) {
	var qs []*entities.Questionnaires
	err := s.gormDB.WithContext(ctx).Table("questionnaires").
		Joins("left join frameworks_questionnaires fq on fq.questionnaires_uuid = questionnaires.questionnaires_uuid").
		Joins("left join company_frameworks cf on cf.frameworks_uuid = fq.frameworks_uuid").
		Preload("Options").
		Preload("Answer", "questionnaire_answers.company_uuid = ?", companyUuid).
		Preload("Answer.Options").
		Preload("Answer.Created").
		Where("questionnaires.category = ? AND cf.company_uuid = ?", category, companyUuid).
		Order("created_at desc").
		Find(&qs).Error

	if err != nil {
		return nil, err
	}

	return qs, err
}

func (s *sqlRepository) CreateCompanyQuestionnaires(ctx context.Context, companyUUID, userUUID *uuid.UUID) error {
	categories, err := s.getCategories(ctx)
	if err != nil {
		return err
	}
	companyQuestionnaires := make([]entities.CompanyQuestionnaires, 0)

	for _, c := range categories {
		cq := entities.CompanyQuestionnaires{
			CompanyQuestionnairesUuid: uuid.New(),
			CompanyUuid:               *companyUUID,
			Category:                  c,
			CreatedAt:                 nullable.NewNullTime(time.Now()),
			CreatedBy:                 *userUUID,
		}

		companyQuestionnaires = append(companyQuestionnaires, cq)
	}

	result := s.gormDB.WithContext(ctx).Create(&companyQuestionnaires)
	if result.Error != nil {
		err := result.Error
		if db.IsForeignKeyViolationError(err) {
			return errors.WithMessage(err, "violates foreign key constraint")
		}

		return err
	}

	return nil
}

// GetFrameworks TODO find a better logic to use the framework client
func (s *sqlRepository) GetFrameworks(ctx context.Context, companyUuid, userUuid *uuid.UUID) ([]*frameworkEntity.Framework, error) {
	var frameworks []*frameworkEntity.Framework

	err := s.gormDB.WithContext(ctx).Model(&frameworkEntity.Framework{}).
		Order("created_at desc").
		Joins("left join public.company_frameworks ON company_frameworks.frameworks_uuid = frameworks.frameworks_uuid").
		Find(&frameworks, "company_uuid = ?", companyUuid).Error
	if err != nil {
		return nil, err
	}

	return frameworks, nil
}

// GetCompanyByUUID TODO find a better logic to use the company client
func (s *sqlRepository) GetCompanyByUUID(ctx context.Context, companyUuid *uuid.UUID) (*companyEntities.Company, error) {
	company := &companyEntities.Company{}
	err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select("company_uuid",
			"name",
			"to_json(industry_type)",
			"onboarding",
			"address",
			"COALESCE(external_id, '')",
			"COALESCE(knowbe4_token, '')",
			"COALESCE(jira_epic_id, '')",
			"campaign_users",
			"rapid7_site_ids",
		).
		From("public.companies").
		Where(sq.Eq{"company_uuid": companyUuid}).
		RunWith(s.sqlDB).
		QueryRowContext(ctx).Scan(
		&company.CompanyUuid,
		&company.Name,
		&company.IndustryType,
		&company.Onboarding,
		&company.Address,
		&company.ExternalId,
		&company.Knowbe4Token,
		&company.JiraEpicId,
		&company.CampaignUsers,
		pq.Array(&company.Rapid7SiteIds),
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return company, nil
}

// UpdateCompany TODO find a better logic to use the company client
func (r *sqlRepository) UpdateCompany(ctx context.Context, updateData *companyEntities.UpdateCompanyRequest) error {

	sql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Update("public.companies").
		SetMap(updateData.Data).
		Where(sq.Eq{"company_uuid": updateData.CompanyUuid}).RunWith(r.sqlDB)

	_, err := sql.Exec()
	if err != nil {
		return err
	}

	return nil
}

func (s *sqlRepository) AddAnswer(ctx context.Context, companyUUID, userUUID, questionnairesUUID *uuid.UUID, files *entities.EvidenceFiles, answer *entities.Answer) error {
	qaUuid := uuid.New()
	qa := entities.CreateQuestionnaireAnswers{
		QuestionnaireAnswersUuid: qaUuid,
		CompanyUuid:              *companyUUID,
		QuestionnairesUuid:       *questionnairesUUID,
		Comment:                  answer.Comment,
		CreatedAt:                nullable.NewNullTime(time.Now()),
		CreatedBy:                *userUUID,
		Files:                    *files,
	}

	result := s.gormDB.WithContext(ctx).Create(&qa)
	if result.Error != nil {
		err := result.Error
		if db.IsForeignKeyViolationError(err) {
			return errors.WithMessage(err, "violates foreign key constraint")
		}
		return err
	}

	return nil
}

func (s *sqlRepository) UpdateQuestionnaireAnswerByUUID(ctx context.Context, answerUUID *uuid.UUID, data map[string]interface{}) error {
	result := s.gormDB.WithContext(ctx).Model(&entities.CreateQuestionnaireAnswers{}).Where("questionnaire_answers_uuid = ?", answerUUID).Updates(data)
	if result.Error != nil {
		err := result.Error
		if db.IsForeignKeyViolationError(err) {
			return errors.WithMessage(err, "violates foreign key constraint")
		}

		return err
	}
	return nil
}

func (s *sqlRepository) GetQuestionnaireAnswer(ctx context.Context, companyUUID, questionnaireUUID, answerUUID *uuid.UUID) (*entities.QuestionnaireAnswers, error) {
	var qa *entities.QuestionnaireAnswers
	err := s.gormDB.WithContext(ctx).
		Where("questionnaire_answers_uuid = ?", answerUUID).
		Where("company_uuid = ?", companyUUID).
		Where("questionnaires_uuid = ?", questionnaireUUID).
		Find(&qa).Error
	if err != nil {
		return nil, err
	}

	return qa, nil
}

func (s *sqlRepository) getCategories(ctx context.Context) ([]string, error) {
	var categories []string

	err := s.gormDB.WithContext(ctx).
		Model(&entities.Questionnaires{}).
		Distinct().
		Pluck("category", &categories).Error
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (s *sqlRepository) removeQuestionnaireAnswersIfExistsByQuestionnairesUuid(ctx context.Context, questionnairesUuid, companyUuid uuid.UUID) error {
	var qa entities.QuestionnaireAnswers
	result := s.gormDB.WithContext(ctx).First(&qa, "questionnaires_uuid = ?", questionnairesUuid)
	if result.Error != nil {
		if result.Error != gorm.ErrRecordNotFound {
			return result.Error
		}
	}

	if result.RowsAffected > 0 {
		err := s.gormDB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
			err := tx.Delete(&entities.QuestionnaireAnswersOptions{}, "questionnaire_answers_uuid = ?", qa.QuestionnaireAnswersUuid).Error
			if err != nil {
				return err
			}

			err = tx.Delete(&entities.QuestionnaireAnswers{}, "questionnaires_uuid = ? and company_uuid = ?", questionnairesUuid, companyUuid).Error
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	return nil
}
