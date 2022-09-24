package service

import (
	"demo/database"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/spf13/viper"
)

func Callback(c *gin.Context) {
	CHANNEL_SECRET := viper.GetString("linebot.secret")
	CHANNEL_TOKEN := viper.GetString("linebot.token")
	bot, err := linebot.New(CHANNEL_SECRET, CHANNEL_TOKEN)
	if err != nil {
		log.Fatal(err)
	}
	events, err := bot.ParseRequest(c.Request)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			fmt.Printf("err: %v\n", err)
		} else {
			fmt.Printf("err: %v\n", err)
		}
		return
	}
	for _, event := range events {
		fmt.Printf("event.Type: %v\n", event.Type)
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				//some call back condition here
				log.Print(event.Source.UserID)
				log.Print(message.Text)
				database.SaveMessage(event.Source.UserID, message.Text)
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	}
}
