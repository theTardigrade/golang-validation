package handling

import (
	"testing"
)

type lengthStringDummyModel struct {
	x string `validation:"minlen=4,maxlen=8"`
	y string `validation:"required,minlen=4,maxlen=8"`
}

func TestLength_stringInvalidEmpty(t *testing.T) {
	model := lengthStringDummyModel{}
	executeTest(t, model, 2)
}

func TestLength_stringInvalidMin(t *testing.T) {
	model := lengthStringDummyModel{y: "the"}
	executeTest(t, model, 1)
}

func TestLength_stringInvalidMax(t *testing.T) {
	model := lengthStringDummyModel{y: "the things"}
	executeTest(t, model, 1)
}

func TestLength_stringValid(t *testing.T) {
	model := lengthStringDummyModel{y: "these"}
	executeTest(t, model, 0)
}

type lengthSliceDummyModel struct {
	x []string `validation:"minlen=4,maxlen=8"`
}

func TestLength_sliceInvalidEmpty(t *testing.T) {
	model := lengthSliceDummyModel{}
	executeTest(t, model, 1)
}

func TestLength_sliceInvalidMin(t *testing.T) {
	model := lengthSliceDummyModel{
		[]string{"a", "b"},
	}
	executeTest(t, model, 1)
}
func TestLength_sliceInvalidMax(t *testing.T) {
	model := lengthSliceDummyModel{
		[]string{"a", "b", "c", "d", "e", "f", "g", "h", "i"},
	}
	executeTest(t, model, 1)
}

func TestLength_sliceValid(t *testing.T) {
	model := lengthSliceDummyModel{
		[]string{"a", "b", "c", "d"},
	}
	executeTest(t, model, 0)

	model.x = []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	executeTest(t, model, 0)
}
