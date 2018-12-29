package handler

import (
	"net/http"
	"github.com/rmitsubayashi/bodyweight-server/src/handler/exercise"

)

type RouterImpl struct {}

func NewRouter() *RouterImpl {
	return &RouterImpl{}
}

func (r *RouterImpl) Route() {
	http.HandleFunc("/", NewDefaultHandler().pong)
	http.HandleFunc("/user/exercises", exercise.NewExerciseHandler().GetExercises)
}