package service

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"html/template"
	"strings"
	"time"

	"github.com/google/uuid"
	authError "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/auth/errors"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/auth/meta"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/aws_client/cognito"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company"
	companyEntities "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/onboarding"
	onboardingEntities "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/onboarding/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/salesforce"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/ses"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/user/constants"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/user/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/user/repository"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/cfg"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/db"
	appErrors "github.com/nurdsoft/redesign-grp-trust-portal-api/shared/errors"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/nullable"
	"github.com/pkg/errors"
	"golang.org/x/exp/slices"
)

// Service stores business logic
type Service interface {
	CreateUser(ctx context.Context, CompanyUUID, UserUUID uuid.UUID, reqBody *entities.CreateUserRequestBody) (*entities.User, error)
	UpdateCompanyUserLink(ctx context.Context, companyUser *entities.CompanyUser) (*entities.CompanyUser, error)
	DeleteCompanyUserLink(ctx context.Context, companyUUID, reqUserUUID, userUUID *uuid.UUID) error
	CreateCompanyUser(ctx context.Context, req *entities.CreateCompanyAndUserRequest) (*entities.CreateCompanyAndUserResponse, error)
	GetContextUserCompanyInfo(ctx context.Context) (*entities.GetUserCompanyInfoByUserIdResponse, error)
	GetContextUserCompanyInfoInternal(ctx context.Context) (*entities.GetUserCompanyInfoByUserIdResponse, error)
	GetSFUserAndCompanyInfo(ctx context.Context, req *entities.GetSFUserAndCompanyInfoRequest) (*entities.GetSFUserAndCompanyInfoResponse, error)
	ActivateUserByEmail(ctx context.Context, req *entities.ActivateDeactivateUserRequest) (*entities.ActivateDeactivateUserResponse, error)
	GetUserByUsername(ctx context.Context, username string) (*entities.User, error)
	UpdateCurrentCompany(ctx context.Context, req *entities.UpdateCurrentCompanyRequest) error
	GetCompanyUsers(ctx context.Context, companyUUID, userUUID uuid.UUID) ([]*entities.User, error)
	ResendUserInvite(ctx context.Context, CompanyUUID, UserUUID uuid.UUID, reqBody *entities.CreateUserRequestBody) error
	UpdateUserDetails(ctx context.Context, reqBody *entities.UpdateUserRequest) (*entities.User, error)
	GetUserByUuid(ctx context.Context, userUUID uuid.UUID) (*entities.User, error)
	ListCompaniesForUser(ctx context.Context, userUUID uuid.UUID, keyword string) ([]*entities.GetCompaniesResponse, error)
}

type service struct {
	repo             repository.Repository
	companyClient    company.Client
	sfClient         salesforce.Client
	cognitoClient    cognito.Client
	onboardingClient onboarding.Client
	emailClient      ses.Client
	commonConfig     cfg.Config
}

// New service for user.
func New(
	repo repository.Repository,
	companyClient company.Client,
	sfClient salesforce.Client,
	cognitoClient cognito.Client,
	onboardingClient onboarding.Client,
	emailClient ses.Client,
	commonConfig cfg.Config,
) Service {
	svc := &service{
		repo:             repo,
		companyClient:    companyClient,
		sfClient:         sfClient,
		cognitoClient:    cognitoClient,
		onboardingClient: onboardingClient,
		emailClient:      emailClient,
		commonConfig:     commonConfig,
	}

	return svc
}

func (s *service) GetCompanyUsers(ctx context.Context, companyUUID, userUUID uuid.UUID) ([]*entities.User, error) {
	return s.repo.GetCompanyUsers(ctx, companyUUID, userUUID)
}

func (s *service) UpdateCurrentCompany(ctx context.Context, req *entities.UpdateCurrentCompanyRequest) error {
	err := s.repo.UpdateCurrentCompany(ctx, req.UserUUID, req.CompanyUUID)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) GetUserByUuid(ctx context.Context, userUUID uuid.UUID) (*entities.User, error) {
	return s.repo.FindByUUID(ctx, userUUID)
}

func (s *service) GetUserByUsername(ctx context.Context, username string) (*entities.User, error) {
	user, err := s.repo.FindByUsername(ctx, username)
	if err != nil {
		return nil, errors.Wrap(err, "finding user")
	}

	return user, nil
}

func (s *service) ActivateUserByEmail(ctx context.Context, req *entities.ActivateDeactivateUserRequest) (*entities.ActivateDeactivateUserResponse, error) {
	user, err := s.repo.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if user.Username == "" || user.FirstName == "" || user.LastName == "" {
		// update details in case they are null (portal invite flow)
		var updatedDetails map[string]interface{} = map[string]interface{}{
			"first_name": req.FirstName,
			"last_name":  req.LastName,
			"username":   req.Username,
		}
		err = s.repo.UpdateUserDetails(ctx, user.UserUuid, updatedDetails)
		if err != nil {
			return nil, err
		}
	}

	companiesOfUser, err := s.repo.GetCompanyUserByUserId(ctx, user.UserUuid)
	if err != nil {
		return nil, err
	}

	if len(companiesOfUser) > 0 {
		// set status in company_users for this user to "ACTIVE" for all companies
		err = s.repo.UpdateUserCompanyStatus(ctx, user.UserUuid, "ACTIVE")
		if err != nil {
			return nil, err
		}
	}

	_, err = s.repo.ActivateOrDeactivateByEmail(ctx, req.Email, true)
	if err != nil {
		return nil, err
	}

	return &entities.ActivateDeactivateUserResponse{
		Message: "user activated successfully",
	}, nil
}

func (s *service) CreateUser(ctx context.Context, companyUUID, userUUID uuid.UUID, reqBody *entities.CreateUserRequestBody) (*entities.User, error) {
	var createdUser *entities.User

	user := &entities.User{
		FirstName:    reqBody.FirstName,
		LastName:     reqBody.LastName,
		Email:        reqBody.Email,
		IsFirstLogin: true,
		CreatedBy:    userUUID,
		Group:        s.getUserGroupByRole(reqBody.Role),
	}

	createdUser, err := s.repo.Create(ctx, user)
	if err != nil {
		if !db.IsAlreadyExistError(err) {
			return nil, err
		}
	}

	if createdUser == nil {
		u, err := s.repo.FindByEmail(ctx, reqBody.Email)
		if err != nil {
			return nil, err
		}
		createdUser = &entities.User{
			UserUuid:           u.UserUuid,
			CurrentCompanyUuid: u.CurrentCompanyUuid,
			Username:           u.Username,
			FirstName:          u.FirstName,
			LastName:           u.LastName,
			Email:              u.Email,
			Phone:              u.Phone,
			IsFirstLogin:       u.IsFirstLogin,
			ExternalID:         u.ExternalID,
			CreatedAt:          u.CreatedAt,
			UpdatedAt:          u.UpdatedAt,
			CreatedBy:          u.CreatedBy,
			UpdatedBy:          u.UpdatedBy,
			CompanyRole:        u.CompanyRole,
			CompanyStatus:      u.CompanyStatus,
		}
	}

	createdUser.CompanyStatus = "PENDING"
	createdUser.CompanyRole = reqBody.Role

	// create company_user
	_, err = s.repo.LinkCompanyAndUser(ctx, &entities.CompanyUser{
		CompanyUuid: companyUUID,
		UserUuid:    createdUser.UserUuid,
		Role:        createdUser.CompanyRole,
		Status:      createdUser.CompanyStatus,
		CreatedBy:   userUUID,
	})
	if err != nil {
		if !repository.IsCompanyUserAlreadyExistsError(err) {
			return nil, err
		}
	}

	err = s.inviteUser(ctx, companyUUID, userUUID, reqBody)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (s *service) DeleteCompanyUserLink(ctx context.Context, companyUUID, reqUserUUID, userUUID *uuid.UUID) error {
	return s.repo.DeleteCompanyUserLink(ctx, companyUUID, reqUserUUID, userUUID)
}

func (s *service) UpdateCompanyUserLink(ctx context.Context, companyUser *entities.CompanyUser) (*entities.CompanyUser, error) {
	return s.repo.UpdateCompanyUserLink(ctx, companyUser)
}

func (s *service) CreateCompanyUser(ctx context.Context, req *entities.CreateCompanyAndUserRequest) (*entities.CreateCompanyAndUserResponse, error) {
	var userUUID, companyUUID uuid.UUID
	var rapid7_ids []string
	var res *companyEntities.FindCompanyByExternalIdResponse

	user, err := s.repo.FindByEmail(ctx, req.User.Email)
	if err != nil {
		if !repository.IsUserNotFoundError(err) {
			return nil, err
		}
	}

	req.User.Group = s.getUserGroupByRole(req.Company.Role)

	if user == nil {
		createdUser, err := s.repo.Create(ctx, &req.User)
		if err != nil {
			return nil, err
		}
		userUUID = createdUser.UserUuid
	} else {
		userUUID = user.UserUuid
		// phone no needs to be updated in case a user was invited from inside portal
		err := s.repo.UpdateUserDetails(ctx, user.UserUuid, map[string]interface{}{
			"phone": req.User.Phone,
		})
		if err != nil {
			return nil, err
		}

	}
	if req.Company.ExternalID != "" {
		// invite flow from portal may not have external id
		res, err = s.companyClient.FindByExternalId(ctx, &companyEntities.FindCompanyByExternalIdRequest{ExternalId: req.Company.ExternalID})
		if err != nil {
			return nil, err
		}

		sfSess, err := s.sfClient.NewSession(ctx)
		if err != nil {
			return nil, err
		}

		account, err := s.sfClient.GetAccountByID(ctx, sfSess, req.Company.ExternalID)
		if err != nil {
			return nil, err
		}

		for _, v := range strings.Split(account.Rapid7SiteIds, ",") {
			rapid7_ids = append(rapid7_ids, strings.TrimSpace(v))
		}

		if res.Company == nil {
			onboarding := companyEntities.OnboardingGroup{
				{
					Position:  1,
					Name:      onboardingEntities.InviteOtherUsersStep,
					URL:       "/company/information/manage-users",
					Status:    "DRAFT",
					UpdatedBy: userUUID.String(),
				},
				{
					Position:  2,
					Name:      onboardingEntities.ScheduleMeetingsStep,
					URL:       "/company/information/meetings",
					Status:    "DRAFT",
					UpdatedBy: userUUID.String(),
				},
				{
					Position:  3,
					Name:      onboardingEntities.CompanyInfoStep,
					URL:       "/company/information/locations",
					Status:    "DRAFT",
					UpdatedBy: userUUID.String(),
				},

				{
					Position:  4,
					Name:      onboardingEntities.SetTechnicalInformationStep,
					URL:       "/company/information/applications",
					Status:    "DRAFT",
					UpdatedBy: userUUID.String(),
				},
				{
					Position:  5,
					Name:      onboardingEntities.SignAuthorizationStep,
					URL:       "/company/information/authorizations",
					Status:    "DRAFT",
					UpdatedBy: userUUID.String(),
				},
				{
					Position:  6,
					Name:      onboardingEntities.UploadPoliciesAndProceduresStep,
					URL:       "/policies-procedures",
					Status:    "DRAFT",
					UpdatedBy: userUUID.String(),
				},
				{
					Position:  7,
					Name:      onboardingEntities.SecurityAwarenessTrainingStep,
					URL:       "/security-awareness",
					Status:    "DRAFT",
					UpdatedBy: userUUID.String(),
				},
			}

			req.Company.Onboarding = onboarding
			req.Company.Knowbe4Token = account.Knowbe4Token

			if len(rapid7_ids) > 0 {
				req.Company.Rapid7SiteIds = rapid7_ids
			}

			companyUUID, err = s.companyClient.CreateCompany(ctx, &req.Company)
			if err != nil {
				return nil, err
			}
		} else {
			companyUUID = res.Company.CompanyUuid
			// Update knowbe4_token, jira_epdi_id & rapid7 site_ids if the company already exists
			if account.Knowbe4Token != "" && len(rapid7_ids) > 0 {
				updateCompanyReq := companyEntities.UpdateCompanyRequest{
					CompanyUuid: companyUUID,
					Data: map[string]interface{}{
						"knowbe4_token":   account.Knowbe4Token,
						"rapid7_site_ids": rapid7_ids,
						"jira_epic_id":    account.JiraEpicId,
						"updated_at":      nullable.NewNullTime(time.Now().UTC()),
					},
				}

				err = s.companyClient.UpdateCompany(ctx, &updateCompanyReq)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	status := "ACTIVE"
	if req.User.ExternalID == "" {
		status = "PENDING"
	}

	// for user-invites via portal, company_id won't be passed
	if companyUUID != uuid.Nil {
		_, err = s.repo.LinkCompanyAndUser(ctx, &entities.CompanyUser{
			CompanyUuid: companyUUID,
			UserUuid:    userUUID,
			Role:        "admin",
			Status:      status,
		})

		if err != nil {
			if !repository.IsCompanyUserAlreadyExistsError(err) {
				return nil, err
			}
		}
	}

	companyAndUserRes := &entities.CreateCompanyAndUserResponse{
		CompanyUuid: companyUUID,
		UserUuid:    userUUID,
	}

	return companyAndUserRes, nil
}

func (s *service) GetContextUserCompanyInfoInternal(ctx context.Context) (*entities.GetUserCompanyInfoByUserIdResponse, error) {
	return s.getUserCompanyInfo(ctx, false)
}

func (s *service) GetContextUserCompanyInfo(ctx context.Context) (*entities.GetUserCompanyInfoByUserIdResponse, error) {
	return s.getUserCompanyInfo(ctx, true)

}

func (s *service) GetSFUserAndCompanyInfo(ctx context.Context, req *entities.GetSFUserAndCompanyInfoRequest) (*entities.GetSFUserAndCompanyInfoResponse, error) {
	splitted := strings.Split(req.Token, "?")

	contactId := splitted[0]
	accountId := splitted[1]

	contactIdDecoded, err := base64.StdEncoding.DecodeString(contactId)
	if err != nil {
		return nil, err
	}

	accountIdDecoded, err := base64.StdEncoding.DecodeString(accountId)
	if err != nil {
		return nil, err
	}

	sfSess, err := s.sfClient.NewSession(ctx)
	if err != nil {
		return nil, err
	}

	contact, err := s.sfClient.GetContactByID(ctx, sfSess, string(contactIdDecoded))
	if err != nil {
		return nil, err
	}

	account, err := s.sfClient.GetAccountByID(ctx, sfSess, string(accountIdDecoded))
	if err != nil {
		return nil, err
	}

	industry := "Entertainment"
	if account.Industry != "" {
		industry = account.Industry
	}

	res := &entities.GetSFUserAndCompanyInfoResponse{
		User: entities.SFUserResponse{
			Username:   contact.Email,
			FirstName:  contact.FirstName,
			LastName:   contact.LastName,
			Email:      contact.Email,
			Phone:      contact.Phone,
			ExternalID: contact.ID,
		},
		Company: []entities.SFCompanyResponse{
			{
				Name:         account.Name,
				IndustryType: []string{industry},
				Role:         "admin",
				ExternalID:   account.ID,
			},
		},
	}

	return res, nil
}

func (s *service) getUserInfoResponse(ref *entities.User, companies []*companyEntities.Company, currentCompanyUuid string) *entities.GetUserCompanyInfoByUserIdResponse {

	var comp companyEntities.Company
	var isMfaAppEnabled, isMfaSmsEnabled bool

	for _, c := range companies {
		if c.CompanyUuid.String() == currentCompanyUuid {
			comp = *c
		}
		c.Onboarding = nil
		c.Address = nil
	}

	if slices.Contains(ref.Mfa, "app") {
		isMfaAppEnabled = true
	}
	if slices.Contains(ref.Mfa, "sms") {
		isMfaSmsEnabled = true
	}

	return &entities.GetUserCompanyInfoByUserIdResponse{
		UserUuid:        ref.UserUuid,
		Username:        ref.Username,
		FirstName:       ref.FirstName,
		LastName:        ref.LastName,
		Email:           ref.Email,
		Phone:           ref.Phone,
		IsFirstLogin:    ref.IsFirstLogin,
		Company:         &comp,
		Companies:       companies,
		Status:          ref.CompanyStatus,
		Group:           ref.Group,
		Job:             ref.JobTitle,
		IsMfaSmsEnabled: isMfaSmsEnabled,
		IsMfaAppEnabled: isMfaAppEnabled,
	}
}

func (s *service) getUserCompanyInfo(ctx context.Context, updateIsFirstLogin bool) (*entities.GetUserCompanyInfoByUserIdResponse, error) {
	ctxUser := meta.User(ctx)
	var companies []*companyEntities.Company

	if ctxUser == nil {
		return nil, errors.New("user not found in the context")
	}

	user, err := s.repo.FindByEmail(ctx, ctxUser.Email)
	if err != nil {
		return nil, errors.Wrap(err, "finding user")
	}

	if user == nil {
		return nil, &appErrors.ErrNotFound{Message: "User not found"}
	}

	res, err := s.companyClient.GetUserCompaniesByUserUuid(ctx, &companyEntities.GetCompaniesByUserIdRequest{
		UserUuid: user.UserUuid,
	})
	if err != nil {
		return nil, errors.Wrap(err, "finding companies")
	}

	if user.CurrentCompanyUuid == uuid.Nil {
		companyUuid := res.Companies[0].CompanyUuid
		if err = s.repo.UpdateCurrentCompany(ctx, user.UserUuid, companyUuid); err != nil {
			return nil, err
		}
		user.CurrentCompanyUuid = companyUuid
	}

	if updateIsFirstLogin {
		_, err = s.repo.UpdateIsFirstLoginByUserUuid(ctx, user.UserUuid, false)
		if err != nil {
			return nil, errors.Wrap(err, "updating is_first_login in user")
		}
	}

	companyUser, err := s.repo.GetCompanyUser(ctx, user.UserUuid, user.CurrentCompanyUuid)
	if err != nil {
		// bypass company user check for superadmin, engineer, csc user group
		if !repository.IsUserNotFoundError(err) || (user.Group != constants.UserGroupSuperadmin && user.Group != constants.UserGroupEngineer && user.Group != constants.UserGroupCsc) {
			return nil, err
		}
	}

	if companyUser != nil {
		user.CompanyStatus = companyUser.Status
	}

	if user.Group == constants.UserGroupSuperadmin || user.Group == constants.UserGroupEngineer || user.Group == constants.UserGroupCsc {
		allCompanies, err := s.companyClient.GetAllCompanies(ctx, &companyEntities.GetAllCompaniesRequest{})
		if err != nil {
			return nil, err
		}

		// update role of superadmin IFF they are a part of any company
		for _, c := range allCompanies {
			for _, uc := range res.Companies {
				if uc.CompanyUuid == c.CompanyUuid {
					c.UserRole = uc.UserRole
				} else {
					switch user.Group {
					case constants.UserGroupSuperadmin:
						c.UserRole = "superadmin"
					case constants.UserGroupEngineer:
						c.UserRole = "engineer"
					case constants.UserGroupCsc:
						c.UserRole = "csc"
					default:
						c.UserRole = "superadmin"
					}
				}
			}
		}
		companies = allCompanies
	} else {
		companies = res.Companies
	}

	return s.getUserInfoResponse(user, companies, user.CurrentCompanyUuid.String()), nil
}

func (s *service) ResendUserInvite(ctx context.Context, companyUUID, UserUUID uuid.UUID, reqBody *entities.CreateUserRequestBody) error {
	err := s.inviteUser(ctx, companyUUID, UserUUID, reqBody)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) inviteUser(ctx context.Context, companyUUID, UserUUID uuid.UUID, reqBody *entities.CreateUserRequestBody) error {
	user := meta.User(ctx)

	if reqBody.Role == "csc" || reqBody.Role == "engineer" {
		if user.Company.UserRole != "superadmin" {
			return authError.ErrNoPermission
		}
	}

	companyIdEncoded := base64.StdEncoding.EncodeToString([]byte(companyUUID.String()))
	userEmailEncoded := base64.StdEncoding.EncodeToString([]byte(reqBody.Email))

	signupToken := fmt.Sprintf("%s?%s", companyIdEncoded, userEmailEncoded)
	subject := "Your invitation to [re]DESIGN"

	// E.g. domain/user-signup/comapnyId/userEmail
	signupUrl := fmt.Sprintf("https://%s/%s/%s", s.commonConfig.FrontendDomain, "user-signup", signupToken)

	email := `<html><head> <style>p{line-height: 22px;}.signup{width: 240px; height: 32px; background: #436AF3; border-radius: 4.16px; text-decoration: none; margin-top: 40px; margin-bottom: 40px; font-size: 16px; padding-top: 6.5px; padding-left: 61.5px; padding-right: 61.5px; text-align: center; display: block;}.footer-heading{padding-top: 74px; font-weight: 600; font-size: 16px; text-align: center; color: #464646;}.footer-sub-heading{font-size: 12px; text-align: center; color: #B7B7B7;}</style></head><body leftmargin="0" rightmargin="0" topmargin="0" bottommargin="0" marginwidth="0" marginheight="0"> <table width="600" cellpadding="0" cellspacing="0" align="center"> <tr> <td> <img alt="REDESIGN Logo" title="REDESIGN Logo" style="display:block" height="37px" src="https://redesigntrustportal-static-files.s3.us-west-2.amazonaws.com/rdt_logo_small.png"> <p>Welcome to REDESIGN Trust Portal!</p><p>This is the first step to accelerating your Trust subscription journey and Managing Risk, Vulnerabilities, Security Control Policies, Procedures and much more.</p><p>Please follow the on-screen prompts to complete your account creation, and let your Trust journey begin.</p><p>If you have any challenges or issues logging in, please connect with your Trust support team via email at <a href="mailto:security.services@redesign-group.com" target="_blank">security.services@redesign-group.com</a></p><p><a class="signup" style="color: #FFFFFF" href="{{.signupLink}}">Create Account</a></p><p>Thank you,</p><p><a href="mailto:tsullivan@redesign-group.com">Tim Sullivan</a></p></td></tr><tr> <td> <footer class="footer-heading">Welcome to REDESIGN Trust Portal</footer> <footer class="footer-sub-heading">Trust is the only way!</footer> </td></tr></table></body></html>`
	templ := template.Must(template.New("invite").Parse(email))

	var tpl bytes.Buffer
	_ = templ.Execute(&tpl, map[string]interface{}{
		"signupLink": signupUrl,
	})

	err := s.emailClient.SendEmail(ctx, subject, tpl.String(), reqBody.Email)
	if err != nil {
		return err
	}

	s.onboardingClient.UpdateOnboardingStatus(ctx, onboardingEntities.InviteOtherUsers, companyUUID, UserUUID)
	return nil
}

func (s *service) UpdateUserDetails(ctx context.Context, reqBody *entities.UpdateUserRequest) (*entities.User, error) {
	ctxUser := meta.User(ctx)
	if ctxUser == nil {
		return nil, errors.New("user not found in the context")
	}

	details := map[string]interface{}{
		"first_name": reqBody.FirstName,
		"last_name":  reqBody.LastName,
		"phone":      reqBody.Phone,
		"job":        reqBody.JobTitle,
		"mfa":        reqBody.MfaMethods,
	}

	err := s.cognitoClient.UpdateUserAttributes(ctx, map[string]string{
		"first_name": reqBody.FirstName,
		"last_name":  reqBody.LastName,
	})
	if err != nil {
		return nil, err
	}

	err = s.repo.UpdateUserDetails(ctx, ctxUser.Uuid, details)
	if err != nil {
		return nil, err
	}

	user, err := s.repo.FindByUUID(ctx, ctxUser.Uuid)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *service) ListCompaniesForUser(ctx context.Context, userUUID uuid.UUID, keyword string) ([]*entities.GetCompaniesResponse, error) {
	var companiesByPermission []*entities.GetCompaniesResponse
	var allCompanies, userCompanies, companies []*companyEntities.Company

	user, err := s.repo.FindByUUID(ctx, userUUID)
	if err != nil {
		return nil, err
	}

	// list companies that a user is a part of
	res, err := s.companyClient.GetUserCompaniesByUserUuid(ctx, &companyEntities.GetCompaniesByUserIdRequest{
		UserUuid: user.UserUuid,
		Keyword:  keyword,
	})
	if err != nil {
		return nil, err
	}

	if res != nil {
		userCompanies = res.Companies
	}

	if user.Group == constants.UserGroupSuperadmin || user.Group == constants.UserGroupEngineer || user.Group == constants.UserGroupCsc {
		// if superadmin return access to all companies
		allCompanies, err = s.companyClient.GetAllCompanies(ctx, &companyEntities.GetAllCompaniesRequest{
			Keyword: keyword,
		})

		if err != nil {
			return nil, err
		}

		// update role of superadmin IFF they are a part of any company
		for _, c := range allCompanies {
			for _, uc := range userCompanies {
				if uc.CompanyUuid == c.CompanyUuid {
					c.UserRole = uc.UserRole
				} else {
					switch user.Group {
					case constants.UserGroupSuperadmin:
						c.UserRole = "superadmin"
					case constants.UserGroupEngineer:
						c.UserRole = "engineer"
					case constants.UserGroupCsc:
						c.UserRole = "csc"
					default:
						c.UserRole = "superadmin"
					}
				}
			}
		}
		companies = allCompanies
	} else {
		companies = userCompanies
	}

	for _, company := range companies {
		c := &entities.GetCompaniesResponse{
			CompanyUuid: company.CompanyUuid,
			CompanyName: company.Name,
			CompanyRole: company.UserRole,
			UserGroup:   user.Group,
		}

		companiesByPermission = append(companiesByPermission, c)
	}

	return companiesByPermission, nil
}

func (s *service) getUserGroupByRole(role string) string {
	switch role {
	case "admin":
		return "customer"
	case "user":
		return "customer"
	case "csc":
		return "csc"
	case "engineer":
		return "engineer"
	case "superadmin":
		return "superadmin"
	}
	return "customer"
}
