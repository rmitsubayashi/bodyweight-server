package server

type TransactionExercise struct {
	TransactionID int `db:"transaction_id"`
	ExerciseID    int `db:"exercise_id"`
	Amount        int `db:"amount"`
}
