package usecase

import (
	"github.com/rmitsubayashi/bodyweight-server/src/model/client"
)

type ExerciseUseCaseImpl struct {}

func (uc *ExerciseUseCaseImpl) GetExerciseList(categoryID int) ([]*client.Exercise, error) {
	return []*client.Exercise{
		&client.Exercise {
			ID: 1,
			Title: "One arm pushup",
			Level: 10,
			Description: "put your arms where blow your chest. Push down.",
			MeasurementType: "REP",
			CategoryID: categoryID,
			TargetSets: []client.Set {
				client.Set {
					SetNumber: 1,
					Value: 10,
				},
				client.Set {
					SetNumber: 2,
					Value: 10,
				},
			},
			Quantity: -1,
		},

		&client.Exercise {
			ID: 24,
			Title: "Hindu pushups",
			Level: 6,
			Description: "put your arms in a normla position. Lift your butt up.",
			MeasurementType: "REP",
			CategoryID: categoryID,
			TargetSets: []client.Set {
				client.Set {
					SetNumber: 1,
					Value: 20,
				},
				client.Set {
					SetNumber: 2,
					Value: 20,
				},
			},
			Quantity: 3,
		},
	}, nil
}

func NewExerciseUseCase() *ExerciseUseCaseImpl {
	return &ExerciseUseCaseImpl{}
}
