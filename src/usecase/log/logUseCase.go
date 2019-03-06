package log

import (
	"github.com/rmitsubayashi/bodyweight-server/src/model/client"
)

type LogUseCase interface {
	GetLogList(userID int) ([]*client.Log, error)
	GetLogInfo(logID int) (*client.Log, error)
	RecordLog(log client.Log, uid int) (*client.Feedback, error)
}