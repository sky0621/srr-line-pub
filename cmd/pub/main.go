package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	goji "goji.io"

	"goji.io/pat"

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
	s, t := parseFlag()
	setConfig() // FIXME flagからtomlのパスを読む（無しなら「.」カレント）ようにする
	// FIXME ミドルウェアを検討
	mux := goji.NewMux()
	mux.HandleC(pat.Post("/srr/linebot/"), pub.NewPubHandler(*s, *t))
	http.ListenAndServe(":7171", mux) // FIXME tomlから読む

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
