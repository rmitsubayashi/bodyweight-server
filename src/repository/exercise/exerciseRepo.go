package exercise

import (
	"github.com/rmitsubayashi/bodyweight-server/src/model/server"
)

type ExerciseRepo interface {
	GetExercise(exerciseID int) (*server.Exercise, error)
	FindMaxSingleSetValue(exerciseID int) (int, error)
	FindMaxTotalSetValue(exerciseID int) (int, error)
	FindUserExercises(userID int, categoryID int) (*[]server.UserExercise, map[int]server.Exercise, error)
	FindRandomExercise(categoryID int, minLevel int, maxLevel int, seed int, count int) (*[]server.Exercise, error)
	AddUserExercise(exercise *server.UserExercise) error
	RemoveUserExercise(uid int, exerciseID int, amount int) error
	FindDefaultExercise(catID int, level int) (*server.Exercise, error)
}
