package exercise

import (
	"github.com/rmitsubayashi/bodyweight-server/src/model/client"
)

type ExerciseUseCase interface {
	GetExerciseList(categoryID int) (*[]client.Exercise, error)
}
