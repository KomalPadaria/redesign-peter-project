package service

import (
	"context"
	"testing"
	"time"

	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/nullable"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_settings/websites/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_settings/websites/repository"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_service_CreateWebsite(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := repository.NewMockRepository(ctrl)
	ctx := context.Background()

	companyUUID := uuid.New()
	userUUID := uuid.New()
	websiteUUID := uuid.New()

	expected := &entities.CompanyWebsite{
		CompanyWebsiteUuid: websiteUUID,
		CompanyUuid:        companyUUID,
		Url:                "www.test.com",
		IndustryType:       "Entertainment",
		Zip:                "123123",
		Country:            "US",
		State:              "ST",
		City:               "testcity",
		Address1:           "testaddress1",
		Address2:           "",
		CreatedAt:          nullable.NewNullTime(time.Now()),
		CreatedBy:          userUUID,
	}
	repo.EXPECT().CreateWebsite(gomock.Any(), gomock.Any()).Return(expected, nil)

	svc := New(repo)
	Convey("Given company_id, user_id and website details", t, func() {
		website := &entities.CompanyWebsite{
			Url:          "www.test.com",
			IndustryType: "Entertainment",
			Zip:          "123123",
			Country:      "US",
			State:        "ST",
			City:         "testcity",
			Address1:     "testaddress1",
			Address2:     "",
		}
		Convey("Call the CreateWebsite function", func() {
			actual, err := svc.CreateWebsite(ctx, &companyUUID, &userUUID, website)
			Convey("The value should be equal to the expected value", func() {
				So(err, ShouldBeNil)
				So(actual, ShouldNotBeNil)
				So(actual, ShouldResemble, expected)
			})
		})
	})
}

func Test_service_UpdateWebsite(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := repository.NewMockRepository(ctrl)
	ctx := context.Background()

	companyUUID := uuid.New()
	userUUID := uuid.New()
	websiteUUID := uuid.New()

	expected := &entities.CompanyWebsite{
		CompanyWebsiteUuid: websiteUUID,
		CompanyUuid:        companyUUID,
		Url:                "www.test.com",
		IndustryType:       "Entertainment",
		Zip:                "123123",
		Country:            "US",
		State:              "ST",
		City:               "testcity",
		Address1:           "testaddress1",
		Address2:           "",
		CreatedAt:          nullable.NewNullTime(time.Now()),
		UpdatedAt:          nullable.NewNullTime(time.Now()),
		CreatedBy:          userUUID,
		UpdatedBy:          userUUID,
	}

	repo.EXPECT().UpdateWebsite(gomock.Any(), gomock.Any()).Return(expected, nil)

	svc := New(repo)
	Convey("Given company_id, user_id, application_id and website details", t, func() {

		website := &entities.CompanyWebsite{}
		Convey("Call the UpdateWebsite function", func() {
			actual, err := svc.UpdateWebsite(ctx, &companyUUID, &userUUID, &websiteUUID, website)
			Convey("The value should be equal to the expected value", func() {
				So(err, ShouldBeNil)
				So(actual, ShouldNotBeNil)
				So(actual.CreatedAt, ShouldResemble, expected.CreatedAt)
			})
		})
	})
}

func Test_service_GetWebsiteById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	companyUUID, _ := uuid.Parse("a0f9e860-f3fa-4994-a8d4-eb156dfce805")
	userUUID, _ := uuid.Parse("64d1802e-1fa4-4732-b182-0719199108a8")
	websiteUUID, _ := uuid.Parse("a0f9e860-f3fa-4994-a8d4-eb156dfce805")

	repo := repository.NewMockRepository(ctrl)
	ctx := context.Background()
	expected := &entities.CompanyWebsite{
		CompanyWebsiteUuid: websiteUUID,
		CompanyUuid:        companyUUID,
		Url:                "test.com",
		IndustryType:       "Entertainment",
		Zip:                "11758",
		Country:            "US",
		State:              "ND",
		City:               "Massapequa",
		Address1:           "4267 Merrick Rd",
		Address2:           "",
		CreatedAt:          nullable.NewNullTime(time.Now()),
		UpdatedAt:          nullable.NewNullTime(time.Now()),
		CreatedBy:          userUUID,
		UpdatedBy:          userUUID,
	}

	repo.EXPECT().GetWebsiteById(gomock.Any(), gomock.Any()).Return(expected, nil)

	svc := New(repo)
	Convey("Given company_id, user_id and application_id", t, func() {
		companyUUID, _ = uuid.Parse("a0f9e860-f3fa-4994-a8d4-eb156dfce805")
		userUUID, _ = uuid.Parse("64d1802e-1fa4-4732-b182-0719199108a8")
		websiteUUID, _ = uuid.Parse("a0f9e860-f3fa-4994-a8d4-eb156dfce805")

		Convey("Call the GetWebsiteById function", func() {
			actual, err := svc.GetWebsiteById(ctx, &companyUUID, &userUUID, &websiteUUID)
			Convey("The value should be equal to the expected value", func() {
				So(err, ShouldBeNil)
				So(actual, ShouldNotBeNil)
				So(actual, ShouldEqual, expected)
			})
		})
	})
}

func Test_service_GetAllWebsites(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userUuid := uuid.New()

	repo := repository.NewMockRepository(ctrl)
	ctx := context.Background()
	expected := make([]*entities.CompanyWebsite, 0)
	expected = append(expected, &entities.CompanyWebsite{
		CompanyWebsiteUuid: uuid.New(),
		CompanyUuid:        uuid.New(),
		Url:                "test.com",
		IndustryType:       "Entertainment",
		Zip:                "11758",
		Country:            "US",
		State:              "ND",
		City:               "Massapequa",
		Address1:           "4267 Merrick Rd",
		Address2:           "",
		CreatedAt:          nullable.NewNullTime(time.Now()),
		UpdatedAt:          nullable.NewNullTime(time.Now()),
		CreatedBy:          userUuid,
		UpdatedBy:          userUuid,
	})

	repo.EXPECT().GetAllWebsites(gomock.Any(), gomock.Any()).Return(expected, nil)

	svc := New(repo)
	Convey("Given company_id, user_id", t, func() {
		companyUUID, _ := uuid.Parse("a0f9e860-f3fa-4994-a8d4-eb156dfce805")
		userUUID, _ := uuid.Parse("64d1802e-1fa4-4732-b182-0719199108a8")

		Convey("Call the GetAllWebsites function", func() {
			actual, err := svc.GetAllWebsites(ctx, &companyUUID, &userUUID)
			Convey("The value should be equal to the expected value", func() {
				So(err, ShouldBeNil)
				So(actual, ShouldNotBeNil)
				So(actual, ShouldResemble, expected)
			})
		})
	})
}

func Test_service_DeleteWebsite(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := repository.NewMockRepository(ctrl)
	ctx := context.Background()
	expected, _ := uuid.Parse("a0f9e860-f3fa-4994-a8d4-eb156dfce805")
	repo.EXPECT().DeleteWebsite(gomock.Any(), gomock.Any()).Return(&expected, nil)

	svc := New(repo)
	Convey("Given company_id, user_id and website details", t, func() {
		companyUUID, _ := uuid.Parse("a0f9e860-f3fa-4994-a8d4-eb156dfce805")
		userUUID, _ := uuid.Parse("64d1802e-1fa4-4732-b182-0719199108a8")
		websiteUUID, _ := uuid.Parse("a0f9e860-f3fa-4994-a8d4-eb156dfce805")

		Convey("Call the DeleteWebsite function", func() {
			err := svc.DeleteWebsite(ctx, &companyUUID, &userUUID, &websiteUUID)
			Convey("The value should be equal to the expected value", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}
