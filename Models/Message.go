package models

import "github.com/mailru/easyjson"

type Message struct{
	ChatId 	string 	`json:"chat_id"`
	Msg 		string	`json:"msg"`
}

func CreateMsgFromJson(msg []byte) (*Message,error) {
	message := Message{}
	err := easyjson.Unmarshal(msg,&message)
	return &message,err
}

func(msg *Message) CreateJson() ([]byte,error) {
	return easyjson.Marshal(msg)
}
