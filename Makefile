initdb:
	docker run -itd --name mongo -p 27017:27017 mongo

config:
	echo {} > config.Json

mod:
	go mod init
	go mod tidy

main:
	echo package main > main.go