package exercise

import (
	"net/http"
	"strconv"

	"github.com/rmitsubayashi/bodyweight-server/src/handler/util"
	usecase "github.com/rmitsubayashi/bodyweight-server/src/usecase/exercise"
)

type ExerciseHandler struct {
	UseCase usecase.ExerciseUseCase
}

func NewExerciseHandler() (*ExerciseHandler, error) {
	uc, err := usecase.NewExerciseUseCase()
	if err != nil {
		return nil, err
	}
	return &ExerciseHandler{
		UseCase: uc,
	}, nil
}

func (h *ExerciseHandler) GetExercises(w http.ResponseWriter, r *http.Request) {
	categoryID, err := util.GetQueryParam(r, "category_id")
	if err != nil {
		util.SendError(w, err, http.StatusBadRequest)
		return
	}
	categoryIDInt, err := strconv.Atoi(categoryID)
	if err != nil {
		util.SendError(w, err, http.StatusBadRequest)
	}

	exercises, err := h.UseCase.GetExerciseList(categoryIDInt)
	if err != nil {
		//TODO split errors
		util.SendError(w, err, http.StatusInternalServerError)
		return
	}

	util.SendData(w, exercises)
}
