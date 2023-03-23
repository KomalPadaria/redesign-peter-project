package service

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_tech_info/externalinfra/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_tech_info/externalinfra/repository"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/nullable"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_service_CreateTechInfoExternalInfra(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := repository.NewMockRepository(ctrl)
	ctx := context.Background()

	companyUUID := uuid.New()
	userUUID := uuid.New()
	techInfoExternalInfraUuid := uuid.New()

	expected := &entities.TechInfoExternalInfra{
		TechInfoExternalInfraUuid: techInfoExternalInfraUuid,
		CompanyUuid:               companyUUID,
		IpFrom:                    "192.168.1.1",
		IpTo:                      "192.168.1.10",
		Env:                       "test env",
		Location:                  "test location",
		HasPermissions:            false,
		HasIDsIps:                 false,
		IsWhitelisted:             false,
		Is3rdPartyHosted:          false,
		CreatedAt:                 nullable.NewNullTime(time.Now()),
		CreatedBy:                 userUUID,
	}
	repo.EXPECT().CreateTechInfoExternalInfra(gomock.Any(), gomock.Any()).Return(expected, nil)

	svc := New(repo)
	Convey("Given company_id, user_id and techInfoExternalInfra details", t, func() {
		techInfoExternalInfra := &entities.TechInfoExternalInfra{
			IpFrom:           "192.168.1.1",
			IpTo:             "192.168.1.10",
			Env:              "test env",
			Location:         "test location",
			HasPermissions:   false,
			HasIDsIps:        false,
			IsWhitelisted:    false,
			Is3rdPartyHosted: false,
		}
		Convey("Call the CreateTechInfoExternalInfra function", func() {
			actual, err := svc.CreateTechInfoExternalInfra(ctx, &companyUUID, &userUUID, techInfoExternalInfra)
			Convey("The value should be equal to the expected value", func() {
				So(err, ShouldBeNil)
				So(actual, ShouldNotBeNil)
				So(actual, ShouldResemble, expected)
			})
		})
	})
}

func Test_service_UpdateTechInfoExternalInfra(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := repository.NewMockRepository(ctrl)
	ctx := context.Background()

	companyUUID := uuid.New()
	userUUID := uuid.New()
	techInfoExternalInfraUuid := uuid.New()

	expected := &entities.TechInfoExternalInfra{
		TechInfoExternalInfraUuid: techInfoExternalInfraUuid,
		CompanyUuid:               companyUUID,
		IpFrom:                    "192.168.1.1",
		IpTo:                      "192.168.1.10",
		Env:                       "test env",
		Location:                  "test location",
		HasPermissions:            false,
		HasIDsIps:                 false,
		IsWhitelisted:             false,
		Is3rdPartyHosted:          false,
		CreatedAt:                 nullable.NewNullTime(time.Now()),
		UpdatedAt:                 nullable.NewNullTime(time.Now()),
		CreatedBy:                 userUUID,
		UpdatedBy:                 userUUID,
	}

	repo.EXPECT().UpdateTechInfoExternalInfra(gomock.Any(), gomock.Any()).Return(expected, nil)

	svc := New(repo)
	Convey("Given company_id, user_id, application_id and techInfoExternalInfra details", t, func() {

		techInfoExternalInfra := &entities.TechInfoExternalInfra{}
		Convey("Call the UpdateTechInfoExternalInfra function", func() {
			actual, err := svc.UpdateTechInfoExternalInfra(ctx, &companyUUID, &userUUID, &techInfoExternalInfraUuid, techInfoExternalInfra)
			Convey("The value should be equal to the expected value", func() {
				So(err, ShouldBeNil)
				So(actual, ShouldNotBeNil)
				So(actual.CreatedAt, ShouldResemble, expected.CreatedAt)
			})
		})
	})
}

func Test_service_GetTechInfoExternalInfraById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	companyUUID, _ := uuid.Parse("a0f9e860-f3fa-4994-a8d4-eb156dfce805")
	userUUID, _ := uuid.Parse("64d1802e-1fa4-4732-b182-0719199108a8")
	techInfoExternalInfraUuid, _ := uuid.Parse("a0f9e860-f3fa-4994-a8d4-eb156dfce805")

	repo := repository.NewMockRepository(ctrl)
	ctx := context.Background()
	expected := &entities.TechInfoExternalInfra{
		TechInfoExternalInfraUuid: techInfoExternalInfraUuid,
		CompanyUuid:               companyUUID,
		IpFrom:                    "192.168.1.1",
		IpTo:                      "192.168.1.10",
		Env:                       "test env",
		Location:                  "test location",
		HasPermissions:            false,
		HasIDsIps:                 false,
		IsWhitelisted:             false,
		Is3rdPartyHosted:          false,
		CreatedAt:                 nullable.NewNullTime(time.Now()),
		UpdatedAt:                 nullable.NewNullTime(time.Now()),
		CreatedBy:                 userUUID,
		UpdatedBy:                 userUUID,
	}

	repo.EXPECT().GetTechInfoExternalInfraById(gomock.Any(), gomock.Any()).Return(expected, nil)

	svc := New(repo)
	Convey("Given company_id, user_id and application_id", t, func() {
		companyUUID, _ = uuid.Parse("a0f9e860-f3fa-4994-a8d4-eb156dfce805")
		userUUID, _ = uuid.Parse("64d1802e-1fa4-4732-b182-0719199108a8")
		techInfoExternalInfraUuid, _ = uuid.Parse("a0f9e860-f3fa-4994-a8d4-eb156dfce805")

		Convey("Call the GetTechInfoExternalInfraById function", func() {
			actual, err := svc.GetTechInfoExternalInfraById(ctx, &companyUUID, &userUUID, &techInfoExternalInfraUuid)
			Convey("The value should be equal to the expected value", func() {
				So(err, ShouldBeNil)
				So(actual, ShouldNotBeNil)
				So(actual, ShouldEqual, expected)
			})
		})
	})
}

func Test_service_GetAllTechInfoExternalInfras(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userUuid := uuid.New()

	repo := repository.NewMockRepository(ctrl)
	ctx := context.Background()
	expected := make([]*entities.TechInfoExternalInfra, 0)
	expected = append(expected, &entities.TechInfoExternalInfra{
		TechInfoExternalInfraUuid: uuid.New(),
		CompanyUuid:               uuid.New(),
		IpFrom:                    "192.168.1.1",
		IpTo:                      "192.168.1.10",
		Env:                       "test env",
		Location:                  "test location",
		HasPermissions:            false,
		HasIDsIps:                 false,
		IsWhitelisted:             false,
		Is3rdPartyHosted:          false,
		CreatedAt:                 nullable.NewNullTime(time.Now()),
		UpdatedAt:                 nullable.NewNullTime(time.Now()),
		CreatedBy:                 userUuid,
		UpdatedBy:                 userUuid,
	})

	repo.EXPECT().GetAllTechInfoExternalInfras(gomock.Any(), gomock.Any()).Return(expected, nil)

	svc := New(repo)
	Convey("Given company_id, user_id", t, func() {
		companyUUID, _ := uuid.Parse("a0f9e860-f3fa-4994-a8d4-eb156dfce805")
		userUUID, _ := uuid.Parse("64d1802e-1fa4-4732-b182-0719199108a8")

		Convey("Call the GetAllTechInfoExternalInfras function", func() {
			actual, err := svc.GetAllTechInfoExternalInfras(ctx, &companyUUID, &userUUID)
			Convey("The value should be equal to the expected value", func() {
				So(err, ShouldBeNil)
				So(actual, ShouldNotBeNil)
				So(actual, ShouldResemble, expected)
			})
		})
	})
}

func Test_service_DeleteTechInfoExternalInfra(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := repository.NewMockRepository(ctrl)
	ctx := context.Background()
	expected, _ := uuid.Parse("a0f9e860-f3fa-4994-a8d4-eb156dfce805")
	repo.EXPECT().DeleteTechInfoExternalInfra(gomock.Any(), gomock.Any()).Return(&expected, nil)

	svc := New(repo)
	Convey("Given company_id, user_id and techInfoExternalInfra details", t, func() {
		companyUUID, _ := uuid.Parse("a0f9e860-f3fa-4994-a8d4-eb156dfce805")
		userUUID, _ := uuid.Parse("64d1802e-1fa4-4732-b182-0719199108a8")
		techInfoExternalInfraUuid, _ := uuid.Parse("a0f9e860-f3fa-4994-a8d4-eb156dfce805")

		Convey("Call the DeleteTechInfoExternalInfra function", func() {
			err := svc.DeleteTechInfoExternalInfra(ctx, &companyUUID, &userUUID, &techInfoExternalInfraUuid)
			Convey("The value should be equal to the expected value", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}
