package handling

import (
	"math"
	"testing"

	"github.com/theTardigrade/golang-validation/internal/data"
)

type minIntDummyModel struct {
	x int `validation:"min=-20"`
}

func TestMin_intInvalid(t *testing.T) {
	model := minIntDummyModel{}
	datum := minDatum{}

	for _, f := range []int{-21, -22, -2e6, math.MinInt32} {
		model.x = f

		executeTest(t, model, func(m *data.Main, t *data.Tag) (s []string) {
			s = append(s, datum.FailureMessage(m, t))
			return
		})
	}
}

func TestMin_intValid(t *testing.T) {
	model := minIntDummyModel{}

	for _, f := range []int{-20, -19, 0, 4, 20, math.MaxInt32} {
		model.x = f

		executeTest(t, model, nil)
	}
}

type minUintDummyModel struct {
	x uint `validation:"min=20"`
}

func TestMin_uintInvalid(t *testing.T) {
	model := minUintDummyModel{}
	datum := minDatum{}

	for _, f := range []uint{0, 1, 5, 8, 19} {
		model.x = f

		executeTest(t, model, func(m *data.Main, t *data.Tag) (s []string) {
			s = append(s, datum.FailureMessage(m, t))
			return
		})
	}
}

func TestMin_uintValid(t *testing.T) {
	model := minUintDummyModel{}

	for _, f := range []uint{21, 22, 25, 999, math.MaxUint32} {
		model.x = f

		executeTest(t, model, nil)
	}
}

type minFloat64DummyModel struct {
	x float64 `validation:"min=1.8"`
}

func TestMin_float64Invalid(t *testing.T) {
	model := minFloat64DummyModel{}
	datum := minDatum{}

	for _, f := range []float64{1.799, 1.7, 1, 0, math.SmallestNonzeroFloat64} {
		model.x = f

		executeTest(t, model, func(m *data.Main, t *data.Tag) (s []string) {
			s = append(s, datum.FailureMessage(m, t))
			return
		})
	}
}

func TestMin_float64Valid(t *testing.T) {
	model := minFloat64DummyModel{}

	for _, f := range []float64{25, 2, 1.81, 1.801, math.MaxFloat64} {
		model.x = f

		executeTest(t, model, nil)
	}
}
