package pub

import (
	"github.com/Sirupsen/logrus"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type sqsHandlerIF interface {
	sendMessage(body string) (*sqs.SendMessageOutput, error)
}

type sqsHandler struct {
	cfg *sqsConfig
	cli *sqs.SQS
}

func newSqsHandler(cfg *sqsConfig, credential *Credential) (sqsHandlerIF, error) {
	if cfg.environment == constEnvLocal {
		return &sqsHandlerMock{}, nil
	}

	awsCfg := &aws.Config{
		Credentials: credentials.NewStaticCredentials(credential.AwsAccessKeyID, credential.AwsSecretAccessKey, ""),
		Region:      aws.String(cfg.region),
		Endpoint:    aws.String(cfg.endpoint),
	}

	sess, err := session.NewSession(awsCfg)
	if err != nil {
		return nil, err
	}

	cli := sqs.New(sess)

	return &sqsHandler{
		cfg: cfg,
		cli: cli,
	}, nil
}

func (h *sqsHandler) sendMessage(body string) (*sqs.SendMessageOutput, error) {
	getInput := &sqs.GetQueueUrlInput{QueueName: aws.String(h.cfg.name)}
	gquRes, err := h.cli.GetQueueUrl(getInput)
	if err != nil {
		logrus.Errorf("GetQueueUrl: %#v", err)
		return nil, err
	}

	input := &sqs.SendMessageInput{
		QueueUrl:    gquRes.QueueUrl,
		MessageBody: aws.String(body),
	}
	output, err := h.cli.SendMessage(input)
	if err != nil {
		logrus.Errorf("SendMessage: %#v", err)
		return nil, err
	}
	return output, nil
}
