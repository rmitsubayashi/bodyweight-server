package handler

import (
	"net/http"
	"fmt"
)

type TestHandler struct {

}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

func (h *TestHandler) test(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Hello, World!")
}