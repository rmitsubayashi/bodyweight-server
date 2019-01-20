package server

type Exercise struct {
	ID int
	Title string
	Description string
	Level int
	ImageURL string
	MeasureType int
	CategoryID int
	IsDefault bool
	TargetSets []Set
}