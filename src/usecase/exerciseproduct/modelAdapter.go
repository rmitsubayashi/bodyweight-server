package exerciseproduct

import (
	"sort"

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

func clientExerciseProductToServerTransaction(cep client.ExerciseProduct, uid int) server.Transaction {
	var tes []server.TransactionExercise
	for _, e := range cep.Exercises {
		te := server.TransactionExercise{
			ExerciseID: e.ID,
			Amount:     e.Quantity,
		}
		tes = append(tes, te)
	}

	return server.Transaction{
		Price:     cep.Price,
		UserID:    uid,
		Exercises: tes,
	}
}

func serverExercisesToClientExerciseProduct(se []server.Exercise, title string, price int, sold bool) client.ExerciseProduct {
	eMap := make(map[int]*client.Exercise)
	for _, e := range se {
		ePtr := eMap[e.ID]
		if ePtr == nil {
			ce := client.Exercise{
				ID:          e.ID,
				Title:       e.Title,
				Level:       e.Level,
				ImageURL:    e.ImageURL,
				Description: e.Description,
				CategoryID:  e.CategoryID,
				Quantity:    1,
			}
			eMap[e.ID] = &ce
		} else {
			(*ePtr).SetQuantity((*ePtr).Quantity + 1)
		}
	}

	var ces []client.Exercise
	for _, e := range eMap {
		ces = append(ces, *e)
	}
	sort.Slice(ces, func(i, j int) bool {
		if ces[i].Quantity == ces[j].Quantity {
			return ces[i].Title < ces[j].Title
		}
		return ces[i].Quantity > ces[j].Quantity
	})

	return client.ExerciseProduct{
		Title:     title,
		Exercises: ces,
		Price:     price,
		Sold:      sold,
	}
}
