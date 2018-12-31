package client

type Experience struct {
	ID int `json:"id"`
	Type string `json:"type"`
	Attributes ExperienceAttributes `json:"attributes"`
}

type ExperienceAttributes struct {
	CategoryID int `json:"category_id"`
	Level int `json:"level"`
	NextLevelCurrent int `json:"next_level_current"`
	NextLevelTotal int `json:"next_level_total"`
}

func NewExperience(id int, attrs ExperienceAttributes) *Experience {
	return &Experience{
		ID: id,
		Type: "experience",
		Attributes: attrs,
	}
}