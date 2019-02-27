package client

type Feedback struct {
	ID                    int                `json:"id"`
	Comment               string             `json:"comment"`
	CommentHighlightSpans [][2]int           `json:"comment_highlight_spans"`
	PreviousPoints        int                `json:"previous_points"`
	AfterPoints           int                `json:"after_points"`
	UnlockedExercises     []UnlockedExercise `json:"unlocked_exercises"`
	DroppedExercises      []Exercise  `json:"dropped_exercises"`
}
