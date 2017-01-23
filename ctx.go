package pub

type ctx struct {
	credential  *Credential
	config      *Config
	awsHandler  awsHandlerIF
	lineHandler lineHandlerIF
}

func newCtx(credential *Credential, config *Config) (*ctx, error) {
	ctx := &ctx{
		credential: credential,
		config:     config,
	}

	return ctx, nil
}
