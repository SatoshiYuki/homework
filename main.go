package main

import (
	"homework/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	initConfig()
	//fmt.Println(viper.GetString("MongoDB.Host"))
	//mongoDB.DbConnect()

	router := gin.Default()
	router.POST("/callback", services.RecieveMessage)
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
		"amount":  123,
		"status":  "ok",
		"message": "test123",
	})
}
