package experience

import (
	"github.com/rmitsubayashi/bodyweight-server/src/model/client"
)

type ExperienceUseCaseImpl struct{}

func (*ExperienceUseCaseImpl) GetExperienceGraph(userID int) ([]*client.Experience, error) {
	return []*client.Experience{
		client.NewExperience(0, client.ExperienceAttributes{
			CategoryID:       0,
			Level:            30,
			NextLevelCurrent: 233,
			NextLevelTotal:   300,
		}),
		client.NewExperience(1, client.ExperienceAttributes{
			CategoryID:       1,
			Level:            25,
			NextLevelCurrent: 120,
			NextLevelTotal:   200,
		}),
		client.NewExperience(2, client.ExperienceAttributes{
			CategoryID:       2,
			Level:            11,
			NextLevelCurrent: 34,
			NextLevelTotal:   50,
		}),
		client.NewExperience(3, client.ExperienceAttributes{
			CategoryID:       3,
			Level:            21,
			NextLevelCurrent: 140,
			NextLevelTotal:   150,
		}),
		client.NewExperience(4, client.ExperienceAttributes{
			CategoryID:       4,
			Level:            10,
			NextLevelCurrent: 1,
			NextLevelTotal:   45,
		}),
		client.NewExperience(5, client.ExperienceAttributes{
			CategoryID:       5,
			Level:            34,
			NextLevelCurrent: 0,
			NextLevelTotal:   375,
		}),
	}, nil
}

func NewExperienceUseCase() *ExperienceUseCaseImpl {
	return &ExperienceUseCaseImpl{}
}
