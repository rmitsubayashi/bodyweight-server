package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rmitsubayashi/bodyweight-server/src/handler/exercise"
	"github.com/rmitsubayashi/bodyweight-server/src/handler/exerciseproduct"
	"github.com/rmitsubayashi/bodyweight-server/src/handler/log"
	"github.com/rmitsubayashi/bodyweight-server/src/handler/user"
)

type RouterImpl struct {
}

func NewRouter() *RouterImpl {
	return &RouterImpl{}
}
func (ri *RouterImpl) Route() {
	r := mux.NewRouter()
	exerciseH, err := exercise.NewExerciseHandler()
	if err != nil {
		fmt.Printf(err.Error())
	}
	logH, err := log.NewLogHandler()
	if err != nil {
		fmt.Printf(err.Error())
	}
	exerciseProductH, err := exerciseproduct.NewExerciseProductHandler()
	if err != nil {
		fmt.Printf(err.Error())
	}
	userH, err := user.NewUserHandler()
	if err != nil {
		fmt.Printf(err.Error())
	}

	r.HandleFunc("/", NewDefaultHandler().pong).Methods(http.MethodGet)
	r.HandleFunc("/users/exercises", exerciseH.GetExercises).Methods(http.MethodGet)
	r.HandleFunc("/users/logs", logH.GetLogs).Methods(http.MethodGet)
	r.HandleFunc("/users/logs/{log_id}", logH.GetLog).Methods(http.MethodGet)
	r.HandleFunc("/users/logs", logH.PostLog).Methods(http.MethodPost)
	r.HandleFunc("/shop/exercises", exerciseProductH.GetExerciseProducts).Methods(http.MethodGet)
	r.HandleFunc("/shop/exercises", exerciseProductH.PostExerciseProduct).Methods(http.MethodPost)
	r.HandleFunc("/users", userH.GetUser).Methods(http.MethodGet)
	r.HandleFunc("/users", userH.PostUser).Methods(http.MethodPost)
	http.Handle("/", r)
}
