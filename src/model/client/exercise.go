package client

type Exercise struct {
	ID              int    `json:"id"`
	Title           string `json:"title,omitempty"`
	Level           int    `json:"level,omitempty"`
	ImageURL        string `json:"image_url,omitempty"`
	Description     string `json:"description,omitempty"`
	MeasurementType string `json:"measurement_type,omitempty"`
	CategoryID      int    `json:"category_id,omitempty"`
	TargetSets      []Set  `json:"target_sets,omitempty"`
	Quantity        int    `json:"quantity,omitempty"`
}
