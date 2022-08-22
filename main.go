package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"homework/services"
)

func main() {
	initConfig()

	router := gin.Default()
	router.PUT("/lineBot/Message", services.RecieveMessage)
	router.POST("/lineBot/Message", services.SendMessage)
	router.GET("/lineBot/Message", services.QueryMessage)
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

//func test(context *gin.Context) {
//	context.JSON(http.StatusOK, gin.H{
//		"status":  "200",
//		"message": "test123",
//	})
//}
