package service

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/ses/config"
)

const (
	CharSet = "UTF-8"
)

// Service stores business logic
type Service interface {
	SendEmail(ctx context.Context, subject, body, recipient string) error
}

type service struct {
	config config.Config
}

func New(config config.Config) *service {
	return &service{config: config}
}

func (s *service) SendEmail(ctx context.Context, subject, body, recipient string) error {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region:                        aws.String(s.config.AwsRegion),
			MaxRetries:                    aws.Int(3),
			CredentialsChainVerboseErrors: aws.Bool(true),
		},
	}))

	// Create an SES session.
	svc := ses.New(sess)

	// Assemble the email.
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{
				aws.String(recipient),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(body),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(CharSet),
				Data:    aws.String(subject),
			},
		},
		Source: aws.String(s.config.InviteSenderEmail),
	}

	result, err := svc.SendEmail(input)
	log.Println(result)

	if err != nil {
		return err
	}

	return nil
}
