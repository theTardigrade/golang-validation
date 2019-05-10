package handling

import "testing"

type oddIntDummyModel struct {
	x int `validation:"odd"`
}

func TestOdd_intInvalid(t *testing.T) {
	model := oddIntDummyModel{}

	for _, f := range []int{-8, -4, -2, 0, 2, 4, 8, 1e9} {
		model.x = f
		executeTest(t, model, 1)
	}
}

func TestOdd_intValid(t *testing.T) {
	model := oddIntDummyModel{}

	for _, f := range []int{-9, -3, -1, 1, 3, 9} {
		model.x = f
		executeTest(t, model, 0)
	}
}

type oddUintDummyModel struct {
	x uint `validation:"odd"`
}

func TestOdd_uintInvalid(t *testing.T) {
	model := oddUintDummyModel{}

	for _, f := range []uint{0, 2, 4, 8, 1e9} {
		model.x = f
		executeTest(t, model, 1)
	}
}

func TestOdd_uintValid(t *testing.T) {
	model := oddUintDummyModel{}

	for _, f := range []uint{1, 3, 9} {
		model.x = f
		executeTest(t, model, 0)
	}
}
