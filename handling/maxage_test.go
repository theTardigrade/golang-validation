package handling

import (
	"testing"
	"time"

	"github.com/theTardigrade/golang-validation/data"
)

type maxageTimeDummyModel struct {
	X time.Time `validation:"maxage=20"`
}

func TestMaxage_timeInvalid(t *testing.T) {
	model := maxageTimeDummyModel{}
	datum := maxageDatum{}
	now := time.Now()

	for _, a := range []time.Time{
		now.AddDate(-21, 0, 0),
		now.AddDate(-50, 0, 0),
		now.AddDate(-1e9, 0, 0),
	} {
		model.X = a

		executeTest(t, model, func(m *data.Main, t *data.Tag) (s []string) {
			s = append(s, datum.FailureMessage(m, t))
			return
		})
	}
}

func TestMaxage_timeValid(t *testing.T) {
	model := maxageTimeDummyModel{}
	now := time.Now()

	for _, a := range []time.Time{
		now,
		now.AddDate(-19, 0, 0),
		now.AddDate(-20, 0, 0),
		now.AddDate(20, 0, 0),
		now.AddDate(1e6, 0, 0),
	} {
		model.X = a

		executeTest(t, model, nil)
	}
}
