package experience

import (
	"github.com/rmitsubayashi/bodyweight-server/src/model/client"
)

type ExperienceUseCaseImpl struct{}

func (*ExperienceUseCaseImpl) GetExperienceGraph(userID int) ([]*client.Experience, error) {
	return []*client.Experience {
		&client.Experience {
			ID: 0,
			CategoryID:       0,
			Level:            30,
			NextLevelCurrent: 233,
			NextLevelTotal:   300,
		},
		&client.Experience {
			ID: 1,
			CategoryID:       1,
			Level:            25,
			NextLevelCurrent: 120,
			NextLevelTotal:   200,
		},
		&client.Experience {
			ID: 2,
			CategoryID:       2,
			Level:            11,
			NextLevelCurrent: 34,
			NextLevelTotal:   50,
		},
		&client.Experience {
			ID: 3,
			CategoryID:       3,
			Level:            21,
			NextLevelCurrent: 140,
			NextLevelTotal:   150,
		},
		&client.Experience {
			ID: 4,
			CategoryID:       4,
			Level:            10,
			NextLevelCurrent: 1,
			NextLevelTotal:   45,
		},
		&client.Experience {
			ID: 5,
			CategoryID:       5,
			Level:            34,
			NextLevelCurrent: 0,
			NextLevelTotal:   375,
		},
	}, nil
}

func NewExperienceUseCase() *ExperienceUseCaseImpl {
	return &ExperienceUseCaseImpl{}
}
