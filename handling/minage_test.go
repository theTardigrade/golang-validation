package handling

import (
	"testing"
	"time"

	"github.com/theTardigrade/golang-validation/data"
)

type minageTimeDummyModel struct {
	X time.Time `validation:"minage=20"`
}

func TestMinage_timeInvalid(t *testing.T) {
	model := minageTimeDummyModel{}
	datum := minageDatum{}
	now := time.Now()

	for _, a := range []time.Time{
		now,
		now.AddDate(-19, 0, 0),
		now.AddDate(20, 0, 0),
		now.AddDate(1e6, 0, 0),
	} {
		model.X = a

		executeTest(t, model, func(m *data.Main, t *data.Tag) (s []string) {
			s = append(s, datum.FailureMessage(m, t))
			return
		})
	}
}

func TestMinage_timeValid(t *testing.T) {
	model := minageTimeDummyModel{}
	now := time.Now()

	for _, a := range []time.Time{
		now.AddDate(-20, 0, 0),
		now.AddDate(-21, 0, 0),
		now.AddDate(-1e9, 0, 0),
	} {
		model.X = a

		executeTest(t, model, nil)
	}
}
