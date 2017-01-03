package pub

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type awsHandlerIF interface {
	getSqsClient() *sqs.SQS
}

type awsHandler struct {
	session   *session.Session
	sqsClient *sqs.SQS
}

func newAwsHandler(cfg *awsConfig) (awsHandlerIF, error) {
	if cfg.environment == constEnvLocal {
		return &awsHandlerMock{}, nil
	}
	sess, err := session.NewSession(&aws.Config{Region: aws.String(cfg.sqs.region)})
	if err != nil {
		return nil, err
	}
	cli := sqs.New(sess)

	return &awsHandler{
		session:   sess,
		sqsClient: cli,
	}, nil
}

func (h *awsHandler) getSqsClient() *sqs.SQS {
	return h.sqsClient
}
