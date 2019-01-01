package client

type ExerciseProduct struct {
	ID         int                       `json:"id"`
	Title     string     `json:"title"`
	Exercises []Exercise `json:"exercises"`
	Price     int        `json:"price`
}