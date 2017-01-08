package pub

type awsHandlerMock struct {
}

func (h *awsHandlerMock) getSqsHandler() sqsHandlerIF {
	return &sqsHandlerMock{}
}
