package pub

import "github.com/Sirupsen/logrus"

// Ctx ...
type Ctx struct {
	config      *Config
	awsHandler  awsHandlerIF
	lineHandler lineHandlerIF
}

// CtxSetting ...
type CtxSetting func(*Ctx) error

// AwsSetting ...
func AwsSetting(awsConfig *awsConfig, credential *Credential) CtxSetting {
	return func(ctx *Ctx) error {
		// FIXME
		return nil
	}
}

// NewCtx ...
func NewCtx(credential *Credential, config *Config) (*Ctx, error) {
	awsHandler, err := newAwsHandler(config.aws, credential)
	if err != nil {
		logrus.Errorf("AWS setting error %#v", err)
		return nil, err
	}
	logrus.Info("AWS connect setting done")

	lineHandler, err := newLineHandler(config.line, credential)
	if err != nil {
		logrus.Errorf("LINE setting error: %#v", err)
		return nil, err
	}
	logrus.Info("LINE connect setting done")

	return &Ctx{
		config:      config,
		awsHandler:  awsHandler,
		lineHandler: lineHandler,
	}, nil
}
