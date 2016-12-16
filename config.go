package pub

import "github.com/BurntSushi/toml"

// config ... 設定全体
type config struct {
	AppName       string `toml:"app_name"`
	*serverConfig `toml:"server"`
	*logConfig    `toml:"logger"`
	*awsConfig    `toml:"aws"`
	*lineConfig   `toml:"line"`
}

// serverConfig ... サーバに関する設定
type serverConfig struct {
	Host string
	Port string
}

// logConfig ... ログに関する設定
type logConfig struct {
	Filepath string
	LogLevel string `toml:"log_level"`
}

// awsConfig ... AWS全般に関する設定
type awsConfig struct {
	awsAccessKeyID     string
	awsSecretAccessKey string
	*sqsConfig         `toml:"sqs"`
}

// sqsConfig ... SQSに関する設定
type sqsConfig struct {
	region   string
	queueURL string `toml:"queue_url"`
}

// lineConfig ... LINEに関する設定
type lineConfig struct {
	channelSecret string
	accessToken   string
}

// newConfig ... 設定の取込
func newConfig(arg *Arg) (*config, error) {
	var cfg config
	_, err := toml.DecodeFile(arg.configFilePath, &cfg)
	if err != nil {
		return nil, err
	}
	cfg.awsConfig.awsAccessKeyID = arg.awsAccessKeyID
	cfg.awsConfig.awsSecretAccessKey = arg.awsSecretAccessKey
	cfg.lineConfig = &lineConfig{
		channelSecret: arg.lineChannelSecret,
		accessToken:   arg.lineAccessToken,
	}
	return &cfg, nil
}
