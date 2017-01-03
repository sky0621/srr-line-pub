package pub

import "github.com/aws/aws-sdk-go/service/sqs"

type awsHandlerMock struct {
}

func (h *awsHandlerMock) getSqsClient() *sqs.SQS {
	return nil
}
