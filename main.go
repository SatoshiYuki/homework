package main

import (
	"homework/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	initConfig()

	router := gin.Default()
	router.POST("/homework/recieveMessage", services.RecieveMessage)
	router.POST("/homework/sendMessage", services.SendMessage)
	router.GET("/homework/queryMessage", services.QueryMessage)
	router.Run(":80")

}

func initConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func test(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "test123",
	})
}
