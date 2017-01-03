package pub

type ctx struct {
	arg         *Arg
	logger      *appLogger
	config      *config
	awsHandler  awsHandlerIF
	lineHandler lineHandlerIF
}
