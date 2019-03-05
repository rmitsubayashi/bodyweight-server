package client

import "sort"

type ExerciseProduct struct {
	ID        int        `json:"id"`
	Title     string     `json:"title"`
	Exercises []Exercise `json:"exercises"`
	Price     int        `json:"price"`
	Sold      bool       `json:"sold"`
}

// only compare critical points for fraud detection
func (ep *ExerciseProduct) IsSameProduct(compEP ExerciseProduct) bool {
	if len(ep.Exercises) != len(compEP.Exercises) {
		return false
	}
	epExercises := make([]Exercise, len(ep.Exercises))
	compEPExercises := make([]Exercise, len(compEP.Exercises))
	copy(epExercises, ep.Exercises)
	copy(compEPExercises, compEP.Exercises)
	sort.Slice(epExercises, func(i, j int) bool {
		if epExercises[i].ID == epExercises[j].ID {
			return epExercises[i].Quantity > epExercises[j].Quantity
		}
		return epExercises[i].ID > epExercises[j].ID
	})
	sort.Slice(compEPExercises, func(i, j int) bool {
		if compEPExercises[i].ID == compEPExercises[j].ID {
			return compEPExercises[i].Quantity > compEPExercises[j].Quantity
		}
		return compEPExercises[i].ID > compEPExercises[j].ID
	})
	for i, e := range epExercises {
		compE := compEPExercises[i]
		if e.ID != compE.ID || e.Quantity != compE.Quantity {
			return false
		}
	}

	if ep.Price != compEP.Price {
		return false
	}

	return true
}
