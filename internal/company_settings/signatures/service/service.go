package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company"
	companyEntities "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_settings/signatures/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company_settings/signatures/repository"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/onboarding"
	onboardingEntities "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/onboarding/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/salesforce"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/user/userclient"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/cfg"
	appErr "github.com/nurdsoft/redesign-grp-trust-portal-api/shared/errors"
	"go.uber.org/zap"
)

type Service interface {
	GetSignatures(ctx context.Context, companyUUID, userUUID *uuid.UUID) ([]*entities.GetSignaturesResponse, error)
	UpdateStatus(ctx context.Context, companyUUID, userUUID, companySignatureUuid *uuid.UUID, req *entities.UpdateStatusRequestBody) (*entities.UpdateStatusResponse, error)
	ViewDocument(ctx context.Context, companyUUID, signatureUuid *uuid.UUID) (string, error)
	Webhook(ctx context.Context, data *entities.DocusignWebhookData) error
}

func New(repo repository.Repository, userClient userclient.Client, companyClient company.Client, onboardingClient onboarding.Client, salesforceClient salesforce.Client, commonConfig cfg.Config, logger *zap.SugaredLogger) Service {
	return &service{
		repo:             repo,
		userClient:       userClient,
		companyClient:    companyClient,
		onboardingClient: onboardingClient,
		sfClient:         salesforceClient,
		commonConfig:     commonConfig,
		logger:           logger,
	}
}

type service struct {
	repo             repository.Repository
	userClient       userclient.Client
	companyClient    company.Client
	onboardingClient onboarding.Client
	sfClient         salesforce.Client
	commonConfig     cfg.Config
	logger           *zap.SugaredLogger
}

func (s *service) Webhook(ctx context.Context, data *entities.DocusignWebhookData) error {
	if data.Event == "envelope-completed" {
		var companyUUID, signatureUUID, userUUID uuid.UUID
		var name, env string
		for _, t := range data.Data.EnvelopeSummary.CustomFields.TextCustomFields {
			switch t.Name {
			case "CompanyUUID":
				companyUUID = uuid.MustParse(t.Value)
			case "SignatureUUID":
				signatureUUID = uuid.MustParse(t.Value)
			case "UserUUID":
				userUUID = uuid.MustParse(t.Value)
			case "E":
				env = t.Value
			}
		}

		// return if the environment not matched with running environment
		if env != s.commonConfig.Env {
			return nil
		}

		for _, e := range data.Data.EnvelopeSummary.EnvelopeDocuments {
			// the documentID is created by the docusign for the document which needs to be signed.
			// Since we expect to add only one document to the docusign template,
			// the documentID expected to be 1 always for the signed document
			if e.DocumentID == "1" {
				name = e.Name
				break
			}
		}

		_, err := s.repo.CreateSignatures(ctx, &entities.CompanySignatures{
			SignatureUuid: signatureUUID,
			CompanyUuid:   companyUUID,
			Name:          name,
			Status:        "Signed",
			SignatureData: *data,
			CreatedBy:     userUUID,
			UpdatedBy:     userUUID,
		})
		if err != nil {
			return err
		}

		s.onboardingClient.UpdateOnboardingStatus(ctx, onboardingEntities.SignAuthorization, companyUUID, userUUID)
	}

	return nil
}

func (s *service) ViewDocument(ctx context.Context, companyUUID, signatureUuid *uuid.UUID) (string, error) {
	cs, err := s.repo.GetCompanySignature(ctx, companyUUID, signatureUuid)
	if err != nil {
		return "", err
	}
	var docBytes string
	for _, ed := range cs.SignatureData.Data.EnvelopeSummary.EnvelopeDocuments {
		// the documentID is created by the docusign for the document which needs to be signed.
		// Since we expect to add only one document to the docusign template,
		// the documentID expected to be 1 always for the signed document
		if ed.DocumentID == "1" {
			docBytes = ed.PDFBytes
			break
		}
	}

	if docBytes == "" {
		return "", &appErr.ErrNotFound{
			Message: "company signature not found",
		}
	}
	return docBytes, nil
}

func (s *service) UpdateStatus(ctx context.Context, companyUUID, userUUID, companySignatureUuid *uuid.UUID, req *entities.UpdateStatusRequestBody) (*entities.UpdateStatusResponse, error) {
	err := s.repo.UpdateStatus(ctx, userUUID, companySignatureUuid, req)
	if err != nil {
		return nil, err
	}

	return &entities.UpdateStatusResponse{
		CompanySignatureUuid: *companySignatureUuid,
		Status:               req.Status,
	}, nil

}

func (s *service) GetSignatures(ctx context.Context, companyUUID, userUUID *uuid.UUID) ([]*entities.GetSignaturesResponse, error) {
	signatures, err := s.repo.GetSignatures(ctx, companyUUID, userUUID)
	if err != nil {
		return nil, err
	}

	ctxUser, err := s.userClient.GetUserByUuid(ctx, *userUUID)
	if err != nil {
		return nil, err
	}

	company, err := s.companyClient.FindByUUID(ctx, &companyEntities.GetCompanyByIdRequest{
		CompanyUuid: *companyUUID,
	})
	if err != nil {
		return nil, err
	}

	username := ctxUser.FirstName + " " + ctxUser.LastName
	email := ctxUser.Email
	job := ctxUser.JobTitle
	clientCompanyName := username + ", " + company.Name

	vulSubName := "Vulnerability Management and Penetration Testing"
	contractNumber, err := s.getSFContractNumberByAccountID(ctx, company.ExternalId)
	if err != nil {
		return nil, err
	}

	for _, sig := range signatures {
		// example document URL for docusign:
		// https://www.docusign.net/Member/PowerFormSigning.aspx?PowerFormId=0324c1ba-514e-45b4-aa58-a272fa7a5806&env=na1&acct=7171320a-7100-4375-97c7-3266016165d9&v=2&Signer_UserName=Jerold Leslie&Signer_Email=jeroldleslie@gmail.com&EnvelopeField_CompanyUUID=a0f9e860-f3fa-4994-a8d4-eb156dfce805&EnvelopeField_SignatureUUID=cfd12bd3-6ae6-4847-aea3-c29452096d7b&EnvelopeField_UserUUID=64d1802e-1fa4-4732-b182-0719199108a8
		var docusignURL string
		if vulSubName != "" {
			docusignURL = fmt.Sprintf("%s&Signer_UserName=%s&Signer_Email=%s&SignerJob=%s&SFSubName=%s&CompanyName=%s&ClientCompanyName=%s&EnvelopeField_CompanyUUID=%s&EnvelopeField_SignatureUUID=%s&EnvelopeField_UserUUID=%s&ContractNumber=%s&EnvelopeField_E=%s",
				sig.DocumentUrl, username, email, job, vulSubName, company.Name, clientCompanyName, companyUUID.String(), sig.SignatureUuid.String(), userUUID.String(), contractNumber, s.commonConfig.Env)

			sig.DocumentViewURI = fmt.Sprintf("/companies/%s/signatures/%s/view", companyUUID.String(), sig.SignatureUuid.String())
		}
		sig.DocumentUrl = docusignURL
	}

	if err != nil {
		return nil, err
	}

	return signatures, nil
}

func (s *service) getSFContractNumberByAccountID(ctx context.Context, externalAccountId string) (string, error) {
	sfSess, err := s.sfClient.NewSession(ctx)
	if err != nil {
		return "", err
	}
	contracts, err := s.sfClient.GetContractsByAccountId(ctx, sfSess, externalAccountId)
	if err != nil {
		return "", err
	}

	for _, c := range contracts {
		if c.Type == "MSSP" || c.Type == "MSP" {
			return c.ContractNumber, nil
		}
	}

	return "", nil
}
