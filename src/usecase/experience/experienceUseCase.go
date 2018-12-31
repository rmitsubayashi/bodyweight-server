package experience

import (
	"github.com/rmitsubayashi/bodyweight-server/src/model/client"
)

type ExperienceUseCase interface {
	GetExperienceGraph(userID int) ([]*client.Experience, error)
}