package transaction

import (
	"time"

	"github.com/rmitsubayashi/bodyweight-server/src/model/server"
)

type TransactionRepo interface {
	AddTransaction(server.Transaction) error
	TransactionExists(server.Transaction, time.Time) (bool, error)
}
