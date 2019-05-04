package handling

import "testing"

type requiredStringDummyModel struct {
	x string `validation:"required"`
}

func TestRequired_stringInvalid(t *testing.T) {
	model := requiredStringDummyModel{}
	executeTest(t, model, 1)
}

func TestRequired_stringValid(t *testing.T) {
	model := requiredStringDummyModel{"x"}
	executeTest(t, model, 0)
}

type requiredPointerDummyModel struct {
	x *string `validation:"required"`
}

func TestRequired_pointerInvalid(t *testing.T) {
	model := requiredPointerDummyModel{}
	executeTest(t, model, 1)
}

func TestRequired_pointerValid(t *testing.T) {
	s := "x"
	model := requiredPointerDummyModel{&s}
	executeTest(t, model, 0)
}
