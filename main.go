package main

import (
	"homework/mongoDB"

	"github.com/spf13/viper"
)

func main() {
	initConfig()
	//fmt.Println(viper.GetString("MongoDB.Host"))
	mongoDB.DbConnect()
}

func initConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	//viper.AddConfigPath("$GOPATH/src")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
