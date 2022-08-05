package models

type Message struct {
	UserId  string `json:"user"`
	Message string `json:"message"`
}
