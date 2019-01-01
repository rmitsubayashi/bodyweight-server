package exerciseproduct

import (
	"github.com/rmitsubayashi/bodyweight-server/src/model/client"
)

type ExerciseProductUseCaseImpl struct{}

func (*ExerciseProductUseCaseImpl) GetAvailableExerciseProducts(userID int) ([]*client.ExerciseProduct, error) {
	return []*client.ExerciseProduct{
		&client.ExerciseProduct{
			ID: 2,
			Title: "chest focus set",
			Exercises: []client.Exercise{
				client.Exercise {
					ID: 41,
					Title:           "Wide arm pushups",
					Level:           5,
					Description:     "put your arms out wide. Push down.",
					MeasurementType: "REP",
					CategoryID:      0,
					TargetSets: []client.Set {
						client.Set {
							SetNumber: 1,
							Value:     10,
						},
						client.Set {
							SetNumber: 2,
							Value:     10,
						},
					},
					Quantity: 5,
				},
				client.Exercise{
					ID: 24,
					Title:           "Spider man pushups",
					Level:           10,
					Description:     "put your arm in front of another. Push down and move forward",
					MeasurementType: "REP",
					CategoryID:      0,
					TargetSets: []client.Set{
						client.Set{
							SetNumber: 1,
							Value:     30,
						},
						client.Set{
							SetNumber: 2,
							Value:     30,
						},
					},
					Quantity: 5,
				},
			},
			Price: 300,
		},
	}, nil
}

func (*ExerciseProductUseCaseImpl) BuyExerciseProduct(userID int, ep client.ExerciseProduct) (*client.Points, error) {
	return &client.Points {
		ID: 4,
		Value: 200,
	}, nil
}

func NewExerciseProductUseCase() *ExerciseProductUseCaseImpl {
	return &ExerciseProductUseCaseImpl{}
}
