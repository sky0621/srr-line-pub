package pub

import "github.com/aws/aws-sdk-go/aws/session"

type awsHandlerIF interface {
	getSqsHandler() sqsHandlerIF
}

type awsHandler struct {
	session    *session.Session
	sqsHandler sqsHandlerIF
	logger     *appLogger
}

func newAwsHandler(cfg *awsConfig, credential *Credential, logger *appLogger) (awsHandlerIF, error) {
	sqsHandler, err := newSqsHandler(cfg.sqs, credential, logger)
	if err != nil {
		return nil, err
	}

	return &awsHandler{
		sqsHandler: sqsHandler,
		logger:     logger,
	}, nil
}

func (h *awsHandler) getSqsHandler() sqsHandlerIF {
	return h.sqsHandler
}
