package service

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/dgrijalva/jwt-go"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/aws_client/cognito/config"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/aws_client/cognito/entities"
	cognitoErr "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/aws_client/cognito/errors"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/meta"
	"github.com/pkg/errors"
)

// Service stores business logic
type Service interface {
	VerifyToken(ctx context.Context, token string) (*entities.AWSCognitoIDTokenClaims, error)
	GetClaimsFromIDToken(ctx context.Context, idToken string) (*entities.AWSCognitoIDTokenClaims, error)
	InviteUser(ctx context.Context, details entities.InviteUserDetails, resend bool) error
	UpdateUserAttributes(ctx context.Context, details map[string]string) error
}

type service struct {
	config config.Config
}

func (s *service) VerifyToken(ctx context.Context, accessToken string) (*entities.AWSCognitoIDTokenClaims, error) {
	pubKeyURL := "https://cognito-idp.%s.amazonaws.com/%s/.well-known/jwks.json"
	formattedURL := fmt.Sprintf(pubKeyURL, s.config.AwsRegion, s.config.UserPoolID)

	c := jwk.NewCache(ctx)

	err := c.Register(formattedURL, jwk.WithRefreshInterval(15*time.Minute))
	if err != nil {
		return nil, err
	}

	keySet, err := c.Refresh(ctx, formattedURL)
	if err != nil {
		return nil, err
	}

	claims := &entities.AWSCognitoIDTokenClaims{}

	// JWT Parse - it's actually doing parsing, validation and returns back a token.
	// Use .Parse or .ParseWithClaims methods from https://github.com/dgrijalva/jwt-go
	token, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {

		// Verify if the token was signed with correct signing method
		// AWS Cognito is using RSA256
		_, ok := token.Method.(*jwt.SigningMethodRSA)
		if !ok {
			return nil, errors.Wrap(cognitoErr.ErrTokenInvalid, fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"]))
		}

		// Get "kid" value from token header
		// "kid" is shorthand for Key ID
		kid, ok := token.Header["kid"].(string)
		if !ok {
			return nil, errors.Wrap(cognitoErr.ErrTokenInvalid, "kid header not found")
		}

		// "kid" must be present in the public keys set
		key, ok := keySet.LookupKeyID(kid)
		if !ok {
			return nil, errors.Wrap(cognitoErr.ErrTokenInvalid, fmt.Sprintf("key %v not found", kid))
		}

		// In our case, we are returning only one key = keys[0]
		// Return token key as []byte{string} type
		var tokenKey interface{}
		if err := key.Raw(&tokenKey); err != nil {
			return nil, errors.Wrap(cognitoErr.ErrTokenInvalid, "failed to create token key")
		}

		return tokenKey, nil
	})

	if err != nil {
		// This place can throw expiration error
		return nil, cognitoErr.ErrTokenInvalid
	}

	// Check if token is valid
	if !token.Valid {
		return nil, cognitoErr.ErrTokenInvalid
	}

	return claims, nil
}

func (s *service) GetClaimsFromIDToken(ctx context.Context, idToken string) (*entities.AWSCognitoIDTokenClaims, error) {
	if idToken == "" {
		return nil, cognitoErr.ErrIdTokenInvalid
	}

	claims := &entities.AWSCognitoIDTokenClaims{}
	token, _, err := new(jwt.Parser).ParseUnverified(idToken, claims)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*entities.AWSCognitoIDTokenClaims)
	if !ok {
		return nil, errors.New("cannot extract claims")
	}

	return claims, nil
}

// Deprecated: We now use SES to send custom invites
func (s *service) InviteUser(ctx context.Context, details entities.InviteUserDetails, resend bool) error {
	var inp *cognitoidentityprovider.AdminCreateUserInput

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		// Provide SDK Config options, such as Region.
		Config: aws.Config{
			Region:                        aws.String(s.config.AwsRegion),
			MaxRetries:                    aws.Int(3),
			CredentialsChainVerboseErrors: aws.Bool(true),
		},
	}))

	cognitoClient := cognitoidentityprovider.New(sess)

	inp = &cognitoidentityprovider.AdminCreateUserInput{
		DesiredDeliveryMediums: aws.StringSlice([]string{"EMAIL"}),
		UserAttributes: []*cognitoidentityprovider.AttributeType{
			{
				Name:  aws.String("custom:first_name"),
				Value: aws.String(details.FirstName),
			},
			{
				Name:  aws.String("custom:last_name"),
				Value: aws.String(details.LastName),
			},
			{
				Name:  aws.String("email"),
				Value: aws.String(details.Email),
			},
		},
		UserPoolId: aws.String(s.config.UserPoolID),
		Username:   aws.String(details.Email),
		ValidationData: []*cognitoidentityprovider.AttributeType{
			{
				Name:  aws.String("company_name"),
				Value: aws.String(details.CompanyName),
			},
			{
				Name:  aws.String("industry_type"),
				Value: aws.String(details.IndustryType),
			},
			{
				Name:  aws.String("company_external_id"),
				Value: aws.String(details.CompanyExternalId),
			},
		},
	}

	if resend {
		inp.MessageAction = aws.String("RESEND")
	}

	_, err := cognitoClient.AdminCreateUser(inp)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) UpdateUserAttributes(ctx context.Context, details map[string]string) error {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region:                        aws.String(s.config.AwsRegion),
			MaxRetries:                    aws.Int(3),
			CredentialsChainVerboseErrors: aws.Bool(true),
		},
	}))

	accessToken := meta.AccessToken(ctx)

	cognitoClient := cognitoidentityprovider.New(sess)
	in := &cognitoidentityprovider.UpdateUserAttributesInput{
		AccessToken: &accessToken,
		UserAttributes: []*cognitoidentityprovider.AttributeType{
			{
				Name:  aws.String("custom:first_name"),
				Value: aws.String(details["first_name"]),
			},
			{
				Name:  aws.String("custom:last_name"),
				Value: aws.String(details["last_name"]),
			},
		},
	}
	_, err := cognitoClient.UpdateUserAttributes(in)

	if err != nil {
		return err
	}

	return nil
}

func New(config config.Config) *service {
	return &service{config: config}
}
