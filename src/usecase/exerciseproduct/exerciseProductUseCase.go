package exerciseproduct

import (
	"github.com/rmitsubayashi/bodyweight-server/src/model/client"
)

type ExerciseProductUseCase interface {
	GetTodayExerciseProducts(userID int) (*[]client.ExerciseProduct, error)
	BuyExerciseProduct(userID int, ep client.ExerciseProduct) error
}