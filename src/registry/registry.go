package registry

import (
	"github.com/rmitsubayashi/bodyweight-server/src/handler"
)

func NewRouter() Router {
	return handler.NewRouter()
}
