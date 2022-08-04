initdb:
	docker run -itd --name mongo -p 27017:27017 mongo

config:
	echo {} > config.Json

mod:
	go mod init
	go mod tidy

goFile:
	echo package main > main.go
	mkdir mongoDB
	echo package mongoDB > mongoDB/mongoDB.go
	mkdir models
	echo package models > models/models.go