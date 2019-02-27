package exercise

import (
	"github.com/rmitsubayashi/bodyweight-server/src/model/client"
	er "github.com/rmitsubayashi/bodyweight-server/src/repository/exercise"
)

type ExerciseUseCaseImpl struct {
	exerciseRepo er.ExerciseRepo
}

func (uc *ExerciseUseCaseImpl) GetExerciseList(categoryID int) (*[]client.Exercise, error) {
	serverUE, eMap, err := uc.exerciseRepo.FindUserExercises(1, categoryID)
	if err != nil {
		return nil, err
	}
	clientUE := serverToClientUserExercises(*serverUE, eMap)
	return &clientUE, nil
}

func NewExerciseUseCase() (*ExerciseUseCaseImpl, error) {
	exerRepo, err := er.NewExerciseRepo()
	if err != nil {
		return nil, err
	}
	return &ExerciseUseCaseImpl{
		exerciseRepo: exerRepo,
	}, nil
}
