package log

import (
	"math"

	"github.com/rmitsubayashi/bodyweight-server/src/model/client"
	mcomm "github.com/rmitsubayashi/bodyweight-server/src/model/common"
	er "github.com/rmitsubayashi/bodyweight-server/src/repository/exercise"
	lr "github.com/rmitsubayashi/bodyweight-server/src/repository/log"
	ur "github.com/rmitsubayashi/bodyweight-server/src/repository/user"
)

type LogUseCaseImpl struct {
	logRepo      lr.LogRepo
	exerciseRepo er.ExerciseRepo
	userRepo     ur.UserRepo
}

func (uc *LogUseCaseImpl) GetLogList(userID int) ([]*client.Log, error) {
	serverLogs, err := uc.logRepo.GetUserLogs(userID)
	if err != nil {
		return nil, err
	}

	var clientLogs []*client.Log
	for _, serverLog := range *serverLogs {
		clientLog := serverToClientLog(serverLog)
		clientLogs = append(clientLogs, &clientLog)
	}

	return clientLogs, nil

}

func (uc *LogUseCaseImpl) GetLogInfo(logID int) (*client.Log, error) {
	serverLog, err := uc.logRepo.GetUserLog(logID)
	if err != nil {
		return nil, err
	}
	//the server log only has an exercise ID so we need to fetch the exercise
	var clientSets []client.Set
	for _, set := range serverLog.Sets {
		exercise, err := uc.exerciseRepo.GetExercise(set.ExerciseID)
		if err != nil {
			return nil, err
		}

		clientSet := serverToClientSet(set, *exercise)
		clientSets = append(clientSets, clientSet)
	}
	clientLog := serverToClientLogWithExerciseDetails(*serverLog, clientSets)
	return &clientLog, nil
}

func (uc *LogUseCaseImpl) RecordLog(log client.Log, uid int) (*client.Feedback, error) {
	logForServer := clientToServerLog(log, uid)
	if err := uc.logRepo.AddLog(logForServer); err != nil {
		return nil, err
	}

	allExerciseIDs := make(map[int]bool)
	for _, s := range log.Sets {
		allExerciseIDs[s.Exercise.ID] = true
	}
	for id := range allExerciseIDs {
		if err := uc.exerciseRepo.RemoveUserExercise(uid, id, 1); err != nil {
			return nil, err
		}
	}

	fbPtr, err := uc.generateFeedback(log, uid)
	if err != nil {
		return nil, err
	}
	fb := *fbPtr
	p := fb.AfterPoints - fb.PreviousPoints
	if err := uc.userRepo.ChangePointsBy(uid, p); err != nil {
		return nil, err
	}
	for _, ue := range fb.UnlockedExercises {
		//unlocked exercises are all default (amount = -1)
		e := clientUnlockedExerciseToServerUserExercise(ue, uid, -1)
		uc.exerciseRepo.AddUserExercise(&e)
	}
	for _, de := range fb.DroppedExercises {
		e := clientExerciseToServerUserExercise(de, uid)
		uc.exerciseRepo.AddUserExercise(&e)
	}

	return fbPtr, nil
}

func (uc *LogUseCaseImpl) generateFeedback(log client.Log, uid int) (*client.Feedback, error) {
	u, err := uc.userRepo.GetUser(uid)
	if err != nil {
		return nil, err
	}

	p := uc.calculatePoints(log.Sets)
	// ue := uc.checkUnlockedExercises(log.Sets)

	return &client.Feedback{
		Comment: "Great job doing your first ever one arm pushup! You're definitely getting stronger!",
		CommentHighlightSpans: [][2]int{
			[2]int{21, 31},
		},
		PreviousPoints: u.Points,
		AfterPoints:    u.Points + p,
		UnlockedExercises: []client.UnlockedExercise{
			client.UnlockedExercise{
				Exercise: client.Exercise{
					ID:    22,
					Title: "close arm pushups",
				},
				LevelUnlocked: 20,
				OtherExercises: []client.Exercise{
					client.Exercise{
						ID:    29,
						Title: "wide-arm pushups",
					},
					client.Exercise{
						ID:    30,
						Title: "Hindu pushups",
					},
				},
			},
		},
		DroppedExercises: []client.Exercise{
			client.Exercise{
				ID:         21,
				CategoryID: 0,
				Title:      "leg-raised pushups",
				Level:      6,
				Quantity:   2,
			},
		},
	}, nil
}

func (uc *LogUseCaseImpl) calculatePoints(sets []client.Set) int {
	total := 0
	for _, set := range sets {
		var relativeAmount int
		switch mt := set.Exercise.MeasurementType; mt {
		case mcomm.MEASUREMENT_TYPE_REPS:
			relativeAmount = set.Value * 100
		case mcomm.MEASUREMENT_TYPE_SECONDS:
			relativeAmount = set.Value * 10
		}

		point := int(math.Pow(2, float64(set.Exercise.Level))) * relativeAmount
		total += point
	}
	return total
}

func NewLogUseCase() (*LogUseCaseImpl, error) {
	logRepo, err := lr.NewLogRepo()
	if err != nil {
		return nil, err
	}

	exerRepo, err := er.NewExerciseRepo()
	if err != nil {
		return nil, err
	}

	usrRepo, err := ur.NewUserRepo()
	if err != nil {
		return nil, err
	}

	return &LogUseCaseImpl{
		logRepo:      logRepo,
		exerciseRepo: exerRepo,
		userRepo:     usrRepo,
	}, nil
}
