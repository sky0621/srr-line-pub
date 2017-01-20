package pub

import "github.com/aws/aws-sdk-go/aws/session"

type awsHandlerIF interface {
	getSqsHandler() sqsHandlerIF
}

type awsHandler struct {
	session    *session.Session
	sqsHandler sqsHandlerIF
}

func newAwsHandler(cfg *awsConfig, credential *Credential) (awsHandlerIF, error) {
	sqsHandler, err := newSqsHandler(cfg.sqs, credential)
	if err != nil {
		return nil, err
	}

	return &awsHandler{sqsHandler: sqsHandler}, nil
}

func (h *awsHandler) getSqsHandler() sqsHandlerIF {
	return h.sqsHandler
}
