package models

import (
	"encoding/json"
	"io"
	"time"

	"github.com/DEEBBLUE/ExProtos/api/Types"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type(
	ExchangeType struct{
		Type string `json:"type"`
	}
	ExchangeStatus struct{
		Status       string						`json:"status"`
	}
	ExchangeCurrency struct{
		Currency 		 string						`json:"curryncy"`
	}
	ExchangeData struct{
		Amount 			 float32					`json:"amount"`
		Details 		 string						`json:"details"`
		Currency 		 ExchangeCurrency	`json:"exchnage_currency"`
	}
	Exchange struct{
		ExchangeId 	 int 						`json:"exchange_id"`
		ClientId 	 	 int 						`json:"client_id"`
		OperId     	 int						`json:"oper_id"`
		TimeStart  	 time.Time			`json:"time_start"`
		TimeEnd 	 	 time.Time			`json:"time_end"`
		Rate       	 float32				`json:"rate"`

		DataIn 		 	 ExchangeData		`json:"data_in"`
		DataOut 		 ExchangeData		`json:"data_out"`

		Status 			 ExchangeStatus `json:"exchnage_status"`
		Type 				 ExchangeType		`json:"exchange_type"`
	}
)

func(ex *Exchange) CreateFromJson(exchange io.ReadCloser) (error){
	return json.NewDecoder(exchange).Decode(&ex)
}

func(ex *Exchange) CreateJson() ([]byte,error){
	return json.Marshal(ex)
}

func(ex *Exchange) CreateFromGRPC(exchange *Types.Exchange) {
	var dataIn,dataOut ExchangeData
	var status ExchangeStatus
	var tp ExchangeType

	dataIn.CreateFromGRPC(exchange.GetDataIn())
	dataOut.CreateFromGRPC(exchange.GetDataOut())

	status.CreateFromGRPC(exchange.GetStatus())

	tp.CreateFromGRPC(exchange.GetType())

	ex.ExchangeId = int(exchange.GetExchangeId())
	ex.ClientId = int(exchange.GetClientId())
	ex.OperId = int(exchange.GetOperId())
	ex.TimeStart = exchange.GetTimeStart().AsTime()
	ex.TimeEnd = exchange.GetTimeEnd().AsTime()
	ex.Rate = exchange.GetRate()

	ex.DataIn = dataIn
	ex.DataOut = dataOut

	ex.Status = status
	ex.Type = tp 
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

func(t *ExchangeType) CreateFromGRPC(tp Types.ExchangeType){
	switch tp {
		case Types.ExchangeType_CR:
			t.Type = "CR"
		case 	Types.ExchangeType_RC:
			t.Type = "RC"
	}
}

func(ex *Exchange) CreateGRPC() *Types.Exchange{
	return &Types.Exchange{
		ExchangeId: int32(ex.ExchangeId),		
		ClientId: int32(ex.ClientId),
		OperId: int32(ex.OperId),
		TimeStart: timestamppb.New(ex.TimeStart),
		TimeEnd: timestamppb.New(ex.TimeEnd),
		Rate: ex.Rate,

		DataIn: ex.DataIn.CreateGRPC(),
		DataOut: ex.DataOut.CreateGRPC(),
		
		Status: ex.Status.CreateGRPC(),
		Type: ex.Type.CreateGRPC(),
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
		case "SBER":
			return Types.ExchangeCurrency_SBER
		case "ALFA":
			return Types.ExchangeCurrency_ALFA
		case "TINK":
			return Types.ExchangeCurrency_TINK
		case "GAZ":
			return Types.ExchangeCurrency_GAZ
		default:
			return Types.ExchangeCurrency_SBP
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

func(t *ExchangeType) CreateGRPC() Types.ExchangeType{
	switch t.Type {
		case "CR":
			return Types.ExchangeType_CR
		default:
			return Types.ExchangeType_RC
	}
}
