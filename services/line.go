package services

import (
	"fmt"
	"homework/models"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/spf13/viper"
)

var (
	client *linebot.Client
	err    error
)

func RecieveMessage(context *gin.Context) {

	client, err = linebot.New(viper.GetString("LINE.CHANNEL_SECRET"), viper.GetString("LINE.CHANNEL_ACCESS_TOKEN"))
	if err != nil {
		fmt.Println(err)
	}

	events, err := client.ParseRequest(context.Request)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			context.Writer.WriteHeader(400)
		} else {
			context.Writer.WriteHeader(500)
		}

		return
	}

	for _, event := range events {

		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:

				c := DbConnect()
				err = c.Insert(models.Message{UserId: event.Source.UserID, Message: message.Text})
				if err != nil {
					panic(err)
				}

				// if _, err = client.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
				// 	log.Println(err.Error())
				// }
			}
		}
	}
}
