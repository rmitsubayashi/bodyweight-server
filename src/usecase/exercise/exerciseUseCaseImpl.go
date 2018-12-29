package usecase

import (
	"github.com/rmitsubayashi/bodyweight-server/src/model/client"
)

type ExerciseUseCaseImpl struct {}

func (uc *ExerciseUseCaseImpl) GetExercisesByCategory(categoryID int) ([]client.Exercise, error) {
	return []client.Exercise {
		client.Exercise {
			ID: 1,
		},
	}, nil
}

func NewExerciseUseCase() *ExerciseUseCaseImpl {
	return &ExerciseUseCaseImpl{}
}
