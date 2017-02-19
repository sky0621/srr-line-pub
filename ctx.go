package pub

import (
	"github.com/sky0621/go-lib/log"
	"github.com/sky0621/srr-line-pub/global"
)

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

// CtxOption ...
type CtxOption func(ctx *Ctx) error

// NewCtx ...
func NewCtx(credential *Credential, config *Config) (*Ctx, error) {
	awsHandler, err := newAwsHandler(config.aws, credential)
	if err != nil {
		global.L.Logf(log.E, "AWS setting error %#v", err)
		return nil, err
	}
	global.L.Log(log.I, "AWS connect setting done")

	lineHandler, err := newLineHandler(config.line, credential)
	if err != nil {
		global.L.Logf(log.E, "LINE setting error: %#v", err)
		return nil, err
	}
	global.L.Log(log.I, "LINE connect setting done")

	return &Ctx{
		config:      config,
		awsHandler:  awsHandler,
		lineHandler: lineHandler,
	}, nil
}
