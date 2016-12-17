package pub

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

// ==================================================
// 外部モジュール用いてデコードするため、可視性は公開で。
// ==================================================

// Config ... 設定全体
type Config struct {
	AppName      string        `toml:"app_name"`
	ServerConfig *ServerConfig `toml:"server"`
	LogConfig    *LogConfig    `toml:"logger"`
	AwsConfig    *AwsConfig    `toml:"aws"`
	LineConfig   *LineConfig   `toml:"line"`
}

// ServerConfig ... サーバに関する設定
type ServerConfig struct {
	Host string
	Port string
}

// LogConfig ... ログに関する設定
type LogConfig struct {
	Filepath string
	LogLevel string `toml:"log_level"`
}

// AwsConfig ... AWS全般に関する設定
type AwsConfig struct {
	AccessKeyID     string
	SecretAccessKey string
	SqsConfig       *SqsConfig `toml:"sqs"`
}

// SqsConfig ... SQSに関する設定
type SqsConfig struct {
	Region   string
	QueueURL string `toml:"queue_url"`
}

// LineConfig ... LINEに関する設定
type LineConfig struct {
	ChannelSecret string
	AccessToken   string
}

// newConfig ... 設定の取込
func newConfig(arg *Arg) (*Config, error) {
	var cfg Config
	_, err := toml.DecodeFile(arg.configFilePath, &cfg)
	if err != nil {
		return nil, err
	}
	fmt.Println(arg.configFilePath)
	cfg.AwsConfig = &AwsConfig{
		AccessKeyID:     arg.awsAccessKeyID,
		SecretAccessKey: arg.awsSecretAccessKey,
	}
	cfg.LineConfig = &LineConfig{
		ChannelSecret: arg.lineChannelSecret,
		AccessToken:   arg.lineAccessToken,
	}
	return &cfg, nil
}

// --------------------------------

func (c *Config) String() string {
	if c == nil {
		return "<nil>"
	}
	return fmt.Sprintf(
		`
		AppName: %s
		ServerConfig: %v
		LogConfig: %v
		AwsConfig: %v
		LineConfig: %v
		`,
		c.AppName,
		c.ServerConfig.String(),
		c.LogConfig.String(),
		c.AwsConfig.String(),
		c.LineConfig.String(),
	)
}

func (s *ServerConfig) String() string {
	if s == nil {
		return "<nil>"
	}
	return fmt.Sprintf(
		`\n
		  Host: %s
		  Port: %s
		`,
		s.Host,
		s.Port,
	)
}

func (l *LogConfig) String() string {
	if l == nil {
		return "<nil>"
	}
	return fmt.Sprintf(
		`\n
		  Filepath: %s
		  LogLevel: %s
		`,
		l.Filepath,
		l.LogLevel,
	)
}

func (a *AwsConfig) String() string {
	if a == nil {
		return "<nil>"
	}
	return fmt.Sprintf(
		`\n
		  AccessKeyID: %s
		  SecretAccessKey: %s
		  SqsConfig: %v
		`,
		a.AccessKeyID,
		a.SecretAccessKey,
		a.SqsConfig.String(),
	)
}

func (s *SqsConfig) String() string {
	if s == nil {
		return "<nil>"
	}
	return fmt.Sprintf(
		`\n
		    Region: %s
		    QueueURL: %s
		`,
		s.Region,
		s.QueueURL,
	)
}

func (l *LineConfig) String() string {
	if l == nil {
		return "<nil>"
	}
	return fmt.Sprintf(
		`\n
		  AccessToken: %s
		  ChannelSecret: %s
		`,
		l.AccessToken,
		l.ChannelSecret,
	)
}
