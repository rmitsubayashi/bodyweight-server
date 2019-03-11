package client

type UnlockedExercise struct {
	Exercise       Exercise   `json:"exercise"`
	OtherExercises []Exercise `json:"other_exercises"`
}
