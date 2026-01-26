package pkg_test

import (
	"testing"

	"github.com/runattekky/footyrate/models"
	"github.com/runattekky/footyrate/pkg"
)

func TestCompareLogic(t *testing.T) {
	t.Run("Test Comparing Logic", func(t *testing.T) {
		winner := models.Footballer{
			Name:     "RunAt",
			ELO:      1500,
			ImgURL:   "",
			K_Factor: 20,
			Rounds:   60,
		}

		loser := models.Footballer{
			Name:     "Minato",
			ELO:      1400,
			ImgURL:   "",
			K_Factor: 20,
			Rounds:   55,
		}

		pkg.UpdateELO(&winner, &loser)
		want := float32(1500 + 7.2)
		assertFloat(t, winner.GetELO(), want)
	})
}
