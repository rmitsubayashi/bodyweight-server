package log

import (
	"github.com/rmitsubayashi/bodyweight-server/src/model/client"
)

type LogUseCaseImpl struct{}

func (uc *LogUseCaseImpl) GetLogList(userID int) ([]*client.Log, error) {
	return []*client.Log{
		&client.Log{
			ID:         1,
			CategoryID: 0,
			Date:       "2018-12-29",
			Sets: []client.Set{
				client.Set{
					ExerciseID:      21,
					ExerciseTitle:   "One arm pushups",
					MeasurementType: "REP",
					SetNumber:       1,
					Value:           1,
				},
				client.Set{
					ExerciseID:      20,
					ExerciseTitle:   "Half one-arm pushups",
					MeasurementType: "REP",
					SetNumber:       2,
					Value:           13,
				},
				client.Set{
					ExerciseID:      20,
					ExerciseTitle:   "Half one-arm pushups",
					MeasurementType: "REP",
					SetNumber:       3,
					Value:           4,
				},
			},
			Memo: "Did my first fucking one arm pushup!!!",
		},

		&client.Log{
			ID:         1,
			CategoryID: 0,
			Date:       "2018-12-31",
			Sets: []client.Set{
				client.Set{
					ExerciseID:      21,
					ExerciseTitle:   "One arm pushups",
					MeasurementType: "REP",
					SetNumber:       1,
					Value:           2,
				},
				client.Set{
					ExerciseID:      20,
					ExerciseTitle:   "Half one-arm pushups",
					MeasurementType: "REP",
					SetNumber:       2,
					Value:           12,
				},
			},
			Memo: "Did two fucking one arm pushups!!!",
		},
	}, nil
}

func (uc *LogUseCaseImpl) GetLogInfo(logID int) (*client.Log, error) {
	return &client.Log{
		ID:         logID,
		CategoryID: 0,
		Date:       "2018-12-31",
		Sets: []client.Set{
			client.Set{
				ExerciseID:      21,
				ExerciseTitle:   "One arm pushups",
				MeasurementType: "REP",
				SetNumber:       1,
				Value:           2,
			},
			client.Set{
				ExerciseID:      20,
				ExerciseTitle:   "Half one-arm pushups",
				MeasurementType: "REP",
				SetNumber:       2,
				Value:           12,
			},
		},
		Memo: "Did two fucking one arm pushups!!!",
	}, nil
}

func (uc *LogUseCaseImpl) RecordLog(log client.Log) (*client.Feedback, error) {
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
		UnlockedExercises: []client.UnlockedExercise{
			client.UnlockedExercise{
				Exercise: client.ExerciseTitle{
					ExerciseID: 24,
					Title:      "close arm pushups",
				},
				LevelUnlocked: 20,
				OtherExercises: []client.ExerciseTitle{
					client.ExerciseTitle{
						ExerciseID: 29,
						Title:      "wide-arm pushups",
					},
					client.ExerciseTitle{
						ExerciseID: 30,
						Title:      "Hindu pushups",
					},
				},
			},
		},
		DroppedExercises: []client.ExerciseTitle{
			client.ExerciseTitle{
				ExerciseID: 21,
				Title:      "leg-raised pushups",
			},
			client.ExerciseTitle{
				ExerciseID: 22,
				Title:      "one-legged pushups",
			},
		},
	}, nil
}

func NewLogUseCase() *LogUseCaseImpl {
	return &LogUseCaseImpl{}
}
