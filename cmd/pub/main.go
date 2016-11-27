package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/zenazn/goji"

	"github.com/fsnotify/fsnotify"
	pub "github.com/sky0621/srr-line-pub"
	"github.com/spf13/viper"
)

// flagでconfig.tomlのパスを取得
// config.tomlをViperでパース
// gojiでWebサーバ起動
// LINEからのメッセージならSQSに投入

func main() {
	os.Exit(realMain())
}

func realMain() (exitCode int) {
	// treat panic
	defer func() {
		err := recover()
		if err != nil {
			log.Printf("Panic occured. ERR: %+v", err)
			// FIXME 後始末

		}
	}()

	return wrappedMain()
}

func wrappedMain() (exitCode int) {
	s := flag.String("s", "ChannelSecret", "ChannelSecret")
	t := flag.String("t", "AccessToken", "AccessToken")
	flag.Parse()

	setConfig()

	h := &pub.PubHandler{
		ChannelSecret: *s,
		AccessToken:   *t,
	}

	goji.Get("srr/linebot/", h)
	goji.Serve()
	// http.ListenAndServe(":7171", mux)

	return exitCode
}

func setConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("設定ファイル読み込みエラー: %s \n", err))
	}

	viper.SetConfigType("toml")

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("設定ファイルが変更されました:", in.Name)
	})
}
