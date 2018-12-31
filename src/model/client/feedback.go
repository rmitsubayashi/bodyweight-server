package client

type Feedback struct {
	ID int `json:"id"`
	Type string `json:"type"`
	Attributes FeedbackAttributes `json:"attributes"`
}

type FeedbackAttributes struct {
	Comment string `json:"comment"`
	CommentHighlightSpans [][2]int `json:"comment_highlight_spans"`
	PreviousExperience ExperienceAttributes `json:"previous_experience"`
	AfterExperience ExperienceAttributes `json:"after_experience"`
	PreviousPoints int `json:"previous_points"`
	AfterPoints int `json:"after_points"`
	UnlockedExercises []ExerciseAttributes `json:"unlocked_exercises"`
}

func NewFeedback(id int, attrs FeedbackAttributes) *Feedback {
	return &Feedback {
		ID: id,
		Type: "feedback",
		Attributes: attrs,
	}
}