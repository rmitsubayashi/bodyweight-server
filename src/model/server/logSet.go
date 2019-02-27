package server

import (
	"time"
)

type LogSet struct {
	ID         int        `db:"id"`
	LogID      int        `db:"log_id"`
	ExerciseID int        `db:"exercise_id"`
	SetNumber  int        `db:"set_number"`
	Value      int        `db:"value"`
	DeletedAt  *time.Time `db:"deleted_at"`
}

func (ls *LogSet) SetLogID(id int) {
	ls.LogID = id
}
