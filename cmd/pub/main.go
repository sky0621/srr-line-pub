package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"

	pub "github.com/sky0621/srr-line-pub"
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
	s, t := parseFlag()
	setConfig() // FIXME flagからtomlのパスを読む（無しなら「.」カレント）ようにする

	// FIXME ミドルウェアを検討
	mux := web.New()
	mux.Use(middleware.SubRouter)
	goji.Handle("srr/linebot/*", mux)
	mux.Post("/", pub.NewPubHandler(*s, *t))

	// FIXME うまく動かない・・・。Echoに変えるかな・・・。

	flag.Set("bind", viper.GetString("server.port"))
	goji.Serve()

	return exitCode
}

func parseFlag() (s *string, t *string) {
	s = flag.String("s", "ChannelSecret", "ChannelSecret")
	t = flag.String("t", "AccessToken", "AccessToken")
	flag.Parse()
	return
}

// FIXME Viperから構造体への変換タイミングは要検討
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
