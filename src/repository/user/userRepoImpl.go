package user

import (
	"database/sql"
	"fmt"

	"github.com/rmitsubayashi/bodyweight-server/src/repository"
)

type UserRepoImpl struct {
	conn *sql.DB
}

func (ur *UserRepoImpl) AddUser() error {
	statement := `
	INSERT INTO user (
		firebaseuid
	) VALUES (?)
	`
	preparedSt, err := ur.conn.Prepare(statement)
	if err != nil {
		return err
	}
	result, err := preparedSt.Exec("firebaseuid")
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
