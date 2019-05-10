package handling

import (
	"testing"
)

type maxlenStringDummyModel struct {
	x string `validation:"maxlen=4"`
	y string `validation:"required,maxlen=4"`
}

func TestMaxlen_stringInvalidEmpty(t *testing.T) {
	model := maxlenStringDummyModel{}
	executeTest(t, model, 1)
}

func TestMaxlen_stringInvalidValue(t *testing.T) {
	model := maxlenStringDummyModel{y: "these"}
	executeTest(t, model, 1)
}

func TestMaxlen_stringValid(t *testing.T) {
	model := maxlenStringDummyModel{y: "the"}
	executeTest(t, model, 0)
}

type maxlenSliceDummyModel struct {
	x []string `validation:"maxlen=4"`
}

func TestMaxlen_sliceInvalidEmpty(t *testing.T) {
	model := maxlenSliceDummyModel{}
	executeTest(t, model, 0)
}

func TestMaxlen_sliceInvalidValue(t *testing.T) {
	model := maxlenSliceDummyModel{
		[]string{"a", "b", "c", "d", "e"},
	}
	executeTest(t, model, 1)
}

func TestMaxlen_sliceValid(t *testing.T) {
	model := maxlenSliceDummyModel{
		[]string{"a", "b", "c", "d"},
	}
	executeTest(t, model, 0)

	model.x = []string{"a", "b"}
	executeTest(t, model, 0)
}
