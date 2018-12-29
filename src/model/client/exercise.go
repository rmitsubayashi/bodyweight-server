package client

type Exercise struct {
	ID int `json:"id"`
	Type string `json:"type"`
	Attributes ExerciseAttributes `json:"attributes"`
}

type ExerciseAttributes struct {
	Title string `json:"title"`
	Level int `json:"level"`
	ImageURL string `json:"image_url"`
	Description string `json:"description"`
	MeasurementType string `json:"measurement_type"`
	CategoryID int `json:"category_id"`
	TargetSets []Set `json:"target_sets"`
	Quantity int `json:"quantity"`
}

func NewExercise(id int, attrs ExerciseAttributes) *Exercise {
	return &Exercise {
		ID: id,
		Type: "exercise",
		Attributes: attrs,
	}
}