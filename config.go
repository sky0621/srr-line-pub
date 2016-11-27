package srrlinepub

import "github.com/BurntSushi/toml"

// config ... 設定全体
type config struct {
	AppName       string `toml:"app_name"`
	*serverConfig `toml:"server"`
	*logConfig    `toml:"logger"`
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

// newConfig ... 設定の取込
func newConfig(configpath string) (*config, error) {
	var cfg config
	_, err := toml.DecodeFile(configpath, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
