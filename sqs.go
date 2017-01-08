package pub

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type sqsHandlerIF interface {
	sendMessage(body string) (*sqs.SendMessageOutput, error)
}

type sqsHandler struct {
	cfg    *sqsConfig
	cli    *sqs.SQS
	logger *appLogger
}

func newSqsHandler(cfg *sqsConfig, logger *appLogger) (sqsHandlerIF, error) {
	if cfg.environment == constEnvLocal {
		return &sqsHandlerMock{}, nil
	}

	// Credentialは環境変数セット済の前提
	awsCfg := &aws.Config{
		Region:   aws.String(cfg.region),
		Endpoint: aws.String(cfg.endpoint),
	}
	awsCfg.Credentials = credentials.NewEnvCredentials()

	sess, err := session.NewSession(awsCfg)
	if err != nil {
		return nil, err
	}

	cli := sqs.New(sess)

	return &sqsHandler{
		cfg:    cfg,
		cli:    cli,
		logger: logger,
	}, nil
}

func (h *sqsHandler) sendMessage(body string) (*sqs.SendMessageOutput, error) {
	getInput := &sqs.GetQueueUrlInput{QueueName: aws.String(h.cfg.name)}
	gquRes, err := h.cli.GetQueueUrl(getInput)
	if err != nil {
		h.logger.entry.Errorf("GetQueueUrl: %#v", err)
		return nil, err
	}

	input := &sqs.SendMessageInput{
		QueueUrl:    gquRes.QueueUrl,
		MessageBody: aws.String(body),
	}
	output, err := h.cli.SendMessage(input)
	if err != nil {
		h.logger.entry.Errorf("SendMessage: %#v", err)
		return nil, err
	}
	return output, nil
}
