package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rmitsubayashi/bodyweight-server/src/handler/exercise"
	"github.com/rmitsubayashi/bodyweight-server/src/handler/exerciseproduct"
	"github.com/rmitsubayashi/bodyweight-server/src/handler/experience"
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
	r.HandleFunc("/shop/exercises", exerciseproduct.NewExerciseProductHandler().GetExerciseProducts).Methods(http.MethodGet)
	r.HandleFunc("/shop/exercises/purchase", exerciseproduct.NewExerciseProductHandler().PostExerciseProduct).Methods(http.MethodPost)
	r.HandleFunc("/users/experiences", experience.NewExperienceHandler().GetExperiences).Methods(http.MethodGet)
	http.Handle("/", r)
}
