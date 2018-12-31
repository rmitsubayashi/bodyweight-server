package experience

import (
	"github.com/rmitsubayashi/bodyweight-server/src/model/client"
)

type ExperienceUseCaseImpl struct {}

func (*ExperienceUseCaseImpl) GetExperienceGraph(userID int) ([]*client.Experience, error) {
	return []*client.Experience {
		client.NewExperience(2, client.ExperienceAttributes{
			CategoryID: 1,
			Level: 30,
			NextLevelCurrent: 233,
			NextLevelTotal: 300,
		}),
	}, nil
}

func NewExperienceUseCase() *ExperienceUseCaseImpl {
	return &ExperienceUseCaseImpl{}
}