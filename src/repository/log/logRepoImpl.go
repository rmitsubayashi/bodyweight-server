package log

import (
	"github.com/rmitsubayashi/bodyweight-server/src/model/server"
	"github.com/rmitsubayashi/bodyweight-server/src/repository"
)

type LogRepoImpl struct {
	db: repository.DB
}

func (*LogRepoImpl) GetUserLogs(userID int) ([]server.Log, error) {

}

func NewLogRepo() (*LogRepoImpl, error) {
	db, err := repository.NewDBConnection()
	if err != nil {
		return nil, err
	}
	return &LogRepoImpl{
		db: db,
	}, nil
}
