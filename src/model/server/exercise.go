package server

import (
	"time"
)

type Exercise struct {
	ID              int         `db:"id"`
	Title           string      `db:"title"`
	Description     string      `db:"description"`
	Level           int         `db:"level"`
	ImageURL        string      `db:"image_url"`
	MeasurementType string      `db:"measurement_type"`
	CategoryID      int         `db:"category_id"`
	IsDefault       bool        `db:"is_default"`
	TargetSets      []TargetSet `db:"target_set"`
	DeletedAt       *time.Time  `db:"deleted_at"`
}

func (e *Exercise) SetTargetSets(ts []TargetSet) {
	e.TargetSets = ts
}
