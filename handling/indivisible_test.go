package handling

import (
	"testing"

	"github.com/theTardigrade/golang-validation/data"
)

type indivisibleIntDummyModel struct {
	x int `validation:"indivisible=-20,indivisible=5"`
}

func TestIndivisible_intInvalid(t *testing.T) {
	model := indivisibleIntDummyModel{}
	datum := indivisibleDatum{}

	for _, f := range []int{60, -20, 0, 20, 4e6} {
		model.x = f

		executeTest(t, model, func(m *data.Main, t *data.Tag) (s []string) {
			s = append(s, datum.FailureMessage(m, t))
			return
		})
	}
}

func TestIndivisible_intValid(t *testing.T) {
	model := indivisibleIntDummyModel{}

	for _, f := range []int{-9, 1, 9, 21, 44} {
		model.x = f

		executeTest(t, model, nil)
	}
}

type indivisibleUintDummyModel struct {
	x uint `validation:"indivisible=20,indivisible=5"`
}

func TestIndivisible_uintInvalid(t *testing.T) {
	model := indivisibleUintDummyModel{}
	datum := indivisibleDatum{}

	for _, f := range []uint{60, 0, 20, 4e6} {
		model.x = f

		executeTest(t, model, func(m *data.Main, t *data.Tag) (s []string) {
			s = append(s, datum.FailureMessage(m, t))
			return
		})
	}
}

func TestIndivisible_uintValid(t *testing.T) {
	model := indivisibleUintDummyModel{}

	for _, f := range []uint{1, 9, 21, 44} {
		model.x = f

		executeTest(t, model, nil)
	}
}
