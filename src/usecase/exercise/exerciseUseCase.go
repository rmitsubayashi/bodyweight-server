package usecase

import (
	"github.com/rmitsubayashi/bodyweight-server/src/model/client"
)

type ExerciseUseCase interface {
	GetExercisesByCategory(categoryID int) ([]client.Exercise, error)
}