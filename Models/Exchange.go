package models

import (

	"github.com/DEEBBLUE/ExProtos/api/Types"
)

type(
	ExchangeStatus struct{
		Status       string
	}
	ExchangeCurrency struct{
		Currency 		 string
	}
	ExchangeData struct{
		Amount 			 float32	
		Details 		 string
		Currency 		 ExchangeCurrency
	}
	Exchange struct{
		ExchangeId 	 int
		ClientId 	 	 int
		OperId     	 int
		TimeStart  	 string
		TimeEnd 	 	 string
		Rate       	 float32

		DataIn 		 	 ExchangeData
		DataOut 		 ExchangeData

		Status 			 ExchangeStatus
	}
)

func(ex *Exchange) CreateFromGRPC(exchange *Types.Exchange) {
	var dataIn,dataOut ExchangeData
	var status ExchangeStatus

	dataIn.CreateFromGRPC(exchange.GetDataIn())
	dataOut.CreateFromGRPC(exchange.GetDataOut())

	status.CreateFromGRPC(exchange.GetStatus())

	ex.ExchangeId = int(exchange.GetExchangeId())
	ex.ClientId = int(exchange.GetClientId())
	ex.OperId = int(exchange.GetOperId())
	ex.TimeStart = exchange.GetTimeStart()
	ex.TimeEnd = exchange.GetTimeEnd()
	ex.Rate = exchange.GetRate()

	ex.DataIn = dataIn
	ex.DataOut = dataOut

	ex.Status = status

}

func(data *ExchangeData) CreateFromGRPC(dataGRPC *Types.ExchangeData){
	var currency ExchangeCurrency
	currency.CreateFromGRPC(dataGRPC.GetCurrency())

	data.Amount = dataGRPC.GetAmount()
	data.Details = dataGRPC.GetDetails()
	data.Currency = currency
}

func(curr *ExchangeCurrency) CreateFromGRPC(currGRPC Types.ExchangeCurrency){
	curr.Currency = currGRPC.String()
}

func(stat *ExchangeStatus) CreateFromGRPC(status Types.ExchangeStatus){
	stat.Status = status.String()
}

func(ex *Exchange) CreateGRPC() *Types.Exchange{
	return &Types.Exchange{
		ExchangeId: int32(ex.ExchangeId),		
		ClientId: int32(ex.ClientId),
		OperId: int32(ex.OperId),
		TimeStart: ex.TimeStart,
		TimeEnd: ex.TimeEnd,
		Rate: ex.Rate,

		DataIn: ex.DataIn.CreateGRPC(),
		DataOut: ex.DataOut.CreateGRPC(),
		
		Status: ex.Status.CreateGRPC(),
	}
}

func(data *ExchangeData) CreateGRPC() *Types.ExchangeData{
	return &Types.ExchangeData{
		Amount: data.Amount,
		Details: data.Details,
		Currency: data.Currency.CreateGRPC(),
	}
}

func(curr *ExchangeCurrency) CreateGRPC() Types.ExchangeCurrency{
	switch curr.Currency {
		case "BTC":
			return Types.ExchangeCurrency_BTC
		case "USDT":
			return Types.ExchangeCurrency_USDT
		default:
			return Types.ExchangeCurrency_FIAT
	}		
}

func(stat *ExchangeStatus) CreateGRPC() Types.ExchangeStatus{
	switch stat.Status {
		case "CREATED":
			return Types.ExchangeStatus_CREATED
		case "EXINPROCESSED":
			return Types.ExchangeStatus_EXINPROCESSED
		case "COMPELETED":
			return Types.ExchangeStatus_COMPELETED
		default:
			return Types.ExchangeStatus_CANCELED
	}
}
