package mongoDB

import (
	"fmt"
	"homework/models"

	"github.com/spf13/viper"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func DbConnect() *mgo.Collection {
	session, err := mgo.Dial(viper.GetString("MongoDB.Host"))
	if err != nil {
		panic(err)
	}

	db := session.DB(viper.GetString("MongoDB.DB"))
	c := db.C(viper.GetString("MongoDB.Collection"))

	// err = c.Insert(models.Message{MessageId: "1", UserId: "1", Message: "test123"})
	// if err != nil {
	// 	panic(err)
	// }

	result := []models.Message{}

	err = c.Find(bson.M{"messageid": "1"}).All(&result)
	if err != nil {
		panic(err)
	}
	for _, value := range result {
		fmt.Println(value)
	}

	return c
}
