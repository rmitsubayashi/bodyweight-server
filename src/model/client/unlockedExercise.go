package client

type UnlockedExercise struct {
	Exercise       ExerciseTitle   `json:"exercise"`
	LevelUnlocked  int             `json:"level_unlocked"`
	OtherExercises []ExerciseTitle `json:"other_exercises"`
}
