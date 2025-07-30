package models

import "github.com/mailru/easyjson"

type Message struct{
	From 	string 	`json:"from"`
	To 		string	`json:"to"`	
	Msg 	string	`json:"msg"`
}

func CreateMsgFromJson(msg []byte) (*Message,error) {
	message := Message{}
	err := easyjson.Unmarshal(msg,&message)
	return &message,err
}

func(msg *Message) CreateJson() ([]byte,error) {
	return easyjson.Marshal(msg)
}
