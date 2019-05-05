package handling

import (
	"testing"
)

type numberIntDummyModel struct {
	x int `validation:"min=-4,max=20,even,indivisible=3"`
}

func TestNumber_intInvalidMin(t *testing.T) {
	model := numberIntDummyModel{-8}
	executeTest(t, model, 1)
}

func TestNumber_intInvalidMax(t *testing.T) {
	model := numberIntDummyModel{22}
	executeTest(t, model, 1)
}

func TestNumber_intInvalidDivisible(t *testing.T) {
	model := numberIntDummyModel{}

	for _, n := range []int{5, 7} {
		model.x = n
		executeTest(t, model, 1)
	}
}

func TestNumber_intInvalidIndivisible(t *testing.T) {
	model := numberIntDummyModel{}

	for _, n := range []int{6, 12, 18} {
		model.x = n
		executeTest(t, model, 1)
	}
}

func TestNumber_intValid(t *testing.T) {
	model := numberIntDummyModel{}

	for _, n := range []int{4, 8, 20} {
		model.x = n
		executeTest(t, model, 0)
	}
}

type numberUintDummyModel struct {
	x uint `validation:"min=4,max=18,divisible=2,indivisible=4"`
}

func TestNumber_uintInvalidMin(t *testing.T) {
	model := numberUintDummyModel{2}
	executeTest(t, model, 1)
}

func TestNumber_uintInvalidMax(t *testing.T) {
	model := numberUintDummyModel{22}
	executeTest(t, model, 1)
}

func TestNumber_uintInvalidDivisible(t *testing.T) {
	model := numberUintDummyModel{}

	for _, n := range []uint{5, 7} {
		model.x = n
		executeTest(t, model, 1)
	}
}

func TestNumber_uintInvalidIndivisible(t *testing.T) {
	model := numberUintDummyModel{}

	for _, n := range []uint{4, 8, 12, 16} {
		model.x = n
		executeTest(t, model, 1)
	}
}

func TestNumber_uintValid(t *testing.T) {
	model := numberUintDummyModel{}

	for _, n := range []uint{6, 10, 14, 18} {
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
