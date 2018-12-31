package user

import (
	"github.com/rmitsubayashi/bodyweight-server/src/model/client"
)

type UserUseCaseImpl struct {}

func (*UserUseCaseImpl) GetUserInfo(userID int) (*client.Points, error) {
	return client.NewPoints(4, client.PointsAttributes{
		Points: 200,
	}), nil
}

func (*UserUseCaseImpl) RegisterNewUser() (error) {
	return nil
}

func NewUserUseCase() *UserUseCaseImpl {
	return &UserUseCaseImpl{}
}