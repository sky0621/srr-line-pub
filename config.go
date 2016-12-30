package pub

import (
	"fmt"

	"github.com/spf13/viper"
)

// ==================================================
// 外部モジュール用いてデコードするため、可視性は公開で。
// ==================================================

// Config ... 設定全体
type Config struct {
	Arg     *Arg
	AppName string `toml:"app_name"`
	Server  *ServerConfig
	Logger  *LoggerConfig
	Aws     *AwsConfig
	Line    *LineConfig
}

// ServerConfig ...
type ServerConfig struct {
	Host string `toml:"server.host"`
	Port string `toml:"server.port"`
}

// LoggerConfig ...
type LoggerConfig struct {
	Filepath string `toml:"logger.filepath"`
	Level    string `toml:"logger.log_level"`
}

// AwsConfig ...
type AwsConfig struct {
	Sqs *SqsConfig
}

// SqsConfig ...
type SqsConfig struct {
	Region   string `toml:"aws.sqs.region"`
	QueueURL string `toml:"aws.sqs.queue_url"`
}

// LineConfig ...
type LineConfig struct {
	WebhookUrl string `toml:"line.webhook_url"`
}

// newConfig ... 設定の取込
func newConfig(arg *Arg) *Config {
	viper.SetConfigFile(arg.configFilePath)
	err := viper.ReadInConfig()
	if err != nil { // 設定ファイルの読み取りエラー対応
		panic(fmt.Errorf("設定ファイル読み込みエラー: %s \n", err))
	}
	return &Config{
		Arg:     arg,
		AppName: viper.GetString("app_name"),
		Server: &ServerConfig{
			Host: viper.GetString("server.host"),
			Port: viper.GetString("server.port"),
		},
		Logger: &LoggerConfig{
			Filepath: viper.GetString("logger.filepath"),
			Level:    viper.GetString("logger.log_level"),
		},
		Aws: &AwsConfig{
			Sqs: &SqsConfig{
				Region:   viper.GetString("aws.sqs.region"),
				QueueURL: viper.GetString("aws.sqs.queue_url"),
			},
		},
		Line: &LineConfig{
			WebhookUrl: viper.GetString("line.webhook_url"),
		},
	}
}
