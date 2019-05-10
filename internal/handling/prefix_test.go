package handling

import "testing"

type prefixStringDummyModel struct {
	x string `validation:"prefix=the "`
	y string `validation:"required,prefix=the "`
}

func TestPrefix_stringInvalidEmpty(t *testing.T) {
	model := prefixStringDummyModel{}
	executeTest(t, model, 2)
}

func TestPrefix_stringInvalidValue(t *testing.T) {
	model := prefixStringDummyModel{x: "a thing", y: "a thing"}
	executeTest(t, model, 2)
}

func TestPrefix_stringValid(t *testing.T) {
	model := prefixStringDummyModel{y: "the things"}
	executeTest(t, model, 0)
}
