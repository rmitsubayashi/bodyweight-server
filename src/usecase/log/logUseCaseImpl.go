package log

import (
	"github.com/rmitsubayashi/bodyweight-server/src/model/client"
)

type LogUseCaseImpl struct{}

func (uc *LogUseCaseImpl) GetLogList(userID int) ([]*client.Log, error) {
	return []*client.Log{
		client.NewLog(1, client.LogAttributes{
			CategoryID: 0,
			Date:       "2018-12-29",
			Sets: []client.SetAttributes{
				client.SetAttributes{
					ExerciseID:      21,
					ExerciseTitle:   "One arm pushups",
					MeasurementType: "REP",
					SetNumber:       1,
					Value:           1,
				},
				client.SetAttributes{
					ExerciseID:      20,
					ExerciseTitle:   "Half one-arm pushups",
					MeasurementType: "REP",
					SetNumber:       2,
					Value:           13,
				},
				client.SetAttributes{
					ExerciseID:      20,
					ExerciseTitle:   "Half one-arm pushups",
					MeasurementType: "REP",
					SetNumber:       3,
					Value:           4,
				},
			},
			Memo: "Did my first fucking one arm pushup!!!",
		}),

		client.NewLog(1, client.LogAttributes{
			CategoryID: 0,
			Date:       "2018-12-31",
			Sets: []client.SetAttributes{
				client.SetAttributes{
					ExerciseID:      21,
					ExerciseTitle:   "One arm pushups",
					MeasurementType: "REP",
					SetNumber:       1,
					Value:           2,
				},
				client.SetAttributes{
					ExerciseID:      20,
					ExerciseTitle:   "Half one-arm pushups",
					MeasurementType: "REP",
					SetNumber:       2,
					Value:           12,
				},
			},
			Memo: "Did two fucking one arm pushups!!!",
		}),
	}, nil
}

func (uc *LogUseCaseImpl) GetLogInfo(logID int) (*client.Log, error) {
	return client.NewLog(logID, client.LogAttributes{
		CategoryID: 0,
		Date:       "2018-12-31",
		Sets: []client.SetAttributes{
			client.SetAttributes{
				ExerciseID:      21,
				ExerciseTitle:   "One arm pushups",
				MeasurementType: "REP",
				SetNumber:       1,
				Value:           2,
			},
			client.SetAttributes{
				ExerciseID:      20,
				ExerciseTitle:   "Half one-arm pushups",
				MeasurementType: "REP",
				SetNumber:       2,
				Value:           12,
			},
		},
		Memo: "Did two fucking one arm pushups!!!",
	}), nil
}

func (uc *LogUseCaseImpl) RecordLog(log client.Log) (*client.Feedback, error) {
	return client.NewFeedback(2, client.FeedbackAttributes{
		Comment: "Great job doing your first ever one arm pushup! You're definitely getting stronger!",
		CommentHighlightSpans: [][2]int{
			[2]int{21, 31},
		},
		PreviousExperience: client.ExperienceAttributes {
			CategoryID: 0,
			Level: 20,
			NextLevelCurrent: 210,
			NextLevelTotal: 300,
		},
		AfterExperience: client.ExperienceAttributes {
			CategoryID: 0,
			Level: 20,
			NextLevelCurrent: 230,
			NextLevelTotal: 300,
		},
		PreviousPoints: 2300,
		AfterPoints: 2400,
	}), nil
}

func NewLogUseCase() *LogUseCaseImpl {
	return &LogUseCaseImpl{}
}
