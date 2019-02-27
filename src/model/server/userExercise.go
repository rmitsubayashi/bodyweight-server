package server

type UserExercise struct {
	ID         int `db:"id"`
	UserID     int `db:"user_id"`
	ExerciseID int `db:"exercise_id"`
	Amount     int `db:"amount"`
}
