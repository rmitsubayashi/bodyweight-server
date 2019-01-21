package log

import (
	"database/sql"

	"github.com/rmitsubayashi/bodyweight-server/src/model/server"
	"github.com/rmitsubayashi/bodyweight-server/src/repository"
)

type LogRepoImpl struct {
	conn *sql.DB
}

func (*LogRepoImpl) GetUserLogs(userID int) (*[]server.Log, error) {
	return nil, nil
}

func NewLogRepo() (*LogRepoImpl, error) {
	conn, err := repository.NewDBConnection()
	if err != nil {
		return nil, err
	}
	return &LogRepoImpl{
		conn: conn,
	}, nil
}
