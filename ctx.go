package pub

type ctx struct {
	config      *Config
	awsHandler  awsHandlerIF
	lineHandler lineHandlerIF
}
