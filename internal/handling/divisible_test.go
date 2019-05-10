package handling

import (
	"testing"
)

type divisibleIntDummyModel struct {
	x int `validation:"divisible=-20"`
}

func TestDivisible_intInvalid(t *testing.T) {
	model := divisibleIntDummyModel{}

	for _, f := range []int{-9, 1, 9, 21, 44} {
		model.x = f
		executeTest(t, model, 1)
	}
}

func TestDivisible_intValid(t *testing.T) {
	model := divisibleIntDummyModel{}

	for _, f := range []int{60, -20, 0, 20, 4e6} {
		model.x = f
		executeTest(t, model, 0)
	}
}

type divisibleUintDummyModel struct {
	x uint `validation:"divisible=20"`
}

func TestDivisible_uintInvalid(t *testing.T) {
	model := divisibleUintDummyModel{}

	for _, f := range []uint{1, 9, 21, 44} {
		model.x = f
		executeTest(t, model, 1)
	}
}

func TestDivisible_uintValid(t *testing.T) {
	model := divisibleUintDummyModel{}

	for _, f := range []uint{60, 0, 20, 4e6} {
		model.x = f
		executeTest(t, model, 0)
	}
}
