package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/line/line-bot-sdk-go/linebot"

	"github.com/aws/aws-sdk-go/service/sqs"

	pub "github.com/sky0621/srr-line-pub"
)

const (
	AWS_REGION = "ap-northeast-1" // アジアパシフィック (東京)
	QUEUE_URL  = "https://sqs.ap-northeast-1.amazonaws.com/065886101085/sri-line-message"
)

// まずは愚直に
func main() {
	// LINE-API
	ls := flag.String("ls", "ChannelSecret", "LINE ChannelSecret")
	lt := flag.String("lt", "AccessToken", "LINE AccessToken")
	// AWS
	ak := flag.String("ak", "accessKeyId", "AWS AdcessKeyId")
	as := flag.String("as", "secretAccessKey", "AWS SecretAccessKey")
	flag.Parse()
	arg := pub.NewArg(ls, lt, ak, as)

	e := echo.New()
	e.Use(middleware.Logger())

	e.POST("/", func(c echo.Context) error {
		bot, err := linebot.New(&ls, &lt)
		if err != nil {
			fmt.Println(err)
			return
		}
		events, err := bot.ParseRequest(c.Request())
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, event := range events {
			&sqs.
		}

		// FIXME events を処理する。
		// FIXME aws-sdk-go/sqs にてキューに投入

		return c.JSON(http.StatusOK, interface{})
	})
}
