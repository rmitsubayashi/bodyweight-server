package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rmitsubayashi/bodyweight-server/src/handler/exercise"
	"github.com/rmitsubayashi/bodyweight-server/src/handler/log"
)

type RouterImpl struct{}

func NewRouter() *RouterImpl {
	return &RouterImpl{}
}
func (ri *RouterImpl) Route() {
	r := mux.NewRouter()
	r.HandleFunc("/", NewDefaultHandler().pong).Methods(http.MethodGet)
	r.HandleFunc("/users/exercises", exercise.NewExerciseHandler().GetExercises).Methods(http.MethodGet)
	r.HandleFunc("/users/logs", log.NewLogHandler().GetLogs).Methods(http.MethodGet)
	r.HandleFunc("/users/logs/{log_id}", log.NewLogHandler().GetLog).Methods(http.MethodGet)
	r.HandleFunc("/users/logs", log.NewLogHandler().PostLog).Methods(http.MethodPost)
	http.Handle("/", r)
}
