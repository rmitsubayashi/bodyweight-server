package server

type TargetSet struct {
	ExerciseID int `db:"exercise_id"`
	SetNumber  int `db:"set_number"`
	Value      int `db:"value"`
}
