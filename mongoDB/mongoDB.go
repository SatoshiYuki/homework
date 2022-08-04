package mongoDB

import "gopkg.in/mgo.v2"

func DbConnect() *mgo.Session {
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}

	return session
}
