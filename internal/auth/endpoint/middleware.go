// Package endpoint for auth
package endpoint

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-kit/kit/endpoint"
	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/auth/entity"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/auth/errors"
	authMeta "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/auth/meta"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/auth/permissions"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/aws_client/cognito"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/user/userclient"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/jwt"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/meta"
	"go.uber.org/zap"
)

// Middleware is a main middleware
type Middleware interface {
	SecureServiceWithCognitoEndpoint(featureName, method string) endpoint.Middleware
	SecureServiceWithRedesignEndpoint() endpoint.Middleware
	SecureServiceWithRedesignWebhookEndpoint() endpoint.Middleware
}

// New returns new middleware
func New(cognitoClient cognito.Client, userClient userclient.Client, logger *zap.SugaredLogger) Middleware {
	return &middleware{cognitoClient, userClient, logger}
}

type middleware struct {
	cognitoClient cognito.Client
	userClient    userclient.Client
	logger        *zap.SugaredLogger
}

func (s *middleware) SecureServiceWithRedesignWebhookEndpoint() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			token := meta.RedesignWebhookToken(ctx)
			err := jwt.VerifyWebhookToken(token)
			if err != nil {
				return nil, err
			}

			return next(ctx, req)
		}
	}
}

func (s *middleware) SecureServiceWithRedesignEndpoint() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			token := meta.RedesignToken(ctx)
			claims, err := jwt.VerifyToken(token)
			if err != nil {
				return nil, err
			}

			ctx = authMeta.WithUser(ctx, &entity.User{
				Username: claims.Username,
				Email:    claims.Email,
			})

			return next(ctx, req)
		}
	}
}

// SecureServiceWithCognitoEndpoint try to authorize by service cognito token
func (s *middleware) SecureServiceWithCognitoEndpoint(featureName, method string) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			token := meta.RawToken(ctx)
			claims, err := s.cognitoClient.VerifyToken(ctx, token)
			if err != nil {
				return nil, err
			}

			isFirstLogin, err := strconv.ParseBool(claims.RedesignUserIsFirstLogin)
			if err != nil {
				return nil, err
			}

			ctx = authMeta.WithUser(ctx, &entity.User{
				Uuid:         uuid.MustParse(claims.RedesignUserUserUuid),
				Username:     claims.RedesignUserUsername,
				FirstName:    claims.RedesignUserFirstName,
				LastName:     claims.RedesignUserLastName,
				Email:        claims.RedesignUserEmail,
				Phone:        claims.RedesignUserPhone,
				UserGroup:    claims.RedesignUserGroup,
				IsFirstLogin: isFirstLogin,
				Company: &entity.Company{
					Uuid:         uuid.MustParse(claims.RedesignCompanyCompanyUuid),
					Name:         claims.RedesignCompanyName,
					UserRole:     claims.RedesignCompanyUserRole,
					Type:         claims.RedesignCompanyType,
					IndustryType: claims.RedesignCompanyIndustryType,
					ExternalId:   claims.RedesignCompanyExternalId,
				},
			})

			userCompanyInfo, err := s.userClient.GetContextUserCompanyInfoInternal(ctx)
			if err != nil {
				return nil, err
			}

			permissionPath := fmt.Sprintf("%s_%s_%s", userCompanyInfo.Company.Type, userCompanyInfo.Group, userCompanyInfo.Company.UserRole)
			isAllowed := s.isAccessAllowed(featureName, permissionPath, method)
			if !isAllowed {
				return nil, errors.ErrNoPermission
			}

			return next(ctx, req)
		}
	}
}

func (s *middleware) isAccessAllowed(featureName, permissionPath, method string) bool {
	s.logger.Info(fmt.Sprintf("checking access with featureName=%s and permissionPath=%s", featureName, permissionPath))
	permission := permissions.Permissions[featureName][permissionPath]
	switch permission {
	case "ro":
		if method == "GET" {
			return true
		}
	case "rw":
		return true
	case "na":
		return false
	}

	return false
}

// func (s *middleware) isAccessAllowed(ctx context.Context, role string, group string, allowedRolesAndGroups []string) bool {
// 	if len(allowedRolesAndGroups) == 1 {
// 		if allowedRolesAndGroups[0] == "all" {
// 			return true
// 		}
// 	}

// 	for _, allowedRole := range allowedRolesAndGroups {
// 		if role == allowedRole {
// 			return true
// 		}

// 		if group == allowedRole {
// 			return true
// 		}
// 	}

// 	return false
// }
