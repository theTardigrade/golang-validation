package handling

import (
	"testing"

	"github.com/theTardigrade/validation/internal/data"
)

type evenIntDummyModel struct {
	x int `validation:"even"`
}

func TestEven_intInvalid(t *testing.T) {
	model := evenIntDummyModel{}
	datum := evenDatum{}

	for _, f := range []int{-9, 1, 9, 21, 45, 9999, 1e9 - 1} {
		model.x = f

		executeTest(t, model, func(m *data.Main, t *data.Tag) (s []string) {
			s = append(s, datum.FailureMessage(m, t))
			return
		})
	}
}

func TestEven_intValid(t *testing.T) {
	model := evenIntDummyModel{}

	for _, f := range []int{60, -20, 0, 20, 4e6, 1e9} {
		model.x = f

		executeTest(t, model, nil)
	}
}

type evenUintDummyModel struct {
	x uint `validation:"even"`
}

func TestEven_uintInvalid(t *testing.T) {
	model := evenUintDummyModel{}
	datum := evenDatum{}

	for _, f := range []uint{1, 9, 21, 45, 999, 1e9 - 1} {
		model.x = f

		executeTest(t, model, func(m *data.Main, t *data.Tag) (s []string) {
			s = append(s, datum.FailureMessage(m, t))
			return
		})
	}
}

func TestEven_uintValid(t *testing.T) {
	model := evenUintDummyModel{}

	for _, f := range []uint{60, 0, 20, 4e6, 1e9} {
		model.x = f

		executeTest(t, model, nil)
	}
}
