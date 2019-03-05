package user

import (
	"github.com/rmitsubayashi/bodyweight-server/src/model/server"
)

type UserRepo interface {
	AddUser() error
	GetUser(int) (*server.User, error)
	ChangePointsBy(userID int, amount int) error
}
