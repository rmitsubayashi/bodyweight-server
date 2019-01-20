package user

import (
	"github.com/rmitsubayashi/bodyweight-server/src/repository"
)

type UserRepoImpl struct {
	db repository.DB
}

func (*UserRepoImpl) AddUser() error {

}

func NewUserRepo() (*UserRepoImpl, error) {
	db, err := repository.NewDBConnection(cfg)
	if err != nil {
		return nil, err
	}
	return &UserRepoImpl{
		db: db,
	}, nil
}
