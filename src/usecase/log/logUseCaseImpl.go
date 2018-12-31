package log

import (
	"github.com/rmitsubayashi/bodyweight-server/src/model/client"
)

type LogUseCaseImpl struct{}

func (uc *LogUseCaseImpl) GetLogList(userID int) ([]*client.Log, error) {
	return []*client.Log{
		client.NewLog(1, client.LogAttributes{}),
	}, nil
}

func (uc *LogUseCaseImpl) GetLogInfo(logID int) (*client.Log, error) {
	return client.NewLog(3, client.LogAttributes{}), nil
}

func (uc *LogUseCaseImpl) RecordLog(log client.Log) (*client.Feedback, error) {
	return client.NewFeedback(2, client.FeedbackAttributes{}), nil
}

func NewLogUseCase() *LogUseCaseImpl {
	return &LogUseCaseImpl{}
}
