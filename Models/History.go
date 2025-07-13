package models

import "github.com/DEEBBLUE/ExProtos/api/Types"

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
