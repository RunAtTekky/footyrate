package pkg

import (
	"math"
	"testing"
)

func TestELOchange(t *testing.T) {
	t.Run("1500 (Winner) vs 1400 (Loser) ELO should give 0.64 expected score for winner", func(t *testing.T) {
		var winnerELO float32 = 1500.0
		var loserELO float32 = 1400.0
		got := GetExpectedScore(winnerELO, loserELO)
		want := float32(0.64)

		assertFloat(t, got, want)
	})
}

func assertFloat(t *testing.T, got, want float32) {
	t.Helper()
	diff_limit := 0.01
	diff := float64(got - want)
	if math.Abs(diff) > diff_limit {
		t.Errorf("Wrong ELO change, got %.2f but want %.2f", got, want)
	}
}
