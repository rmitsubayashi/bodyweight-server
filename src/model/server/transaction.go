package server

type Transaction struct {
	ID        int                   `db:"id"`
	Price     int                   `db:"price"`
	CreatedAt string                `db:"created_at"`
	UserID    int                   `db:"user_id"`
	Exercises []TransactionExercise `db:"transaction_exercise"`
}
