package models

import "github.com/DEEBBLUE/ExProtos/api/Types"

type User struct{
	TgId 				int	
	ChatId 			int
	OwnerId 		int
	VerifStatus string
	Role 				string
	Balance 		int
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
