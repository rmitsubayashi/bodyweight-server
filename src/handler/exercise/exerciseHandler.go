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

func NewExerciseHandler() *ExerciseHandler {
	return &ExerciseHandler{
		UseCase: usecase.NewExerciseUseCase(),
	}
}

func (h *ExerciseHandler) GetExercises(w http.ResponseWriter, r *http.Request) {
	categoryID, err := util.GetQueryParam(r, "category_id")
	if err != nil {
		util.SendError(w, err, http.StatusBadRequest)
		return
	}
	categoryIDInt, _ := strconv.Atoi(categoryID)
	exercises, err := h.UseCase.GetExercisesByCategory(categoryIDInt)
	if err != nil {
		util.SendError(w, err, http.StatusInternalServerError)
		return
	}
	util.SendData(w, exercises, "exercises")
}
