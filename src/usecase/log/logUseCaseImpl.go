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
					Exercise: client.Exercise{
						ID:              21,
						Title:           "One arm pushups",
						MeasurementType: "REP",
					},
					SetNumber: 1,
					Value:     1,
				},
				client.Set{
					Exercise: client.Exercise{
						ID:              20,
						Title:           "Half one-arm pushups",
						MeasurementType: "REP",
					},
					SetNumber: 2,
					Value:     13,
				},
				client.Set{
					Exercise: client.Exercise{
						ID:              20,
						Title:           "Half one-arm pushups",
						MeasurementType: "REP",
					},
					SetNumber: 3,
					Value:     4,
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
					Exercise: client.Exercise{
						ID:              21,
						Title:           "One arm pushups",
						MeasurementType: "REP",
					},
					SetNumber: 1,
					Value:     2,
				},
				client.Set{
					Exercise: client.Exercise{
						ID:              20,
						Title:           "Half one-arm pushups",
						MeasurementType: "REP",
					},
					SetNumber: 2,
					Value:     12,
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
				Exercise: client.Exercise{
					ID:              21,
					Title:           "One arm pushups",
					MeasurementType: "REP",
				},
				SetNumber: 1,
				Value:     2,
			},
			client.Set{
				Exercise: client.Exercise{
					ID:              20,
					Title:           "Half one-arm pushups",
					MeasurementType: "REP",
				},
				SetNumber: 2,
				Value:     12,
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

func NewLogUseCase() *LogUseCaseImpl {
	return &LogUseCaseImpl{}
}
