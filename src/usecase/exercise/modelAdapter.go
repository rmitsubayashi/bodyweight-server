package exercise

import (
	"github.com/rmitsubayashi/bodyweight-server/src/model/client"
	"github.com/rmitsubayashi/bodyweight-server/src/model/server"
)

func serverToClientUserExercises(sue []server.UserExercise, eMap map[int]server.Exercise) []client.Exercise {
	var ces []client.Exercise
	for _, ue := range sue {
		e := eMap[ue.ExerciseID]
		var clientTargetSets []client.Set
		for _, targetSet := range e.TargetSets {
			set := client.Set{
				SetNumber: targetSet.SetNumber,
				Value:     targetSet.Value,
			}
			clientTargetSets = append(clientTargetSets, set)
		}
		ce := client.Exercise{
			ID:              ue.ExerciseID,
			Title:           e.Title,
			Level:           e.Level,
			ImageURL:        e.ImageURL,
			Description:     e.Description,
			MeasurementType: e.MeasurementType,
			TargetSets:      clientTargetSets,
			Quantity:        ue.Amount,
		}
		ces = append(ces, ce)
	}

	return ces
}
