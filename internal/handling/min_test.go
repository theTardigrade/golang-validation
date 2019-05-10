package handling

import (
	"math"
	"testing"
)

type minIntDummyModel struct {
	x int `validation:"min=-20"`
}

func TestMin_intInvalid(t *testing.T) {
	model := minIntDummyModel{}

	for _, f := range []int{-500, -50, -22, -21, math.MinInt32} {
		model.x = f
		executeTest(t, model, 1)
	}
}

func TestMin_intValid(t *testing.T) {
	model := minIntDummyModel{}

	for _, f := range []int{-20, -19, -18, 0, 18, 19, 20, math.MaxInt32} {
		model.x = f
		executeTest(t, model, 0)
	}
}

type minUintDummyModel struct {
	x uint `validation:"min=20"`
}

func TestMin_uintInvalid(t *testing.T) {
	model := minUintDummyModel{}

	for _, f := range []uint{0, 1, 2, 19} {
		model.x = f
		executeTest(t, model, 1)
	}
}

func TestMin_uintValid(t *testing.T) {
	model := minUintDummyModel{}

	for _, f := range []uint{20, 21, 22, 25, 500, math.MaxUint32} {
		model.x = f
		executeTest(t, model, 0)
	}
}

type minFloat64DummyModel struct {
	x float64 `validation:"min=1.8"`
}

func TestMin_float64Invalid(t *testing.T) {
	model := minFloat64DummyModel{}

	for _, f := range []float64{-1.799, -1.79, -1.7, -1, 0, 1, 1.7, 1.79, 1.799, math.SmallestNonzeroFloat64} {
		model.x = f
		executeTest(t, model, 1)
	}
}

func TestMin_float64Valid(t *testing.T) {
	model := minFloat64DummyModel{}

	for _, f := range []float64{1.801, 1.81, 1.9, 2, 500, math.MaxFloat64} {
		model.x = f
		executeTest(t, model, 0)
	}
}
