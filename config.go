package pub

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config ...
type Config struct {
	environment string
	appName     string
	server      *serverConfig
	logger      *loggerConfig
	aws         *awsConfig
	line        *lineConfig
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		environment: viper.GetString("environment"),
		appName:     viper.GetString("app_name"),
		server:      newServerConfig(),
		logger:      newLoggerConfig(),
		aws:         newAwsConfig(),
		line:        newLineConfig(),
	}
}

func (c *Config) String() string {
	return fmt.Sprintf("environment: %s, appName: %s, server: %s, logger: %s, aws: %s, line: %s", c.environment, c.appName, c.server.String(), c.logger.String(), c.aws.String(), c.line.String())
}

type serverConfig struct {
	host string
	port string
}

func newServerConfig() *serverConfig {
	return &serverConfig{
		host: viper.GetString("server.host"),
		port: viper.GetString("server.port"),
	}
}

func (c *serverConfig) String() string {
	return fmt.Sprintf("host: %s, port: %s", c.host, c.port)
}

type loggerConfig struct {
	environment string
	appName     string
	filepath    string
	level       string
	server      *serverConfig
}

func newLoggerConfig() *loggerConfig {
	return &loggerConfig{
		environment: viper.GetString("environment"),
		appName:     viper.GetString("app_name"),
		filepath:    viper.GetString("logger.filepath"),
		level:       viper.GetString("logger.log_level"),
		server:      newServerConfig(),
	}
}

func (c *loggerConfig) String() string {
	return fmt.Sprintf("environment: %s, appName: %s, filepath: %s, level: %s, server: %s", c.environment, c.appName, c.filepath, c.level, c.server.String())
}

type awsConfig struct {
	environment string
	sqs         *sqsConfig
}

func newAwsConfig() *awsConfig {
	return &awsConfig{
		environment: viper.GetString("environment"),
		sqs:         newSqsConfig(),
	}
}

func (c *awsConfig) String() string {
	return fmt.Sprintf("environment: %s, sqs: %s", c.environment, c.sqs.String())
}

type sqsConfig struct {
	environment string
	region      string
	endpoint    string
	name        string
}

func newSqsConfig() *sqsConfig {
	return &sqsConfig{
		environment: viper.GetString("environment"),
		region:      viper.GetString("aws.sqs.region"),
		endpoint:    viper.GetString("aws.sqs.endpoint"),
		name:        viper.GetString("aws.sqs.name"),
	}
}

func (c *sqsConfig) String() string {
	return fmt.Sprintf("environment: %s, region: %s, endpoint: %s, name: %s", c.environment, c.region, c.endpoint, c.name)
}

type lineConfig struct {
	environment string
	webhookURL  string
}

func newLineConfig() *lineConfig {
	return &lineConfig{
		environment: viper.GetString("environment"),
		webhookURL:  viper.GetString("line.webhook_url"),
	}
}

func (c *lineConfig) String() string {
	return fmt.Sprintf("environment: %s, webhookURL: %s", c.environment, c.webhookURL)
}

// ReadConfig ...
func ReadConfig(configFilePath string) error {
	viper.SetConfigFile(configFilePath)
	return viper.ReadInConfig()
}
