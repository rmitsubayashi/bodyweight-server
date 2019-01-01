package client

type Log struct {
	ID         int           `json:"id"`
	CategoryID int             `json:"cateogry_id"`
	Date       string          `json:"date"`
	Sets       []Set `json:"sets"`
	Memo       string          `json:"memo"`
}
