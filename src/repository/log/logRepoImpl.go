package log

import (
	"github.com/jmoiron/sqlx"

	"github.com/rmitsubayashi/bodyweight-server/src/model/server"
	"github.com/rmitsubayashi/bodyweight-server/src/repository"
)

type LogRepoImpl struct {
	conn *sqlx.DB
}

func (lr *LogRepoImpl) GetUserLogs(userID int) (*[]server.Log, error) {
	getLogsStatement := `
	SELECT * FROM user_log
	WHERE user_id=?
	`

	logs := []server.Log{}
	if err := lr.conn.Select(&logs, getLogsStatement, userID); err != nil {
		return nil, err
	}

	getLogSetsStatement := `
	SELECT * FROM log_set
	WHERE log_id=?
	`
	for _, log := range logs {
		var logSets []server.LogSet
		if err := lr.conn.Select(&logSets, getLogSetsStatement, log.ID); err != nil {
			return nil, err
		}
		log.SetSets(logSets)
	}

	return &logs, nil
}

func (lr *LogRepoImpl) GetUserLog(logID int) (*server.Log, error) {
	getLogStatement := `
	SELECT * FROM user_log
	WHERE id=?
	`
	var log server.Log
	if err := lr.conn.Get(&log, getLogStatement, logID); err != nil {
		return nil, err
	}

	getLogSetsStatement := `
	SELECT * FROM log_set
	WHERE log_id=?
	`
	var logSets []server.LogSet
	if err := lr.conn.Select(&logSets, getLogSetsStatement, logID); err != nil {
		return nil, err
	}
	log.SetSets(logSets)

	return &log, nil
}

func (lr *LogRepoImpl) AddLog(log server.Log) error {
	addLogStatement := `
	INSERT INTO user_log (user_id, category_id)
	VALUES (:user_id, :category_id)
	`
	result, err := lr.conn.NamedExec(addLogStatement, log)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	addLogSetsStatement := `
	INSERT INTO log_set (log_id, exercise_id, set_number, value)
	VALUES (:log_id, :exercise_id, :set_number, :value)
	`

	for _, logSet := range log.Sets {
		logSet.SetLogID(int(id))
		if _, err = lr.conn.NamedExec(addLogSetsStatement, logSet); err != nil {
			return err
		}

	}
	return nil
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
