package user

import (
	"net/http"

	"github.com/rmitsubayashi/bodyweight-server/src/handler/util"
	usecase "github.com/rmitsubayashi/bodyweight-server/src/usecase/user"
)

type UserHandler struct {
	UseCase usecase.UserUseCase
}

func NewUserHandler() (*UserHandler, error) {
	uc, err := usecase.NewUserUseCase()
	if err != nil {
		return nil, err
	}
	return &UserHandler{
		UseCase: uc,
	}, nil
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	userID := 1
	user, err := h.UseCase.GetUserInfo(userID)
	if err != nil {
		//TODO split errors
		util.SendError(w, err, http.StatusInternalServerError)
		return
	}

	util.SendData(w, user)
}

func (h *UserHandler) PostUser(w http.ResponseWriter, r *http.Request) {
	err := h.UseCase.RegisterNewUser()
	if err != nil {
		//TODO split errors
		util.SendError(w, err, http.StatusInternalServerError)
		return
	}

	util.SendData(w, nil)
}
