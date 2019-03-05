package log

import (
	"github.com/rmitsubayashi/bodyweight-server/src/model/client"
)

func calculatePoints(sets []client.Set) int {
	result := 0
	for _, set := range sets {
		targetSetValue := set.Exercise.TargetSets[set.SetNumber-1].Value
		performanceIndex := set.Value / targetSetValue
		levelAdjustedResult := set.Exercise.Level * performanceIndex
		result += levelAdjustedResult
	}

	return result
}