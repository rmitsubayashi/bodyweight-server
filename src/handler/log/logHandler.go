package log

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rmitsubayashi/bodyweight-server/src/handler/util"
	"github.com/rmitsubayashi/bodyweight-server/src/model/client"
	usecase "github.com/rmitsubayashi/bodyweight-server/src/usecase/log"
)

type LogHandler struct {
	UseCase usecase.LogUseCase
}

func NewLogHandler() *LogHandler {
	return &LogHandler{
		UseCase: usecase.NewLogUseCase(),
	}
}

func (h *LogHandler) GetLogs(w http.ResponseWriter, r *http.Request) {
	//TODO get user id from firebase id
	userID := 1
	logs, err := h.UseCase.GetLogList(userID)
	if err != nil {
		//TODO split errors
		util.SendError(w, err, http.StatusInternalServerError)
		return
	}

	util.SendData(w, logs)
}

func (h *LogHandler) GetLog(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	logID := vars["log_id"]
	logIDInt, err := strconv.Atoi(logID)
	if err != nil {
		util.SendError(w, err, http.StatusBadRequest)
	}
	log, err := h.UseCase.GetLogInfo(logIDInt)
	if err != nil {
		//TODO split errors
		util.SendError(w, err, http.StatusInternalServerError)
		return
	}

	util.SendData(w, log)
}

func (h *LogHandler) PostLog(w http.ResponseWriter, r *http.Request) {
	var l client.Log
	err := util.GetData(r, l)
	if err != nil {
		//TOTO split errors
		util.SendError(w, err, http.StatusInternalServerError)
		return
	}
	feedback, err := h.UseCase.RecordLog(l)
	if err != nil {
		util.SendError(w, err, http.StatusInternalServerError)
		return
	}

	util.SendData(w, feedback)
}
