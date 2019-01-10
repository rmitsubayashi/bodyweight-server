package client

type UnlockedExercise struct {
	Exercise       Exercise   `json:"exercise"`
	LevelUnlocked  int             `json:"level_unlocked"`
	OtherExercises []Exercise `json:"other_exercises"`
}
