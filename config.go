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
	AppName            string `toml:"app_name"`
	ServerHost         string `toml:"server.host"`
	ServerPort         string `toml:"server.port"`
	LogFilepath        string `toml:"logger.filepath"`
	LogLevel           string `toml:"logger.log_level"`
	AwsAccessKeyID     string
	AwsSecretAccessKey string
	SqsRegion          string `toml:"aws.sqs.region"`
	SqsQueueURL        string `toml:"aws.sqs.queue_url"`
	LineChannelSecret  string
	LineAccessToken    string
}

// newConfig ... 設定の取込
func newConfig(arg *Arg) *Config {
	viper.SetConfigFile(arg.configFilePath)
	err := viper.ReadInConfig()
	if err != nil { // 設定ファイルの読み取りエラー対応
		panic(fmt.Errorf("設定ファイル読み込みエラー: %s \n", err))
	}
	return &Config{
		AppName:            viper.GetString("app_name"),
		ServerHost:         viper.GetString("server.host"),
		ServerPort:         viper.GetString("server.port"),
		LogFilepath:        viper.GetString("logger.filepath"),
		LogLevel:           viper.GetString("logger.log_level"),
		AwsAccessKeyID:     arg.awsAccessKeyID,
		AwsSecretAccessKey: arg.awsSecretAccessKey,
		SqsRegion:          viper.GetString("aws.sqs.region"),
		SqsQueueURL:        viper.GetString("aws.sqs.queue_url"),
		LineChannelSecret:  arg.lineChannelSecret,
		LineAccessToken:    arg.lineAccessToken,
	}
}
