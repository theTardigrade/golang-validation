package handling

import "testing"

type suffixStringDummyModel struct {
	x string `validation:"suffix=s"`
	y string `validation:"required,suffix=s"`
}

func TestSuffix_stringInvalidEmpty(t *testing.T) {
	model := suffixStringDummyModel{}
	executeTest(t, model, 2)
}

func TestSuffix_stringInvalidValue(t *testing.T) {
	model := suffixStringDummyModel{x: "dog", y: "cat"}
	executeTest(t, model, 2)
}

func TestSuffix_stringValid(t *testing.T) {
	model := suffixStringDummyModel{x: "dogs", y: "cats"}
	executeTest(t, model, 0)
}
