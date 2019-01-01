package user

import (
	"github.com/rmitsubayashi/bodyweight-server/src/model/client"
)

type UserUseCaseImpl struct {}

func (*UserUseCaseImpl) GetUserInfo(userID int) (*client.Points, error) {
	return &client.Points {
		ID: 4,
		Value: 200,
	}, nil
}

func (*UserUseCaseImpl) RegisterNewUser() (error) {
	return nil
}

func NewUserUseCase() *UserUseCaseImpl {
	return &UserUseCaseImpl{}
}