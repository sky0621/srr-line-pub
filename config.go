package pub

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config ...
type Config struct {
	server *serverConfig
	logger *loggerConfig
	aws    *awsConfig
	line   *lineConfig
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		server: newServerConfig(),
		logger: newLoggerConfig(),
		aws:    newAwsConfig(),
		line:   newLineConfig(),
	}
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

type loggerConfig struct {
	appName  string
	filepath string
	level    string
	server   *serverConfig
}

func newLoggerConfig() *loggerConfig {
	return &loggerConfig{
		appName:  viper.GetString("logger.app_name"),
		filepath: viper.GetString("logger.filepath"),
		level:    viper.GetString("logger.log_level"),
		server:   newServerConfig(),
	}
}

type awsConfig struct {
	sqs *sqsConfig
}

func newAwsConfig() *awsConfig {
	return &awsConfig{
		sqs: newSqsConfig(),
	}
}

type sqsConfig struct {
	region   string
	endpoint string
	name     string
}

func newSqsConfig() *sqsConfig {
	return &sqsConfig{
		region:   viper.GetString("aws.sqs.region"),
		endpoint: viper.GetString("aws.sqs.endpoint"),
		name:     viper.GetString("aws.sqs.name"),
	}
}

func (c *Config) String() string {
	return fmt.Sprintf(`
		server[host:%s, port:%s],
		logger[appName:%s, filepath:%s, level:%s],
		aws[sqs[region:%s, endpoint:%s, name:%s]],
		line[webhookURL:%s]
	`,
		c.server.host, c.server.port,
		c.logger.appName, c.logger.filepath, c.logger.level,
		c.aws.sqs.region, c.aws.sqs.endpoint, c.aws.sqs.name,
		c.line.webhookURL,
	)
}

type lineConfig struct {
	webhookURL string
}

func newLineConfig() *lineConfig {
	return &lineConfig{
		webhookURL: viper.GetString("line.webhook_url"),
	}
}

// ReadConfig ...
func ReadConfig(configFilePath string) error {
	viper.SetConfigFile(configFilePath)
	return viper.ReadInConfig()
}
