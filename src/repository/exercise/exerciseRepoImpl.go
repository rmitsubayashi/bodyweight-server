package exercise

import (
	"github.com/jmoiron/sqlx"

	"github.com/rmitsubayashi/bodyweight-server/src/model/server"
	"github.com/rmitsubayashi/bodyweight-server/src/repository"
)

type ExerciseRepoImpl struct {
	conn *sqlx.DB
}

func (er *ExerciseRepoImpl) GetExercise(exerciseID int) (*server.Exercise, error) {
	getExerciseStatement := `
	SELECT * FROM exercise
	WHERE id=?
	`
	var exercise server.Exercise
	if err := er.conn.Get(&exercise, getExerciseStatement, exerciseID); err != nil {
		return nil, err
	}

	getTargetSetsStatement := `
	SELECT * FROM target_set
	WHERE exercise_id=?
	`
	var ts []server.TargetSet
	if err := er.conn.Select(&ts, getTargetSetsStatement, exerciseID); err != nil {
		return nil, err
	}
	exercise.SetTargetSets(ts)

	return &exercise, nil
}

func (er *ExerciseRepoImpl) FindMaxSingleSetValue(exerciseID int) (int, error) {
	return -1, nil
}

func (er *ExerciseRepoImpl) FindMaxTotalSetValue(exerciseID int) (int, error) {
	return -1, nil
}

func (er *ExerciseRepoImpl) FindUserExercises(userID int, categoryID int) (*[]server.UserExercise, map[int]server.Exercise, error) {
	getExercisesStatement := `
	SELECT user_exercise.id, user_exercise.user_id, user_exercise.exercise_id, user_exercise.amount
	FROM user_exercise
	INNER JOIN exercise ON user_exercise.exercise_id=exercise.id
	WHERE user_exercise.user_id=? AND exercise.category_id=?
	`

	var ue []server.UserExercise
	if err := er.conn.Select(&ue, getExercisesStatement, userID, categoryID); err != nil {
		return nil, nil, err
	}

	m := make(map[int]server.Exercise)
	for _, e := range ue {

		// since the user exercise has only the exercise ID, get the exercise info as well
		getExerciseStatement := `
		SELECT * FROM exercise
		WHERE id=?
		`
		var eInfo server.Exercise
		if err := er.conn.Get(&eInfo, getExerciseStatement, e.ExerciseID); err != nil {
			return nil, nil, err
		}
		m[eInfo.ID] = eInfo
	}

	return &ue, m, nil
}

func (er *ExerciseRepoImpl) AddUserExercise(e *server.UserExercise) error {
	checkExistingExerciseStatement := `
	SELECT COUNT(*) FROM user_exercise WHERE user_id=? AND exercise_id=?
	`
	var rowCt *int
	if err := er.conn.Get(&rowCt, checkExistingExerciseStatement, e.UserID, e.ExerciseID); err != nil {
		return err
	}
	if *rowCt == 0 {
		insertExerciseStatement := `
		INSERT INTO user_exercise (user_id, exercise_id, amount) VALUES (?, ?, ?)
		`
		_, err := er.conn.Exec(insertExerciseStatement, e.UserID, e.ExerciseID, e.Amount)
		return err
	} else {
		updateExerciseStatement := `
		UPDATE user_exercise SET amount = amount + ? WHERE user_id=? AND exercise_id=?
		`
		_, err := er.conn.Exec(updateExerciseStatement, e.Amount, e.UserID, e.ExerciseID)
		return err
	}
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
