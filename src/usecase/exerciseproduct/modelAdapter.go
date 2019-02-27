package exerciseproduct

import (
	"github.com/rmitsubayashi/bodyweight-server/src/model/client"
	"github.com/rmitsubayashi/bodyweight-server/src/model/server"
)

func clientExerciseProductToServerUserExercises(cep client.ExerciseProduct, uid int) []server.UserExercise {
	var ues []server.UserExercise
	for _, e := range cep.Exercises {
		ue := server.UserExercise{
			UserID:     uid,
			ExerciseID: e.ID,
			Amount:     e.Quantity,
		}
		ues = append(ues, ue)
	}

	return ues
}
