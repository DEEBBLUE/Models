package models

import (

	"github.com/DEEBBLUE/ExProtos/api/Types"
	easyjson "github.com/mailru/easyjson"
)

type User struct{
	TgId 				int 		`json:"tg_id"`
	ChatId 			int			`json:"chat_id"`
	OwnerId 		int			`json:"owner_id"`
	VerifStatus string 	`json:"verif_status"`
	Role 				string	`json:"role"`
	Balance 		int			`json:"balance"`
}

func(user *User) CreateFromGRPC(userGRPC *Types.User) {
	user.TgId = int(userGRPC.GetTgId())
	user.ChatId = int(userGRPC.GetChatId())
	user.OwnerId = int(userGRPC.GetOwnerId())
	user.VerifStatus = userGRPC.GetVerifStatus().String()
	user.Role = userGRPC.GetRole().String()
	user.Balance = int(userGRPC.GetBalance())
}

func(user *User) CreateGRPC() *Types.User{
	var verifStatus Types.Verif
	var role Types.Role

	switch user.VerifStatus {
		case "INPROCESED":
			verifStatus = Types.Verif_INPROCESED	
		case "VERIFED":
			verifStatus = Types.Verif_VERIFED
		default:
			verifStatus = Types.Verif_UNVERIFED
	}
	switch user.Role{
		case "OPER":
			role = Types.Role_ROLE_OPER
		case "ADMIN":
			role = Types.Role_ROLE_ADMIN
		default:
			role = Types.Role_ROLE_CLIENT
	}

	return &Types.User{
		TgId: int32(user.TgId),
		ChatId: int32(user.ChatId),
		OwnerId: int32(user.OwnerId),
		Balance: int32(user.Balance),
		VerifStatus: verifStatus,
		Role: role,
	}	
}

func(user *User) CreateJson() ([]byte,error){
	return easyjson.Marshal(user)
}
func(user *User) CreateFromJson(usr []byte)(error){
	return easyjson.Unmarshal(usr,user)
}
