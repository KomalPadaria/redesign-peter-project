package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/salesforce/config"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/salesforce/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/salesforce/errors"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/transport/http/client"
	appError "github.com/nurdsoft/redesign-grp-trust-portal-api/shared/errors"
)

const (
	oauthEndpoint = "services/oauth2/token"
	logPrefix     = "salesforce"
)

type Service interface {
	NewSession(ctx context.Context) (*SFSession, error)
	GetAccountByID(ctx context.Context, session *SFSession, accountId string) (*entities.Account, error)
	GetContactByID(ctx context.Context, session *SFSession, contactId string) (*entities.Contact, error)
	GetSubscriptionsByID(ctx context.Context, session *SFSession, subscriptionsID string) (*entities.Subscription, error)
	GetSubscriptionsByIDs(ctx context.Context, session *SFSession, subscriptionsIDs []string) ([]entities.Subscription, error)
	GetSubscriptionsByAccountID(ctx context.Context, session *SFSession, accountId string) ([]entities.Subscription, error)
	GetContractsByAccountId(ctx context.Context, session *SFSession, accountId string) ([]entities.Contract, error)
	UpdateJiraEpicIdForAccount(ctx context.Context, session *SFSession, accountId, jiraEpicId string) (*entities.Account, error)
	GetConfig() config.Config
}

func New(cfg config.Config, httpClient *http.Client) Service {
	return &service{
		config:     cfg,
		httpClient: httpClient,
	}
}

type service struct {
	config     config.Config
	httpClient *http.Client
}

func (s *service) GetSubscriptionsByIDs(ctx context.Context, session *SFSession, subscriptionsIDs []string) ([]entities.Subscription, error) {
	subs, err := s.getSubscriptionsByIDs(ctx, session, subscriptionsIDs)
	if err != nil {
		log.Println(logPrefix, "get subscription by ids,", err)
		return nil, err
	}

	return subs, nil
}

func (s *service) GetConfig() config.Config {
	return s.config
}

func (s *service) GetSubscriptionsByID(ctx context.Context, session *SFSession, subscriptionsID string) (*entities.Subscription, error) {
	subs, err := s.getSubscriptionsByIDs(ctx, session, []string{subscriptionsID})
	if err != nil {
		log.Println(logPrefix, "get subscription by ids,", err)
		return nil, err
	}
	if len(subs) == 1 {
		return &subs[0], nil
	}
	return nil, nil
}

func (s *service) GetContractsByAccountId(ctx context.Context, session *SFSession, accountId string) ([]entities.Contract, error) {
	query := fmt.Sprintf("SELECT+Id,+ContractNumber,+Type__c,+Project__c,+Seats_Licensed__c,+Seats_Active__c,+StartDate,+EndDate+FROM+Contract+WHERE+AccountId+=+'%s'+AND+SBQQ__ActiveContract__c+=+1", accountId)

	url := fmt.Sprintf("%s/services/data/%s/query?q=%s",
		s.config.ApiHost, s.config.ApiVersion, query)

	data, err := session.httpRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
	}

	contractResponse := &entities.ContractResponse{}

	err = json.Unmarshal(data, contractResponse)
	if err != nil {
		log.Println(logPrefix, "contract response json decode failed,", err)
		return nil, err
	}

	return contractResponse.Contracts, nil
}

func (s *service) GetSubscriptionsByAccountID(ctx context.Context, session *SFSession, accountId string) ([]entities.Subscription, error) {
	subs, err := s.getSubscriptionsByAccountId(ctx, session, accountId)
	if err != nil {
		return nil, err
	}

	if len(subs) == 0 {
		return nil, &appError.ErrNotFound{Message: "subscriptions not found"}
	}

	return subs, nil
}

func (s *service) GetContactByID(ctx context.Context, session *SFSession, contactId string) (*entities.Contact, error) {
	url := s.makeUrl("Contact", contactId)

	data, err := session.httpRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println(logPrefix, "http request failed,", err)
		return nil, err
	}

	contact := &entities.Contact{}

	err = json.Unmarshal(data, contact)
	if err != nil {
		log.Println(logPrefix, "json decode failed,", err)
		return nil, err
	}

	return contact, nil
}

func (s *service) GetAccountByID(ctx context.Context, session *SFSession, accountId string) (*entities.Account, error) {
	url := s.makeUrl("Account", accountId)

	data, err := session.httpRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println(logPrefix, "http request failed,", err)
		return nil, err
	}

	account := &entities.Account{}

	err = json.Unmarshal(data, account)
	if err != nil {
		log.Println(logPrefix, "json decode failed,", err)
		return nil, err
	}

	return account, nil
}

func (s *service) UpdateJiraEpicIdForAccount(ctx context.Context, session *SFSession, accountId, jiraEpicId string) (*entities.Account, error) {
	url := s.makeUrl("Account", accountId)

	data := map[string]string{
		"Jira_Epic_Id__c": jiraEpicId,
	}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(data)
	if err != nil {
		log.Fatal(err)
	}

	_, err = session.httpRequest(http.MethodPatch, url, &buf)
	if err != nil {
		log.Println(logPrefix, "http request failed,", err)
		return nil, err
	}

	acc, err := s.GetAccountByID(ctx, session, accountId)
	if err != nil {
		log.Println(logPrefix, "get account", err)
		return nil, err
	}

	return acc, nil
}

func (s *service) NewSession(ctx context.Context) (*SFSession, error) {
	res, err := s.loginPassword(ctx)
	if err != nil {
		return nil, err
	}

	sess := &SFSession{}
	sess.AccessToken = res.AccessToken
	sess.InstanceURL = res.InstanceURL
	sess.ID = res.ID
	sess.TokenType = res.TokenType
	sess.IssuedAt = res.IssuedAt
	sess.Signature = res.Signature
	sess.httpClient = s.httpClient

	return sess, nil
}

func (s *service) getSubscriptionsByAccountId(ctx context.Context, session *SFSession, accountLookupId string) ([]entities.Subscription, error) {
	query := fmt.Sprintf("SELECT+Id,+Name,+SBQQ__ProductName__c,+SBQQ__Contract__c,+SBQQ__ProductId__c,+SBQQ__StartDate__c,+SBQQ__EndDate__c,"+
		"+SBQQ__Quantity__c,+SBQQ__ProductSubscriptionType__c,+SBQQ__OptionType__c,+SBQQ__OrderProduct__c,+Quantity_Billable_Hours__c,"+
		"+Quantity_In_Hour__c,+Quantity_In_Hour_Used__c,+SBQQ__Product__c,+SBQQ__Account__c,+SBQQ__TerminatedDate__c+"+
		"FROM+SBQQ__Subscription__c+"+
		"WHERE+SBQQ__Account__c+=+'%s'", accountLookupId)

	url := fmt.Sprintf("%s/services/data/%s/query?q=%s",
		s.config.ApiHost, s.config.ApiVersion, query)

	data, err := session.httpRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
	}

	sr := &entities.SubscriptionResponse{}

	err = json.Unmarshal(data, sr)
	if err != nil {
		log.Println(logPrefix, "subscription response json decode failed,", err)
		return nil, err
	}

	return sr.Subscriptions, nil
}

func (s *service) getSubscriptionsByIDs(ctx context.Context, session *SFSession, subscriptionsIDs []string) ([]entities.Subscription, error) {
	subIdsStr := "'" + strings.Join(subscriptionsIDs, "','") + "'"

	query := fmt.Sprintf("SELECT+Id,+Name,+SBQQ__ProductName__c,+SBQQ__Contract__c,+SBQQ__ProductId__c,+SBQQ__StartDate__c,+SBQQ__EndDate__c,"+
		"+SBQQ__Quantity__c,+SBQQ__ProductSubscriptionType__c,+SBQQ__OptionType__c,+SBQQ__OrderProduct__c,+Quantity_Billable_Hours__c,"+
		"+Quantity_In_Hour__c,+Quantity_In_Hour_Used__c,+SBQQ__Product__c,+SBQQ__Account__c,+SBQQ__TerminatedDate__c+"+
		"FROM+SBQQ__Subscription__c+"+
		"WHERE+Id+IN+(%s)", subIdsStr)

	url := fmt.Sprintf("%s/services/data/%s/query?q=%s",
		s.config.ApiHost, s.config.ApiVersion, query)

	data, err := session.httpRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
	}

	sr := &entities.SubscriptionResponse{}

	err = json.Unmarshal(data, sr)
	if err != nil {
		log.Println(logPrefix, "subscription response json decode failed,", err)
		return nil, err
	}

	if sr.TotalSize > 0 {
		return sr.Subscriptions, nil
	}

	return nil, nil
}

func (s *service) loginPassword(ctx context.Context) (*entities.LoginPasswordResponse, error) {
	httpClient := client.New(s.config.ApiHost, s.httpClient)

	res := &entities.LoginPasswordResponse{}

	url := fmt.Sprintf("%s?grant_type=password&client_id=%s&client_secret=%s&username=%s&password=%s",
		oauthEndpoint, s.config.ClientID, s.config.ClientSecret, s.config.Username, s.config.Password)

	err := httpClient.Post(ctx, url, nil, nil, res)

	if res.Error != "" {
		return nil, &errors.ErrSalesforceError{
			Message:      res.Error,
			HttpCode:     http.StatusUnauthorized,
			ErrorCode:    "",
			ErrorMessage: fmt.Sprintf("salesforce error:%s", res.ErrorDescription),
		}
	}
	return res, err
}

func (s *service) makeUrl(sObjectType, objectId string) string {
	url := fmt.Sprintf("%s/services/data/%s/sobjects/%s/%s",
		s.config.ApiHost, s.config.ApiVersion, sObjectType, objectId)

	return url
}
