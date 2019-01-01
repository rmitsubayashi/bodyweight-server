package client

type Feedback struct {
	ID int `json:"id"`
	Comment string `json:"comment"`
	CommentHighlightSpans [][2]int `json:"comment_highlight_spans"`
	PreviousExperience Experience `json:"previous_experience"`
	AfterExperience Experience `json:"after_experience"`
	PreviousPoints int `json:"previous_points"`
	AfterPoints int `json:"after_points"`
	UnlockedExercises []Exercise `json:"unlocked_exercises"`
}