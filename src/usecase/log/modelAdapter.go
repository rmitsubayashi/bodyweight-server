package log

import (
	"github.com/rmitsubayashi/bodyweight-server/src/model/client"
	"github.com/rmitsubayashi/bodyweight-server/src/model/server"
	"github.com/rmitsubayashi/bodyweight-server/src/usecase/util"
)

func clientToServerLog(cl client.Log, uid int) server.Log {
	var logSets []server.LogSet
	for _, set := range cl.Sets {
		logSet := server.LogSet{
			LogID:      cl.ID,
			ExerciseID: set.Exercise.ID,
			SetNumber:  set.SetNumber,
			Value:      set.Value,
		}
		logSets = append(logSets, logSet)
	}

	return server.Log{
		ID:         cl.ID,
		UserID:     uid,
		CategoryID: cl.CategoryID,
		Sets:       logSets,
	}
}

func serverToClientLog(sl server.Log) client.Log {
	var sets []client.Set
	for _, logSet := range sl.Sets {
		set := client.Set{
			Exercise: client.Exercise{
				ID: logSet.ExerciseID,
			},
			SetNumber: logSet.SetNumber,
			Value:     logSet.Value,
		}
		sets = append(sets, set)
	}

	return client.Log{
		ID:         sl.ID,
		CategoryID: sl.CategoryID,
		Date:       util.FormatDateTime(sl.Date),
		Sets:       sets,
	}
}

func serverToClientLogWithExerciseDetails(sl server.Log, cs []client.Set) client.Log {
	return client.Log{
		ID:         sl.ID,
		CategoryID: sl.CategoryID,
		Date:       util.FormatDateTime(sl.Date),
		Sets:       cs,
	}
}

func serverToClientSet(ls server.LogSet, e server.Exercise) client.Set {
	var clientTargetSets []client.Set
	for _, targetSet := range e.TargetSets {
		set := client.Set{
			SetNumber: targetSet.SetNumber,
			Value:     targetSet.Value,
		}
		clientTargetSets = append(clientTargetSets, set)
	}
	clientExercise := client.Exercise{
		ID:              e.ID,
		Title:           e.Title,
		Level:           e.Level,
		ImageURL:        e.ImageURL,
		Description:     e.Description,
		MeasurementType: e.MeasurementType,
		TargetSets:      clientTargetSets,
	}
	return client.Set{
		Exercise:  clientExercise,
		SetNumber: ls.SetNumber,
		Value:     ls.Value,
	}
}

func clientUnlockedExerciseToServerUserExercise(ue client.UnlockedExercise, uid int, amount int) server.UserExercise {
	return server.UserExercise{
		ExerciseID: ue.Exercise.ID,
		UserID:     uid,
		Amount:     amount,
	}
}

func clientExerciseToServerUserExercise(e client.Exercise, uid int) server.UserExercise {
	return server.UserExercise{
		ExerciseID: e.ID,
		UserID:     uid,
		Amount:     e.Quantity,
	}
}
