package log

import (
	"github.com/rmitsubayashi/bodyweight-server/src/model/client"
	er "github.com/rmitsubayashi/bodyweight-server/src/repository/exercise"
	lr "github.com/rmitsubayashi/bodyweight-server/src/repository/log"
)

type LogUseCaseImpl struct {
	logRepo      lr.LogRepo
	exerciseRepo er.ExerciseRepo
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

func (uc *LogUseCaseImpl) RecordLog(log client.Log) (*client.Feedback, error) {
	logForServer := clientToServerLog(log)
	if err := uc.logRepo.AddLog(logForServer); err != nil {
		return nil, err
	}

	return &client.Feedback{
		ID:      2,
		Comment: "Great job doing your first ever one arm pushup! You're definitely getting stronger!",
		CommentHighlightSpans: [][2]int{
			[2]int{21, 31},
		},
		PreviousExperience: client.Experience{
			CategoryID:       0,
			Level:            19,
			NextLevelCurrent: 210,
			NextLevelTotal:   300,
		},
		AfterExperience: client.Experience{
			CategoryID:       0,
			Level:            20,
			NextLevelCurrent: 10,
			NextLevelTotal:   310,
		},
		PreviousPoints: 2300,
		AfterPoints:    2400,
		ExperienceDetails: []client.ExperienceDetail{
			client.ExperienceDetail{
				Description: "pushups x 6",
				Experience:  40,
			},
			client.ExperienceDetail{
				Description: " one arm pushup x 1",
				Experience:  60,
			},
		},
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
	}, nil
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
	return &LogUseCaseImpl{
		logRepo:      logRepo,
		exerciseRepo: exerRepo,
	}, nil
}
