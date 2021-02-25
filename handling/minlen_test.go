package handling

import (
	"testing"

	"github.com/theTardigrade/golang-validation/data"
)

type minlenStringDummyModel struct {
	x string `validation:"minlen=4"`
	y string `validation:"required,minlen=4"`
}

func TestMinlen_stringInvalid(t *testing.T) {
	model := minlenStringDummyModel{}
	minlenDatum := minlenDatum{}
	requiredDatum := requiredDatum{}

	for _, v := range [][2]string{
		[...]string{"", ""},
		[...]string{"at", "abc"},
	} {
		model.x = v[0]
		model.y = v[1]

		executeTest(t, model, func(m *data.Main, t *data.Tag) (s []string) {
			switch fieldValue := m.FieldValue.String(); t.Key {
			case "minlen":
				if len(fieldValue) > 0 || m.ContainsTagKey("required") {
					s = append(s, minlenDatum.FailureMessage(m, t))
				}
			case "required":
				if fieldValue == "" {
					s = append(s, requiredDatum.FailureMessage(m, t))
				}
			}
			return
		})
	}
}

func TestMinlen_stringValid(t *testing.T) {
	model := minlenStringDummyModel{}

	for _, v := range [][2]string{
		[...]string{"", "test"},
		[...]string{"test", "test"},
	} {
		model.x = v[0]
		model.y = v[1]

		executeTest(t, model, nil)
	}
}

type minlenSliceDummyModel struct {
	x []string `validation:"minlen=4"`
}

func TestMinlen_sliceInvalid(t *testing.T) {
	model := minlenSliceDummyModel{}
	datum := minlenDatum{}

	for _, s := range [][]string{
		[]string{},
		[]string{"a"},
		[]string{"the", "this", "these"},
	} {
		model.x = s

		executeTest(t, model, func(m *data.Main, t *data.Tag) (s []string) {
			s = append(s, datum.FailureMessage(m, t))
			return
		})
	}
}

func TestMinlen_sliceValid(t *testing.T) {
	model := minlenSliceDummyModel{}

	for _, s := range [][]string{
		[]string{"a", "b", "c", "d", "e"},
		[]string{"0", "1", "2", "3", "4", "5", "6"},
	} {
		model.x = s

		executeTest(t, model, nil)
	}
}
