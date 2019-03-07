package exercise

import (
	"errors"

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

func (er *ExerciseRepoImpl) FindRandomExercise(catID int, minLev int, maxLev int, seed int, count int) (*[]server.Exercise, error) {
	selectQuery := `
	select * FROM exercise WHERE category_id=? AND level BETWEEN ? AND ?
	ORDER BY RAND(?) LIMIT ?
	`
	var e []server.Exercise
	if err := er.conn.Select(&e, selectQuery, catID, minLev, maxLev, seed, count); err != nil {
		return nil, errors.New("error getting random exercise" + err.Error())
	}
	return &e, nil
}

func (er *ExerciseRepoImpl) AddUserExercise(e *server.UserExercise) error {
	checkExistingExerciseStatement := `
	SELECT * FROM user_exercise WHERE user_id=? AND exercise_id=?
	`
	row, err := er.conn.Queryx(checkExistingExerciseStatement, e.UserID, e.ExerciseID)
	if err != nil {
		return errors.New("errrrr" + err.Error())
	}
	if !row.Next() {
		insertExerciseStatement := `
		INSERT INTO user_exercise (user_id, exercise_id, amount) VALUES (?, ?, ?)
		`
		_, err := er.conn.Exec(insertExerciseStatement, e.UserID, e.ExerciseID, e.Amount)
		return err
	} else {
		var qe server.UserExercise
		row.StructScan(&qe)
		// -1 = default exercise
		if qe.Amount == -1 {
			return nil
		}
		updateExerciseStatement := `
		UPDATE user_exercise SET amount = amount + ? WHERE user_id=? AND exercise_id=?
		`
		_, err := er.conn.Exec(updateExerciseStatement, e.Amount, e.UserID, e.ExerciseID)
		return err
	}
}

func (er *ExerciseRepoImpl) RemoveUserExercise(uid int, exerciseID int, amount int) error {
	currAmountStatement := `
	SELECT amount FROM user_exercise WHERE user_id=? AND exercise_id=?
	`
	var currAmount int
	if err := er.conn.Get(&currAmount, currAmountStatement, uid, exerciseID); err != nil {
		return err
	}
	// -1 means it's a default exercise
	if currAmount == -1 {
		return nil
	}
	newAmount := currAmount - amount
	if newAmount < 0 {
		return errors.New("don't have enough exercises")
	}
	if newAmount == 0 {
		removeExerciseStatement := `
		DELETE FROM user_exercise WHERE user_id=? AND exercise_id=?
		`
		_, err := er.conn.Exec(removeExerciseStatement, uid, exerciseID)
		return err
	} else {
		updateExerciseStatement := `
		UPDATE user_exercise SET amount = ? WHERE user_id=? AND exercise_id=?
		`
		_, err := er.conn.Exec(updateExerciseStatement, newAmount, uid, exerciseID)
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
