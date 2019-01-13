package exerciseproduct

import (
	"github.com/rmitsubayashi/bodyweight-server/src/model/client"
)

type ExerciseProductUseCaseImpl struct{}

func (*ExerciseProductUseCaseImpl) GetTodayExerciseProducts(userID int) ([]*client.ExerciseProduct, error) {
	return []*client.ExerciseProduct{
		&client.ExerciseProduct{
			ID: 2,
			Title: "Hand stands",
			Exercises: []client.Exercise{
				client.Exercise {
					ID: 42,
					Title:           "Hand stand pushups",
					Level:           9,
					Description:     "Balance yourself on a wall. Push up",
					MeasurementType: "SECONDS",
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
			},
			Price: 300,
			Sold: false,
		},
		&client.ExerciseProduct{
			ID: 2,
			Title: "Wide arm pushups",
			Exercises: []client.Exercise{
				client.Exercise {
					ID: 41,
					Title:           "Wide arm pushups",
					Level:           5,
					Description:     "put your arms out wide. Push down.",
					MeasurementType: "REP",
					CategoryID:      1,
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
			},
			Price: 300,
			Sold: true,
		},
		&client.ExerciseProduct{
			ID: 22,
			Title: "wide arm pull ups",
			Exercises: []client.Exercise{
				client.Exercise {
					ID: 41,
					Title:           "Wide arm pull ups",
					Level:           6,
					Description:     "put your arms out wide. Pull up.",
					MeasurementType: "REP",
					CategoryID:      2,
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
			},
			Price: 300,
			Sold: true,
		},
		&client.ExerciseProduct{
			ID: 12,
			Title: "front lever",
			Exercises: []client.Exercise{
				client.Exercise {
					ID: 41,
					Title:           "front lever",
					Level:           10,
					Description:     "hold the bar. Lift up.",
					MeasurementType: "SECONDS",
					CategoryID:      3,
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
					Quantity: 3,
				},
			},
			Price: 200,
			Sold: false,
		},
		&client.ExerciseProduct{
			ID: 2,
			Title: "Assisted one legged squats",
			Exercises: []client.Exercise{
				client.Exercise {
					ID: 41,
					Title:           "Assisted one legged squats",
					Level:           7,
					Description:     "Hold on to a door. Push down",
					MeasurementType: "REP",
					CategoryID:      4,
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
					Quantity: 1,
				},
			},
			Price: 100,
			Sold: false,
		},
		&client.ExerciseProduct{
			ID: 2,
			Title: "Leg raises",
			Exercises: []client.Exercise{
				client.Exercise {
					ID: 41,
					Title:           "Leg raises",
					Level:           5,
					Description:     "Lie down. Lift up your legs.",
					MeasurementType: "REP",
					CategoryID:      5,
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
			},
			Price: 300,
			Sold: false,
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
