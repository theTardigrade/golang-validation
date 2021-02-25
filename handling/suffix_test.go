package handling

import (
	"testing"

	"github.com/theTardigrade/golang-validation/data"
)

type suffixStringDummyModel struct {
	x string `validation:"suffix=s"`
	y string `validation:"required,suffix=es"`
}

func TestSuffix_stringInvalid(t *testing.T) {
	model := suffixStringDummyModel{}
	suffixDatum := suffixDatum{}
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
			case "suffix":
				if len(fieldValue) > 0 || m.ContainsTagKey("required") {
					s = append(s, suffixDatum.FailureMessage(m, t))
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

func TestSuffix_stringValid(t *testing.T) {
	model := suffixStringDummyModel{}

	for _, v := range [][2]string{
		[...]string{"", "matrices"},
		[...]string{"tests", "faces"},
		[...]string{"s", "es"},
	} {
		model.x = v[0]
		model.y = v[1]

		executeTest(t, model, nil)
	}
}
