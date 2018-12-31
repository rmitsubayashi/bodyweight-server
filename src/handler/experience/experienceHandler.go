package experience

import (
	"net/http"

	"github.com/rmitsubayashi/bodyweight-server/src/handler/util"
	usecase "github.com/rmitsubayashi/bodyweight-server/src/usecase/experience"
)

type ExperienceHandler struct {
	UseCase usecase.ExperienceUseCase
}

func NewExperienceHandler() *ExperienceHandler {
	return &ExperienceHandler{
		UseCase: usecase.NewExperienceUseCase(),
	}
}

func (h *ExperienceHandler) GetExperiences(w http.ResponseWriter, r *http.Request) {
	userID := 1
	experiences, err := h.UseCase.GetExperienceGraph(userID)
	if err != nil {
		//TODO split errors
		util.SendError(w, err, http.StatusInternalServerError)
		return
	}

	util.SendData(w, experiences)
}
