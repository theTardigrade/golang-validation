package handling

import (
	"testing"
)

type numberIntDummyModel struct {
	x int `validation:"min=-4,max=8,divisible=2"`
}

func TestNumber_intInvalidMin(t *testing.T) {
	model := numberIntDummyModel{-6}
	executeTest(t, model, 1)
}

func TestNumber_intInvalidMax(t *testing.T) {
	model := numberIntDummyModel{10}
	executeTest(t, model, 1)
}

func TestNumber_intInvalidDivisible(t *testing.T) {
	model := numberIntDummyModel{}

	for _, n := range []int{5, 7} {
		model.x = n
		executeTest(t, model, 1)
	}
}

func TestNumber_intValid(t *testing.T) {
	model := numberIntDummyModel{}

	for _, n := range []int{4, 6, 8} {
		model.x = n
		executeTest(t, model, 0)
	}
}

type numberUintDummyModel struct {
	x uint `validation:"min=4,max=8,divisible=2"`
}

func TestNumber_uintInvalidMin(t *testing.T) {
	model := numberUintDummyModel{}
	executeTest(t, model, 1)
}

func TestNumber_uintInvalidMax(t *testing.T) {
	model := numberUintDummyModel{10}
	executeTest(t, model, 1)
}

func TestNumber_uintInvalidDivisible(t *testing.T) {
	model := numberUintDummyModel{}

	for _, n := range []uint{5, 7} {
		model.x = n
		executeTest(t, model, 1)
	}
}

func TestNumber_uintValid(t *testing.T) {
	model := numberUintDummyModel{}

	for _, n := range []uint{4, 6, 8} {
		model.x = n
		executeTest(t, model, 0)
	}
}

type numberFloat64DummyModel struct {
	x float64 `validation:"min=1.4,max=1.8"`
}

func TestNumber_float64InvalidMin(t *testing.T) {
	model := numberFloat64DummyModel{}

	for _, f := range []float64{-4, 0, 1.39, 1.399} {
		model.x = f
		executeTest(t, model, 1)
	}
}

func TestNumber_float64InvalidMax(t *testing.T) {
	model := numberFloat64DummyModel{}

	for _, f := range []float64{25, 2, 1.81, 1.801} {
		model.x = f
		executeTest(t, model, 1)
	}
}

func TestNumber_float64Valid(t *testing.T) {
	model := numberFloat64DummyModel{}

	for _, f := range []float64{1.4, 1.6, 1.8} {
		model.x = f
		executeTest(t, model, 0)
	}
}
