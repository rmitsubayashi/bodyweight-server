package log

import (
	"log"
	"math"

	"github.com/rmitsubayashi/bodyweight-server/src/model/client"
	mcomm "github.com/rmitsubayashi/bodyweight-server/src/model/common"
	er "github.com/rmitsubayashi/bodyweight-server/src/repository/exercise"
	lr "github.com/rmitsubayashi/bodyweight-server/src/repository/log"
	ur "github.com/rmitsubayashi/bodyweight-server/src/repository/user"
)

type LogUseCaseImpl struct {
	logRepo      lr.LogRepo
	exerciseRepo er.ExerciseRepo
	userRepo     ur.UserRepo
}

func (uc *LogUseCaseImpl) GetLogList(userID int) ([]*client.Log, error) {
	serverLogs, err := uc.logRepo.GetUserLogs(userID)
	if err != nil {
		return nil, err
	}

	var clientLogs []*client.Log
	for _, serverLog := range *serverLogs {
		clientLog := serverToClientLog(serverLog)
		clientLogs = append(clientLogs, &clientLog)
	}

	return clientLogs, nil

}

func (uc *LogUseCaseImpl) GetLogInfo(logID int) (*client.Log, error) {
	serverLog, err := uc.logRepo.GetUserLog(logID)
	if err != nil {
		return nil, err
	}
	//the server log only has an exercise ID so we need to fetch the exercise
	var clientSets []client.Set
	for _, set := range serverLog.Sets {
		exercise, err := uc.exerciseRepo.GetExercise(set.ExerciseID)
		if err != nil {
			return nil, err
		}

		clientSet := serverToClientSet(set, *exercise)
		clientSets = append(clientSets, clientSet)
	}
	clientLog := serverToClientLogWithExerciseDetails(*serverLog, clientSets)
	return &clientLog, nil
}

func (uc *LogUseCaseImpl) RecordLog(l client.Log, uid int) (*client.Feedback, error) {
	logForServer := clientToServerLog(l, uid)
	if err := uc.logRepo.AddLog(logForServer); err != nil {
		return nil, err
	}

	allExerciseIDs := make(map[int]bool)
	for _, s := range l.Sets {
		allExerciseIDs[s.Exercise.ID] = true
	}
	for id := range allExerciseIDs {
		if err := uc.exerciseRepo.RemoveUserExercise(uid, id, 1); err != nil {
			return nil, err
		}
	}

	fbPtr, err := uc.generateFeedback(l, uid)
	if err != nil {
		return nil, err
	}
	fb := *fbPtr
	p := fb.AfterPoints - fb.PreviousPoints
	if err := uc.userRepo.ChangePointsBy(uid, p); err != nil {
		return nil, err
	}
	for _, ue := range fb.UnlockedExercises {
		//unlocked exercises are all default exercises (amount = -1)
		e := clientUnlockedExerciseToServerUserExercise(ue, uid, -1)
		if err := uc.exerciseRepo.AddUserExercise(&e); err != nil {
			return nil, err
		}
		if err := uc.userRepo.SetUserLevel(uid, ue.Exercise.CategoryID, ue.Exercise.Level); err != nil {
			return nil, err
		}
	}
	/*
		for _, de := range fb.DroppedExercises {
			e := clientExerciseToServerUserExercise(de, uid)
			uc.exerciseRepo.AddUserExercise(&e)
		}*/

	return fbPtr, nil
}

func (uc *LogUseCaseImpl) generateFeedback(log client.Log, uid int) (*client.Feedback, error) {
	u, err := uc.userRepo.GetUser(uid)
	if err != nil {
		return nil, err
	}

	//we only have the exercise IDs so fetch the required info.
	// (we can have the client pass in the info, but we want the DB to be the SSOT)
	logPtr, err := uc.populateExerciseInfo(log)
	if err != nil {
		return nil, err
	}
	log = *logPtr
	p := uc.calculatePoints(log.Sets)
	ues, err := uc.checkUnlockedExercises(log.Sets, log.CategoryID, uid)
	if err != nil {
		return nil, err
	}

	return &client.Feedback{
		Comment: "Great job doing your first ever one arm pushup! You're definitely getting stronger!",
		CommentHighlightSpans: [][2]int{
			[2]int{21, 31},
		},
		PreviousPoints:    u.Points,
		AfterPoints:       u.Points + p,
		UnlockedExercises: *ues,
		DroppedExercises: []client.Exercise{
			client.Exercise{
				ID:         21,
				CategoryID: 0,
				Title:      "leg-raised pushups",
				Level:      6,
				Quantity:   2,
			},
		},
	}, nil
}

func (uc *LogUseCaseImpl) populateExerciseInfo(l client.Log) (*client.Log, error) {
	eIDSet := make(map[int]bool)
	for _, s := range l.Sets {
		eIDSet[s.Exercise.ID] = true
	}
	es := make(map[int]client.Exercise)
	for eID := range eIDSet {
		ePtr, err := uc.exerciseRepo.GetExercise(eID)
		if err != nil {
			return nil, err
		}
		e := serverToClientExercise(*ePtr)
		log.Printf("%+v\n\n", e)
		es[eID] = e
	}
	var sets []client.Set
	for _, s := range l.Sets {
		newS := client.Set{
			Exercise:  es[s.Exercise.ID],
			SetNumber: s.SetNumber,
			Value:     s.Value,
		}
		sets = append(sets, newS)
	}
	return &client.Log{
		ID:         l.ID,
		CategoryID: l.CategoryID,
		Date:       l.Date,
		Sets:       sets,
	}, nil
}

func (uc *LogUseCaseImpl) calculatePoints(sets []client.Set) int {
	total := 0
	for _, set := range sets {
		var relativeAmount int
		switch mt := set.Exercise.MeasurementType; mt {
		case mcomm.MEASUREMENT_TYPE_REPS:
			relativeAmount = set.Value * 100
		case mcomm.MEASUREMENT_TYPE_SECONDS:
			relativeAmount = set.Value * 10
		}

		point := int(math.Pow(2, float64(set.Exercise.Level))) * relativeAmount
		total += point
	}
	return total
}

func (uc *LogUseCaseImpl) checkUnlockedExercises(sets []client.Set, catID int, uid int) (*[]client.UnlockedExercise, error) {
	var ues []client.UnlockedExercise
	var es []client.Exercise
	for _, s := range sets {
		es = append(es, s.Exercise)
	}

	defaultExists := false
	for _, e := range es {
		se, err := uc.exerciseRepo.GetExercise(e.ID)
		if err != nil {
			return nil, err
		}
		if se.IsDefault {
			defaultExists = true
		}
	}
	if !defaultExists {
		return &ues, nil
	}

	u, err := uc.userRepo.GetUser(uid)
	if err != nil {
		return nil, err
	}
	uCats := u.GetCatLevels()
	catLvl := uCats[catID-1]
	for _, e := range es {
		// if the user's level is lower than the exercise's level,
		// there is nothing to unlock.
		// also, the user can't do exercises higher than the user's level
		if e.Level != catLvl {
			continue
		}
		targetSetIndex := 0
		unlockable := true
		for _, s := range sets {
			if s.Exercise.ID != e.ID {
				continue
			}
			if e.TargetSets[targetSetIndex].Value > s.Value {
				unlockable = false
				break
			}
			targetSetIndex++
			if targetSetIndex >= len(e.TargetSets) {
				break
			}
		}
		if unlockable {
			e, err := uc.exerciseRepo.FindDefaultExercise(catID, catLvl+1)
			if err != nil {
				return nil, err
			}
			// we want the user to get a better picture of what he has unlocked
			otherEs, err := uc.exerciseRepo.FindRandomExercise(catID, catLvl+1, catLvl+1, 0, 2)
			if err != nil {
				return nil, err
			}
			ce := serverToClientExercise(*e)
			var coes []client.Exercise
			for _, otherE := range *otherEs {
				cOtherE := serverToClientExercise(otherE)
				coes = append(coes, cOtherE)
			}
			ue := client.UnlockedExercise{
				Exercise:       ce,
				OtherExercises: coes,
			}
			ues = append(ues, ue)
		}
	}

	return &ues, nil

}

func NewLogUseCase() (*LogUseCaseImpl, error) {
	logRepo, err := lr.NewLogRepo()
	if err != nil {
		return nil, err
	}

	exerRepo, err := er.NewExerciseRepo()
	if err != nil {
		return nil, err
	}

	usrRepo, err := ur.NewUserRepo()
	if err != nil {
		return nil, err
	}

	return &LogUseCaseImpl{
		logRepo:      logRepo,
		exerciseRepo: exerRepo,
		userRepo:     usrRepo,
	}, nil
}
