package user

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/rmitsubayashi/bodyweight-server/src/repository"
)

type UserRepoImpl struct {
	conn *sqlx.DB
}

func (ur *UserRepoImpl) AddUser() error {
	statement := `
	INSERT INTO user (
		firebaseuid
	) VALUES (?)
	`
	result, err := ur.conn.Exec(statement, "firebaseuid")
	if err != nil {
		return err
	}

	fmt.Printf("returned user: %v", result)
	return nil
}

func NewUserRepo() (*UserRepoImpl, error) {
	conn, err := repository.NewDBConnection()
	if err != nil {
		return nil, err
	}
	return &UserRepoImpl{
		conn: conn,
	}, nil
}
