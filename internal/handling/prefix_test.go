package handling

import (
	"testing"

	"github.com/theTardigrade/golang-validation/internal/data"
)

type prefixStringDummyModel struct {
	x string `validation:"prefix=the "`
	y string `validation:"required,prefix=a"`
}

func TestPrefix_stringInvalid(t *testing.T) {
	model := prefixStringDummyModel{}
	prefixDatum := prefixDatum{}
	requiredDatum := requiredDatum{}

	for _, v := range [][2]string{
		[...]string{"", ""},
		[...]string{"ape", "test"},
		[...]string{"test", ""},
	} {
		model.x = v[0]
		model.y = v[1]

		executeTest(t, model, func(m *data.Main, t *data.Tag) (s []string) {
			switch fieldValue := m.FieldValue.String(); t.Key {
			case "prefix":
				if len(fieldValue) > 0 || m.ContainsTagKey("required") {
					s = append(s, prefixDatum.FailureMessage(m, t))
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

func TestPrefix_stringValid(t *testing.T) {
	model := prefixStringDummyModel{}

	for _, v := range [][2]string{
		[...]string{"", "ape"},
		[...]string{"the ape", "aligator"},
		[...]string{"the aligator", "a cat"},
		[...]string{"the ", "a"},
	} {
		model.x = v[0]
		model.y = v[1]

		executeTest(t, model, nil)
	}
}
