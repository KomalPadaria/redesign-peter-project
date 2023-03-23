package service

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	cognitoService "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/aws_client/cognito/service"
	companyMockClient "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company"
	companyEntities "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company/entities"
	onboardingMockClient "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/onboarding"
	salesforceMockClient "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/salesforce"
	emailMockClient "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/ses"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/user/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/user/repository"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/cfg"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/nullable"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_service_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockRepository(ctrl)
	cognitoClient := cognitoService.NewMockService(ctrl)
	companyClient := companyMockClient.NewMockClient(ctrl)
	salesforceClient := salesforceMockClient.NewMockClient(ctrl)
	onboardingClient := onboardingMockClient.NewMockClient(ctrl)
	emailClient := emailMockClient.NewMockClient(ctrl)
	config := cfg.Config{}

	ctx := context.Background()

	companyUUID := uuid.New()
	userUUID := uuid.New()

	companyDetails := &companyEntities.Company{
		CompanyUuid:  companyUUID,
		Name:         "",
		UserRole:     "",
		IndustryType: &companyEntities.IndustryType{"Entertainment"},
		Onboarding:   []companyEntities.Onboarding{},
		Address:      &companyEntities.Address{},
		ExternalId:   "",
	}

	userDetails := &entities.User{
		UserUuid:           userUUID,
		CurrentCompanyUuid: companyUUID,
		Username:           gofakeit.Username(),
		FirstName:          gofakeit.FirstName(),
		LastName:           gofakeit.LastName(),
		Email:              gofakeit.Email(),
		Phone:              gofakeit.Phone(),
		IsFirstLogin:       true,
		ExternalID:         "",
		CreatedAt:          nullable.NullTime{},
		UpdatedAt:          nullable.NullTime{},
		CreatedBy:          [16]byte{},
		UpdatedBy:          [16]byte{},
		CompanyRole:        "",
		CompanyStatus:      "",
	}

	companyClient.EXPECT().FindByUUID(gomock.Any(), gomock.Any()).Return(companyDetails, nil)
	cognitoClient.EXPECT().InviteUser(gomock.Any(), gomock.Any(), false).Return(nil)
	mockRepo.EXPECT().FindByEmail(gomock.Any(), gomock.Any()).Return(userDetails, nil)

	svc := New(mockRepo, companyClient, salesforceClient, cognitoClient, onboardingClient, emailClient, config)
	Convey("Given company_id, user_id and request body details", t, func() {
		reqBody := &entities.CreateUserRequestBody{
			Email: gofakeit.Email(),
			Role:  "admin",
		}
		Convey("Call the CreateUser function", func() {
			actual, err := svc.CreateUser(ctx, companyUUID, userUUID, reqBody)
			Convey("The value should be equal to the expected value", func() {
				So(err, ShouldBeNil)
				So(actual, ShouldNotBeNil)
				So(actual, ShouldResemble, userDetails)
			})
		})
	})
}

func Test_service_ResendUserInvite(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockRepository(ctrl)
	cognitoClient := cognitoService.NewMockService(ctrl)
	companyClient := companyMockClient.NewMockClient(ctrl)
	salesforceClient := salesforceMockClient.NewMockClient(ctrl)
	onboardingClient := onboardingMockClient.NewMockClient(ctrl)
	emailClient := emailMockClient.NewMockClient(ctrl)
	config := cfg.Config{}

	ctx := context.Background()

	companyUUID := uuid.New()
	userUUID := uuid.New()

	companyDetails := &companyEntities.Company{
		CompanyUuid:  companyUUID,
		Name:         "",
		UserRole:     "",
		IndustryType: &companyEntities.IndustryType{"Entertainment"},
		Onboarding:   []companyEntities.Onboarding{},
		Address:      &companyEntities.Address{},
		ExternalId:   "",
	}

	companyClient.EXPECT().FindByUUID(gomock.Any(), gomock.Any()).Return(companyDetails, nil)
	cognitoClient.EXPECT().InviteUser(gomock.Any(), gomock.Any(), true).Return(nil)

	svc := New(mockRepo, companyClient, salesforceClient, cognitoClient, onboardingClient, emailClient, config)
	Convey("Given company_id, user_id and request body details", t, func() {
		reqBody := &entities.CreateUserRequestBody{
			Email: gofakeit.Email(),
			Role:  "admin",
		}
		Convey("Call the ResendUserInvite function", func() {
			err := svc.ResendUserInvite(ctx, companyUUID, userUUID, reqBody)
			Convey("The value should be equal to the expected value", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}
