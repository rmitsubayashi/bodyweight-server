package server

import (
	"time"
)

type Log struct {
	ID         int       `db:"id"`
	UserID     int       `db:"user_id"`
	Date       time.Time `db:"date"`
	CategoryID int       `db:"category_id"`
	Sets       []LogSet
	DeletedAt  *time.Time `db:"deleted_at"`
}

func (l *Log) SetSets(ls []LogSet) {
	l.Sets = ls
}
