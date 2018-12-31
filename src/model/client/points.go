package client

type Points struct {
	ID int `json:"id"`
	Type string `json:"type"`
	Attributes PointsAttributes `json:"attributes"`
}

type PointsAttributes struct {
	Points int `json:"points"`
}

func NewPoints(id int, attrs PointsAttributes) *Points {
	return &Points {
		ID: id,
		Type: "points",
		Attributes: attrs,
	}
}