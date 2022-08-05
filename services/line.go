package services

import (
	"fmt"
	"homework/models"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/spf13/viper"
	"gopkg.in/mgo.v2/bson"
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
			}
		}
	}
}

func SendMessage(context *gin.Context) {

	client, err = linebot.New(viper.GetString("LINE.CHANNEL_SECRET"), viper.GetString("LINE.CHANNEL_ACCESS_TOKEN"))
	if err != nil {
		fmt.Println(err)
	}

	message := context.Query("message")

	_, err := client.PushMessage("U3e39106ee2de0697b17e6c27e641fe7f", linebot.NewTextMessage(message)).Do()
	if err != nil {
		panic(err)
	}
}

func QueryMessage(context *gin.Context) {
	c := DbConnect()
	var messages []models.Message
	err = c.Find(bson.M{"userid": viper.GetString("LINE.UserId")}).All(&messages)
	if err != nil {
		panic(err)
	}

	context.JSON(200, messages)
}
