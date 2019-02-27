package client

type Log struct {
	ID         int    `json:"id"`
	CategoryID int    `json:"category_id"`
	Date       string `json:"date"`
	Sets       []Set  `json:"sets"`
	Memo       string `json:"memo"`
}
