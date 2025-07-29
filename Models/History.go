package models

import (
	easyjson "github.com/mailru/easyjson"
	"github.com/DEEBBLUE/ExProtos/api/Types"
)

type UserHistory struct{
	History	[]Exchange
}

func(hist *UserHistory) CreateGRPC() []*Types.Exchange {
	var resHistory []*Types.Exchange
	for _,ex := range hist.History{
		resHistory = append(resHistory, ex.CreateGRPC())
	}		

	return resHistory
}
func(hist *UserHistory) CreateFromGRPC(history []*Types.Exchange){
	for _,ex := range history{
		var exchange Exchange
		exchange.CreateFromGRPC(ex)

		hist.History = append(hist.History, exchange)
	}
}

func(hist *UserHistory) CreateJson() ([]byte,error) {
	return easyjson.Marshal(hist)
}
