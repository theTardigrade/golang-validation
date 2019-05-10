package handling

import (
	"math"
	"testing"
)

type maxIntDummyModel struct {
	x int `validation:"max=-20"`
}

func TestMax_intInvalid(t *testing.T) {
	model := maxIntDummyModel{}

	for _, f := range []int{-19, 0, 4, 20, math.MaxInt32} {
		model.x = f
		executeTest(t, model, 1)
	}
}

func TestMax_intValid(t *testing.T) {
	model := maxIntDummyModel{}

	for _, f := range []int{-20, -21, -22, math.MinInt32} {
		model.x = f
		executeTest(t, model, 0)
	}
}

type maxUintDummyModel struct {
	x uint `validation:"max=20"`
}

func TestMax_uintInvalid(t *testing.T) {
	model := maxUintDummyModel{}

	for _, f := range []uint{21, 22, 25, 999, math.MaxUint32} {
		model.x = f
		executeTest(t, model, 1)
	}
}

func TestMax_uintValid(t *testing.T) {
	model := maxUintDummyModel{}

	for _, f := range []uint{0, 1, 5, 8, 20} {
		model.x = f
		executeTest(t, model, 0)
	}
}

type maxFloat64DummyModel struct {
	x float64 `validation:"max=1.8"`
}

func TestMax_float64Invalid(t *testing.T) {
	model := maxFloat64DummyModel{}

	for _, f := range []float64{25, 2, 1.81, 1.801, math.MaxFloat64} {
		model.x = f
		executeTest(t, model, 1)
	}
}

func TestMax_float64Valid(t *testing.T) {
	model := maxFloat64DummyModel{}

	for _, f := range []float64{1.799, 1.7, 1, 0, math.SmallestNonzeroFloat64} {
		model.x = f
		executeTest(t, model, 0)
	}
}
