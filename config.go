package pub

import "github.com/spf13/viper"

type config struct {
	environment string
	appName     string
	server      *serverConfig
	logger      *loggerConfig
	aws         *awsConfig
	line        *lineConfig
}

func newConfig() *config {
	return &config{
		environment: viper.GetString("environment"),
		appName:     viper.GetString("app_name"),
		server:      newServerConfig(),
		logger:      newLoggerConfig(),
		aws:         newAwsConfig(),
		line:        newLineConfig(),
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

type sqsConfig struct {
	region   string `toml:"aws.sqs.region"`
	queueURL string `toml:"aws.sqs.queue_url"`
}

func newSqsConfig() *sqsConfig {
	return &sqsConfig{
		region:   viper.GetString("aws.sqs.region"),
		queueURL: viper.GetString("aws.sqs.queue_url"),
	}
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

func readConfig(configFilePath string) error {
	viper.SetConfigFile(configFilePath)
	return viper.ReadInConfig()
}
