package client

type Experience struct {
	ID int `json:"id"`
	CategoryID int `json:"category_id"`
	Level int `json:"level"`
	NextLevelCurrent int `json:"next_level_current"`
	NextLevelTotal int `json:"next_level_total"`
}
