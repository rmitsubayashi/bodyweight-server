package transaction

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/rmitsubayashi/bodyweight-server/src/model/server"
	"github.com/rmitsubayashi/bodyweight-server/src/repository"
)

type TransactionRepoImpl struct {
	conn *sqlx.DB
}

func (tr *TransactionRepoImpl) AddTransaction(t server.Transaction) error {
	insertTransactionStatement := `
	INSERT INTO transaction (price, user_id)
	VALUES (?, ?)
	`
	r, err := tr.conn.Exec(insertTransactionStatement, t.Price, t.UserID)
	if err != nil {
		return err
	}

	id, err := r.LastInsertId()
	if err != nil {
		return err
	}
	for _, e := range t.Exercises {
		insertExerciseStatement := `
		INSERT INTO transaction_exercise (transaction_id, exercise_id, amount)
		VALUES (?, ?, ?)
		`
		if _, err := tr.conn.Exec(insertExerciseStatement, id, e.ExerciseID, e.Amount); err != nil {
			return err
		}
	}

	return nil
}

func (tr *TransactionRepoImpl) TransactionExists(t server.Transaction, date time.Time) (bool, error) {
	location, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		location = time.UTC
		log.Fatal(err)
	}
	localDateTime := date.In(location)
	dateString := localDateTime.Format("2006-01-02")
	selectStatement := `
	SELECT transaction_exercise.exercise_id, transaction_exercise.amount
	FROM transaction INNER JOIN transaction_exercise
	ON transaction.id = transaction_exercise.transaction_id
	WHERE transaction.price=? AND created_at LIKE ?
	`
	var te []server.TransactionExercise
	err = tr.conn.Select(&te, selectStatement, t.Price, dateString+"%")
	if err != nil {
		return false, err
	}
	matchedCt := 0
	for _, e := range t.Exercises {
		log.Printf("e: %+v\n", e)
		for _, compE := range te {
			log.Printf("compE: %v\n", compE)
			if e.ExerciseID == compE.ExerciseID && e.Amount == compE.Amount {
				matchedCt++
				break
			}
		}
	}
	log.Printf("matched: %d", matchedCt)
	return (matchedCt == len(t.Exercises)), nil
}

func NewTransactionRepo() (*TransactionRepoImpl, error) {
	conn, err := repository.NewDBConnection()
	if err != nil {
		return nil, err
	}
	return &TransactionRepoImpl{
		conn: conn,
	}, nil
}
