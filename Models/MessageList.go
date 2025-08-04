package models

import "github.com/mailru/easyjson"

type MessageList struct{
	list []Message `json:"list"`
}


func CreateFromJson(data []byte) (MessageList,error){
	var res MessageList 
	if err := easyjson.Unmarshal(data,&res);err != nil{
		return res,err
	}
	return res,nil
}

func(msgList *MessageList) CreateJson() ([]byte,error) {
	return easyjson.Marshal(msgList)
}
