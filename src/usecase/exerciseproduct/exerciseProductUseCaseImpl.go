package exerciseproduct

import (
	"github.com/rmitsubayashi/bodyweight-server/src/model/client"
)

type ExerciseProductUseCaseImpl struct {}

func (*ExerciseProductUseCaseImpl) GetAvailableExerciseProducts(userID int) ([]*client.ExerciseProduct, error) {
	return []*client.ExerciseProduct {
		client.NewExerciseProduct(2, client.ExerciseProductAttributes{
			Title: "push-up set",
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