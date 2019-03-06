package log

import (
	"github.com/rmitsubayashi/bodyweight-server/src/model/client"
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

	feedback := uc.generateFeedback(log)
	p := feedback.AfterPoints - feedback.PreviousPoints
	if err := uc.userRepo.ChangePointsBy(uid, p); err != nil {
		return nil, err
	}

	return &feedback, nil
}

func (uc *LogUseCaseImpl) generateFeedback(log client.Log) client.Feedback {
	return client.Feedback{
		ID:      2,
		Comment: "Great job doing your first ever one arm pushup! You're definitely getting stronger!",
		CommentHighlightSpans: [][2]int{
			[2]int{21, 31},
		},
		PreviousPoints: 2300,
		AfterPoints:    2400,
		 UnlockedExercises: []client.UnlockedExercise{
			client.UnlockedExercise{
				Exercise: client.Exercise{
					ID:    24,
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
			},
			client.Exercise{
				ID:         22,
				CategoryID: 0,
				Title:      "one-legged pushups",
				Level:      5,
			},
		},
	}
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
