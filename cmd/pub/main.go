package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	// LINE-API
	ls := flag.String("ls", "ChannelSecret", "LINE ChannelSecret")
	lt := flag.String("lt", "AccessToken", "LINE AccessToken")
	// AWS
	ak := flag.String("ak", "accessKeyId", "AWS AdcessKeyId")
	as := flag.String("as", "secretAccessKey", "AWS SecretAccessKey")
	flag.Parse()

	e := echo.New()
	e.Use(middleware.Logger())

	e.POST("/", func(c echo.Context) error {
		bot, err := linebot.New(&s, &t)
		if err != nil {
			fmt.Println(err)
			return
		}
		events, err := bot.ParseRequest(c.Request())
		if err != nil {
			fmt.Println(err)
			return
		}

		// FIXME events を処理する。
		// FIXME aws-sdk-go/sqs にてキューに投入

		return c.JSON(http.StatusOK, interface{})
	})
}
