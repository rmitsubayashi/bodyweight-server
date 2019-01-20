package exercise

import (
	"github.com/rmitsubayashi/bodyweight-server/src/model/server"
	"github.com/rmitsubayashi/bodyweight-server/src/repository"
)

type ExerciseRepoImpl struct{
	db: repository.DB
}

func (er *ExerciseRepoImpl) GetExercise(exerciseID int) (server.Exercise, error) {

}

func (er *ExerciseRepoImpl) FindMaxSingleSetValue(exerciseID int) (int, error) {
	return 0, nil
}

func (er *ExerciseRepoImpl) FindMaxTotalSetValue(exerciseID int) (int, error) {
	return 0, nil
}

func (er *ExerciseRepoImpl) FindUserExercises(userID int, categoryID int) ([]server.UserExercise, error) {

}

func NewExerciseRepo() (*ExerciseRepoImpl, error) {
	db, err := repository.NewDBConnection(cfg)
	if err != nil {
		return nil, err
	}
	return &ExerciseRepoImpl{
		db: db,
	}, nil
}
