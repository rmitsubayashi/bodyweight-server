package usecase

import (
	"github.com/rmitsubayashi/bodyweight-server/src/model/client"
)

type ExerciseUseCaseImpl struct {}

func (uc *ExerciseUseCaseImpl) GetExerciseList(categoryID int) ([]*client.Exercise, error) {
	return []*client.Exercise{
		client.NewExercise(categoryID, client.ExerciseAttributes{}),
	}, nil
}

func NewExerciseUseCase() *ExerciseUseCaseImpl {
	return &ExerciseUseCaseImpl{}
}
