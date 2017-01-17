package pub

type ctx struct {
	logger      *appLogger
	config      *Config
	awsHandler  awsHandlerIF
	lineHandler lineHandlerIF
}
