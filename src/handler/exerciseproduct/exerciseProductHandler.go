package exerciseproduct

import (
	"net/http"

	"github.com/rmitsubayashi/bodyweight-server/src/handler/util"
	"github.com/rmitsubayashi/bodyweight-server/src/model/client"
	usecase "github.com/rmitsubayashi/bodyweight-server/src/usecase/exerciseProduct"
)

type ExerciseProductHandler struct {
	UseCase usecase.ExerciseProductUseCase
}

func NewExerciseProductHandler() *ExerciseProductHandler {
	return &ExerciseProductHandler{
		UseCase: usecase.NewExerciseProductUseCase(),
	}
}

func (h *ExerciseProductHandler) GetExerciseProducts(w http.ResponseWriter, r *http.Request) {
	userID := 1
	exerciseProducts, err := h.UseCase.GetTodayExerciseProducts(userID)
	if err != nil {
		//TODO split errors
		util.SendError(w, err, http.StatusInternalServerError)
		return
	}

	util.SendData(w, exerciseProducts)
}

func (h *ExerciseProductHandler) PostExerciseProduct(w http.ResponseWriter, r *http.Request) {
	var ep client.ExerciseProduct
	err := util.GetData(r, ep)
	if err != nil {
		util.SendError(w, err, http.StatusBadRequest)
		return
	}
	userID := 1
	err = h.UseCase.BuyExerciseProduct(userID, ep)
	if err != nil {
		util.SendError(w, err, http.StatusInternalServerError)
		return
	}

	util.SendData(w, nil)
}
