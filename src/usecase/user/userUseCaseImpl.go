package user

import (
	"github.com/rmitsubayashi/bodyweight-server/src/model/client"
	ur "github.com/rmitsubayashi/bodyweight-server/src/repository/user"
)

type UserUseCaseImpl struct {
	userRepo ur.UserRepo
}

func (uuc *UserUseCaseImpl) GetUserInfo(userID int) (*client.User, error) {
	su, err := uuc.userRepo.GetUser(userID)
	if err != nil {
		return nil, err
	}
	cu := serverToClientUser(*su)
	return &cu, nil
}

func (uuc *UserUseCaseImpl) RegisterNewUser() error {
	return uuc.userRepo.AddUser()
}

func NewUserUseCase() (*UserUseCaseImpl, error) {
	r, err := ur.NewUserRepo()
	if err != nil {
		return nil, err
	}
	return &UserUseCaseImpl{
		userRepo: r,
	}, nil
}
