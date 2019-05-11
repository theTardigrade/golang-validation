package handling

import (
	"testing"

	"github.com/theTardigrade/validation/internal/data"
)

type divisibleIntDummyModel struct {
	x int `validation:"divisible=-20,divisible=5"`
}

func TestDivisible_intInvalid(t *testing.T) {
	model := divisibleIntDummyModel{}
	datum := divisibleDatum{}

	for _, f := range []int{-9, 1, 9, 21, 44} {
		model.x = f

		executeTest(t, model, func(m *data.Main, t *data.Tag) (s []string) {
			s = append(s, datum.FailureMessage(m, t))
			return
		})
	}
}

func TestDivisible_intValid(t *testing.T) {
	model := divisibleIntDummyModel{}

	for _, f := range []int{60, -20, 0, 20, 4e6} {
		model.x = f

		executeTest(t, model, nil)
	}
}

type divisibleUintDummyModel struct {
	x uint `validation:"divisible=20,divisible=5"`
}

func TestDivisible_uintInvalid(t *testing.T) {
	model := divisibleUintDummyModel{}
	datum := divisibleDatum{}

	for _, f := range []uint{1, 9, 21, 44} {
		model.x = f

		executeTest(t, model, func(m *data.Main, t *data.Tag) (s []string) {
			s = append(s, datum.FailureMessage(m, t))
			return
		})
	}
}

func TestDivisible_uintValid(t *testing.T) {
	model := divisibleUintDummyModel{}

	for _, f := range []uint{60, 0, 20, 4e6} {
		model.x = f

		executeTest(t, model, nil)
	}
}
