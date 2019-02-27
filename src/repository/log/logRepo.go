package log

import (
	"github.com/rmitsubayashi/bodyweight-server/src/model/server"
)

type LogRepo interface {
	GetUserLogs(userID int) (*[]server.Log, error)
	GetUserLog(logID int) (*server.Log, error)
	AddLog(log server.Log) error
}
