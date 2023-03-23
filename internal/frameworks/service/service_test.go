package service

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/frameworks/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/frameworks/repository"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/salesforce"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/nullable"
	. "github.com/smartystreets/goconvey/convey"
	"go.uber.org/zap"
)

func Test_service_GetFrameworks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := repository.NewMockRepository(ctrl)

	ctx := context.Background()

	companyUUID := uuid.New()
	userUUID := uuid.New()

	salesforceClient := salesforce.NewMockClient(ctrl)

	logger := zap.NewNop().Sugar()

	svc := New(repo, salesforceClient, logger)

	expected := []*entities.Framework{
		{
			FrameworkUuid: uuid.New(),
			Name:          gofakeit.Word(),
			CreatedAt:     nullable.NullTime{},
			UpdatedAt:     nullable.NullTime{},
			CreatedBy:     uuid.New(),
			UpdatedBy:     uuid.New(),
		},
		{
			FrameworkUuid: uuid.New(),
			Name:          gofakeit.Word(),
			CreatedAt:     nullable.NullTime{},
			UpdatedAt:     nullable.NullTime{},
			CreatedBy:     uuid.New(),
			UpdatedBy:     uuid.New(),
		},
	}

	repo.EXPECT().GetFrameworks(gomock.Any(), gomock.Any(), gomock.Any()).Return(expected, nil)

	Convey("Given company_id and user_id", t, func() {
		reqBody := &entities.GetFrameworksRequest{
			CompanyUuid: companyUUID,
			UserUuid:    userUUID,
		}
		Convey("Call the GetFrameworks function", func() {
			actual, err := svc.GetFrameworks(ctx, &reqBody.CompanyUuid, &reqBody.UserUuid)
			Convey("The value should be equal to the expected value", func() {
				So(err, ShouldBeNil)
				So(actual, ShouldNotBeNil)
				So(actual, ShouldResemble, expected)
			})
		})
	})
}

func Test_service_GetFrameworkControls(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockRepository(ctrl)

	ctx := context.Background()

	companyUUID := uuid.New()
	userUUID := uuid.New()
	frameworkUUID := uuid.New()

	salesforceClient := salesforce.NewMockClient(ctrl)

	logger := zap.NewNop().Sugar()

	svc := New(mockRepo, salesforceClient, logger)
	Convey("Given company_id, user_id and framework_id", t, func() {
		reqBody := &entities.GetFrameworkControlRequest{
			CompanyUuid:   companyUUID,
			UserUuid:      userUUID,
			FrameworkUuid: frameworkUUID,
		}
		Convey("Call the GetFrameworkControls function", func() {
			actual, err := svc.GetFrameworkControls(ctx, &reqBody.CompanyUuid, &reqBody.UserUuid, &reqBody.FrameworkUuid)
			Convey("The value should be equal to the expected value", func() {
				So(err, ShouldBeNil)
				So(actual, ShouldNotBeNil)
				Convey("Control Topic & Domain should not be empty", func() {
					for _, control := range actual {
						So(control.Topic, ShouldNotBeEmpty)
						So(control.Domain, ShouldNotBeEmpty)
					}
				})
			})
		})
	})
}
