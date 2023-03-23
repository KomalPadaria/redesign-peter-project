package aws_client

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

func GetSession(awsRegion string, maxRetries int) *session.Session {
	return session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region:                        aws.String(awsRegion),
			MaxRetries:                    aws.Int(maxRetries),
			CredentialsChainVerboseErrors: aws.Bool(true),
		},
	}))
}
