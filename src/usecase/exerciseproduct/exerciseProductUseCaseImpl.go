package exerciseproduct

import (
	"errors"
	"math/rand"
	"time"

	"github.com/rmitsubayashi/bodyweight-server/src/model/client"
	er "github.com/rmitsubayashi/bodyweight-server/src/repository/exercise"
	tr "github.com/rmitsubayashi/bodyweight-server/src/repository/transaction"
	ur "github.com/rmitsubayashi/bodyweight-server/src/repository/user"
	"github.com/rmitsubayashi/bodyweight-server/src/usecase/util"
)

type ExerciseProductUseCaseImpl struct {
	exerciseRepo    er.ExerciseRepo
	transactionRepo tr.TransactionRepo
	userRepo        ur.UserRepo
}

func (uc *ExerciseProductUseCaseImpl) GetTodayExerciseProducts(userID int) (*[]client.ExerciseProduct, error) {
	u, err := uc.userRepo.GetUser(userID)
	if err != nil {
		return nil, err
	}
	catLevels := u.GetCatLevels()
	var eps []client.ExerciseProduct
	for index, level := range catLevels {
		minLev, maxLev := calculateMinMaxLevel(level)
		catLevel := index + 1
		seed := util.FormatDateSeed()
		userSeed := seed + userID
		rowCt := 1
		exercisesPtr, err := uc.exerciseRepo.FindRandomExercise(catLevel, minLev, maxLev, userSeed, rowCt)
		if err != nil {
			return nil, err
		}
		exercises := *exercisesPtr
		if len(exercises) == 0 {
			continue
		}
		rand.Seed(int64(userSeed))
		amount := rand.Intn(2) + 1
		price := exercises[0].Level * amount * 100
		for i := 1; i < amount; i++ {
			exercises = append(exercises, exercises[0])
		}

		title := exercises[0].Title

		// using ep here just to make st (one less function to write :P)
		ep := serverExercisesToClientExerciseProduct(exercises, title, price, false)
		st := clientExerciseProductToServerTransaction(ep, userID)
		sold, err := uc.transactionRepo.TransactionExists(st, time.Now())
		if err != nil {
			return nil, errors.New(err.Error())
		}
		ep = serverExercisesToClientExerciseProduct(exercises, title, price, sold)

		eps = append(eps, ep)
	}

	return &eps, nil
}

func calculateMinMaxLevel(lvl int) (int, int) {
	var min int
	if lvl == 1 {
		min = 1
	} else {
		min = lvl - 1
	}
	max := lvl
	return min, max
}

func (uc *ExerciseProductUseCaseImpl) BuyExerciseProduct(userID int, ep client.ExerciseProduct) error {
	// check that the user hasn't altered the product
	availableProducts, err := uc.GetTodayExerciseProducts(userID)
	if err != nil {
		return err
	}

	matched := false
	for _, p := range *availableProducts {
		if p.IsSameProduct(ep) {
			matched = true
		}
	}

	if !matched {
		return errors.New("invalid product")
	}

	// check that the user has enough points
	u, err := uc.userRepo.GetUser(userID)
	if err != nil {
		return err
	}
	if u.Points < ep.Price {
		return errors.New("not enough points")
	}

	if err := uc.userRepo.ChangePointsBy(userID, (-1)*ep.Price); err != nil {
		return err
	}

	ues := clientExerciseProductToServerUserExercises(ep, userID)
	for _, ue := range ues {
		if err := uc.exerciseRepo.AddUserExercise(&ue); err != nil {
			return err
		}
	}

	t := clientExerciseProductToServerTransaction(ep, userID)
	err = uc.transactionRepo.AddTransaction(t)

	return err
}

func NewExerciseProductUseCase() (*ExerciseProductUseCaseImpl, error) {
	exerRepo, err := er.NewExerciseRepo()
	if err != nil {
		return nil, err
	}
	tRepo, err := tr.NewTransactionRepo()
	if err != nil {
		return nil, err
	}
	uRepo, err := ur.NewUserRepo()
	if err != nil {
		return nil, err
	}
	return &ExerciseProductUseCaseImpl{
		exerciseRepo:    exerRepo,
		transactionRepo: tRepo,
		userRepo:        uRepo,
	}, nil
}
