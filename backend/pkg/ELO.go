package pkg

import "math"

const places float32 = 100

func roundTo2DecimalPlaces(num float32) float32 {
	rounded := math.Round(float64(num*places)) / float64(places)
	return float32(rounded)
}

func GetExpectedScore(winnerELO, loserELO float32) float32 {
	expectedScoreForWinner := 1 / (1 + math.Pow(10, (float64(loserELO)-float64(winnerELO))/400))
	return roundTo2DecimalPlaces(float32(expectedScoreForWinner))
}

func GetRatingChange(realScore, expectedScore float32, k_factor int) float32 {
	return float32(k_factor) * (realScore - expectedScore)
}
