package handler

import (
	"net/http"
)

type Router struct {}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Route() {
	http.HandleFunc("/", NewDefaultHandler().pong)
}