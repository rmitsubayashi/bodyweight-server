package exercise

import (
	"database/sql"

	"github.com/rmitsubayashi/bodyweight-server/src/model/server"
	"github.com/rmitsubayashi/bodyweight-server/src/repository"
)

type ExerciseRepoImpl struct {
	conn *sql.DB
}

func (er *ExerciseRepoImpl) GetExercise(exerciseID int) (*server.Exercise, error) {
	return nil, nil
}

func (er *ExerciseRepoImpl) FindMaxSingleSetValue(exerciseID int) (int, error) {
	return -1, nil
}

func (er *ExerciseRepoImpl) FindMaxTotalSetValue(exerciseID int) (int, error) {
	return -1, nil
}

func (er *ExerciseRepoImpl) FindUserExercises(userID int, categoryID int) (*[]server.UserExercise, error) {
	return nil, nil
}

func NewExerciseRepo() (*ExerciseRepoImpl, error) {
	conn, err := repository.NewDBConnection()
	if err != nil {
		return nil, err
	}
	return &ExerciseRepoImpl{
		conn: conn,
	}, nil
}
