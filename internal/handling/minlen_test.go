package handling

import (
	"testing"
)

type minlenStringDummyModel struct {
	x string `validation:"minlen=4"`
	y string `validation:"required,minlen=4"`
}

func TestMinlen_stringInvalidEmpty(t *testing.T) {
	model := minlenStringDummyModel{}
	executeTest(t, model, 2)
}

func TestMinlen_stringInvalidMin(t *testing.T) {
	model := minlenStringDummyModel{y: "the"}
	executeTest(t, model, 1)
}

func TestMinlen_stringValid(t *testing.T) {
	model := minlenStringDummyModel{y: "these"}
	executeTest(t, model, 0)
}

type minlenSliceDummyModel struct {
	x []string `validation:"minlen=4"`
}

func TestMinlen_sliceInvalidEmpty(t *testing.T) {
	model := minlenSliceDummyModel{}
	executeTest(t, model, 1)
}

func TestMinlen_sliceInvalidValue(t *testing.T) {
	model := minlenSliceDummyModel{
		[]string{"a", "b"},
	}
	executeTest(t, model, 1)
}

func TestMinlen_sliceValid(t *testing.T) {
	model := minlenSliceDummyModel{
		[]string{"a", "b", "c", "d"},
	}
	executeTest(t, model, 0)

	model.x = []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	executeTest(t, model, 0)
}
