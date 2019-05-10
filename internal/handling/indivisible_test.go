package handling

import (
	"testing"
)

type indivisibleIntDummyModel struct {
	x int `validation:"indivisible=-20"`
}

func TestIndivisible_intInvalid(t *testing.T) {
	model := indivisibleIntDummyModel{}

	for _, f := range []int{60, -20, 0, 20, 4e6} {
		model.x = f
		executeTest(t, model, 1)
	}
}

func TestIndivisible_intValid(t *testing.T) {
	model := indivisibleIntDummyModel{}

	for _, f := range []int{-9, 1, 9, 21, 44} {
		model.x = f
		executeTest(t, model, 0)
	}
}

type indivisibleUintDummyModel struct {
	x uint `validation:"indivisible=20"`
}

func TestIndivisible_uintInvalid(t *testing.T) {
	model := indivisibleUintDummyModel{}

	for _, f := range []uint{60, 0, 20, 4e6} {
		model.x = f
		executeTest(t, model, 1)
	}
}

func TestIndivisible_uintValid(t *testing.T) {
	model := indivisibleUintDummyModel{}

	for _, f := range []uint{1, 9, 21, 44} {
		model.x = f
		executeTest(t, model, 0)
	}
}
