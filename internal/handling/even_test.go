package handling

import "testing"

type evenIntDummyModel struct {
	x int `validation:"even"`
}

func TestEven_intInvalid(t *testing.T) {
	model := evenIntDummyModel{}

	for _, f := range []int{-9, -3, -1, 1, 3, 9} {
		model.x = f
		executeTest(t, model, 1)
	}
}

func TestEven_intValid(t *testing.T) {
	model := evenIntDummyModel{}

	for _, f := range []int{-8, -4, -2, 0, 2, 4, 8, 1e9} {
		model.x = f
		executeTest(t, model, 0)
	}
}

type evenUintDummyModel struct {
	x uint `validation:"even"`
}

func TestEven_uintInvalid(t *testing.T) {
	model := evenUintDummyModel{}

	for _, f := range []uint{1, 3, 9} {
		model.x = f
		executeTest(t, model, 1)
	}
}

func TestEven_uintValid(t *testing.T) {
	model := evenUintDummyModel{}

	for _, f := range []uint{0, 2, 4, 8, 1e9} {
		model.x = f
		executeTest(t, model, 0)
	}
}
