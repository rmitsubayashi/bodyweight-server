package usecase

import (
	"github.com/rmitsubayashi/bodyweight-server/src/model/client"
)

type ExerciseUseCaseImpl struct {}

func (uc *ExerciseUseCaseImpl) GetExerciseList(categoryID int) ([]*client.Exercise, error) {
	return []*client.Exercise{
		client.NewExercise(1, client.ExerciseAttributes{
			Title: "One arm pushup",
			Level: 10,
			Description: "put your arms where blow your chest. Push down.",
			MeasurementType: "REP",
			CategoryID: categoryID,
			TargetSets: []client.SetAttributes{
				client.SetAttributes {
					SetNumber: 1,
					Value: 10,
				},
				client.SetAttributes {
					SetNumber: 2,
					Value: 10,
				},
			},
			Quantity: -1,
		}),

		client.NewExercise(24, client.ExerciseAttributes{
			Title: "Hindu pushups",
			Level: 6,
			Description: "put your arms in a normla position. Lift your butt up.",
			MeasurementType: "REP",
			CategoryID: categoryID,
			TargetSets: []client.SetAttributes{
				client.SetAttributes {
					SetNumber: 1,
					Value: 20,
				},
				client.SetAttributes {
					SetNumber: 2,
					Value: 20,
				},
			},
			Quantity: 3,
		}),
	}, nil
}

func NewExerciseUseCase() *ExerciseUseCaseImpl {
	return &ExerciseUseCaseImpl{}
}
