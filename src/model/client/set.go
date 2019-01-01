package client

type Set struct {
	ExerciseID int `json:"exercise_id"`
	ExerciseTitle string `json:"exercise_title"`
	MeasurementType string `json:"measurement_type"`
	SetNumber int `json:"set_number"`
	Value int `json:"value"`
}