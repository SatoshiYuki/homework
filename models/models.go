package models

type Message struct {
	MessageId string `json:"messageId"`
	UserId    string `json:"user"`
	Message   string `json:"message"`
}
