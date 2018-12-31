package user

import (
	"github.com/rmitsubayashi/bodyweight-server/src/model/client"
)

type UserUseCase interface {
	GetUserInfo(userID int) (*client.Points, error)
	RegisterNewUser() error 
}