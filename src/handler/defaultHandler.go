package handler

import (
	"fmt"
	"net/http"
)

type DefaultHandler struct {
}

func NewDefaultHandler() *DefaultHandler {
	return &DefaultHandler{}
}

func (h *DefaultHandler) pong(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pong")
}
