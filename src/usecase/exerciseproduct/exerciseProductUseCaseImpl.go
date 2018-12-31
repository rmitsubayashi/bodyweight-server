package exerciseproduct

import (
	"github.com/rmitsubayashi/bodyweight-server/src/model/client"
)

type ExerciseProductUseCaseImpl struct{}

func (*ExerciseProductUseCaseImpl) GetAvailableExerciseProducts(userID int) ([]*client.ExerciseProduct, error) {
	return []*client.ExerciseProduct{
		client.NewExerciseProduct(2, client.ExerciseProductAttributes{
			Title: "chest focus set",
			Exercises: []client.Exercise{
				*client.NewExercise(41, client.ExerciseAttributes{
					Title:           "Wide arm pushups",
					Level:           5,
					Description:     "put your arms out wide. Push down.",
					MeasurementType: "REP",
					CategoryID:      0,
					TargetSets: []client.SetAttributes{
						client.SetAttributes{
							SetNumber: 1,
							Value:     10,
						},
						client.SetAttributes{
							SetNumber: 2,
							Value:     10,
						},
					},
					Quantity: 5,
				}),
				*client.NewExercise(24, client.ExerciseAttributes{
					Title:           "Spider man pushups",
					Level:           10,
					Description:     "put your arm in front of another. Push down and move forward",
					MeasurementType: "REP",
					CategoryID:      0,
					TargetSets: []client.SetAttributes{
						client.SetAttributes{
							SetNumber: 1,
							Value:     30,
						},
						client.SetAttributes{
							SetNumber: 2,
							Value:     30,
						},
					},
					Quantity: 5,
				}),
			},
			Price: 300,
		}),
	}, nil
}

func (*ExerciseProductUseCaseImpl) BuyExerciseProduct(userID int, ep client.ExerciseProduct) (*client.Points, error) {
	return client.NewPoints(4, client.PointsAttributes{
		Points: 200,
	}), nil
}

func NewExerciseProductUseCase() *ExerciseProductUseCaseImpl {
	return &ExerciseProductUseCaseImpl{}
}
