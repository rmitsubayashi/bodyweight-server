package client

type ExerciseProduct struct {
	ID int `json:"id"`
	Type string `json:"type"`
	Attributes ExerciseProductAttributes `json:"attributes"`
}

type ExerciseProductAttributes struct {
	Title string `json:"title"`
	Exercises []Exercise `json:"exercises"`
	Price int `json:"price`
}

func NewExerciseProduct(id int, attrs ExerciseProductAttributes) *ExerciseProduct {
	return &ExerciseProduct {
		ID: id,
		Type: "exerciseProduct",
		Attributes: attrs,
	}
}