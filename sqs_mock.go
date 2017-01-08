package pub

import "github.com/aws/aws-sdk-go/service/sqs"

type sqsHandlerMock struct {
}

func (h *sqsHandlerMock) sendMessage(body string) (*sqs.SendMessageOutput, error) {
	return nil, nil
}
