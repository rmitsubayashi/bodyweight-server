package client

type Log struct {
	ID int `json:"id"`
	Type string `json:"type"`
	Attributes LogAttributes `json:"attributes"`
}

type LogAttributes struct {
	CateogryID int `json:"cateogry_id"`
	Date string `json:"date"`
	Sets []Set `json:"sets"`
	Memo string `json:"memo"`
}

func  NewLog(id int, attrs LogAttributes) *Log {
	return &Log {
		ID: id,
		Type: "log",
		Attributes: attrs,
	}
}