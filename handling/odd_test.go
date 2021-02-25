package handling

import (
	"testing"

	"github.com/theTardigrade/golang-validation/data"
)

type oddIntDummyModel struct {
	x int `validation:"odd"`
}

func TestOdd_intInvalid(t *testing.T) {
	model := oddIntDummyModel{}
	datum := oddDatum{}

	for _, f := range []int{60, -20, 0, 20, 4e6, 1e9} {
		model.x = f

		executeTest(t, model, func(m *data.Main, t *data.Tag) (s []string) {
			s = append(s, datum.FailureMessage(m, t))
			return
		})
	}
}

func TestOdd_intValid(t *testing.T) {
	model := oddIntDummyModel{}

	for _, f := range []int{-9, 1, 9, 21, 45, 9999, 1e9 - 1} {
		model.x = f

		executeTest(t, model, nil)
	}
}

type oddUintDummyModel struct {
	x uint `validation:"odd"`
}

func TestOdd_uintInvalid(t *testing.T) {
	model := oddUintDummyModel{}
	datum := oddDatum{}

	for _, f := range []uint{60, 0, 20, 4e6, 1e9} {
		model.x = f

		executeTest(t, model, func(m *data.Main, t *data.Tag) (s []string) {
			s = append(s, datum.FailureMessage(m, t))
			return
		})
	}
}

func TestOdd_uintValid(t *testing.T) {
	model := oddUintDummyModel{}

	for _, f := range []uint{1, 9, 21, 45, 999, 1e9 - 1} {
		model.x = f

		executeTest(t, model, nil)
	}
}
