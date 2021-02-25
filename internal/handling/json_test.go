package handling

import (
	"testing"

	"github.com/theTardigrade/golang-validation/internal/data"
)

type jsonStringDummyModel struct {
	x string `validation:"json"`
	y string `validation:"required,json"`
}

func TestJson_stringInvalid(t *testing.T) {
	model := jsonStringDummyModel{}
	jsonDatum := jsonDatum{}
	requiredDatum := requiredDatum{}

	for _, v := range [][2]string{
		[...]string{"", ""},
		[...]string{"", `test`},
		[...]string{"test", `test`},
		[...]string{"[]", `[]`},
	} {
		model.x = v[0]
		model.y = v[1]

		executeTest(t, model, func(m *data.Main, t *data.Tag) (s []string) {
			switch fieldValue := m.FieldValue.String(); t.Key {
			case "json":
				if len(fieldValue) > 0 || m.ContainsTagKey("required") {
					s = append(s, jsonDatum.FailureMessage(m, t))
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

func TestJson_stringValid(t *testing.T) {
	model := jsonStringDummyModel{}

	for _, v := range [][2]string{
		[...]string{"", `{}`},
		[...]string{`{"test":"test"}`, `{"x":99,"y":{"x":[1,2,3]}}`},
	} {
		model.x = v[0]
		model.y = v[1]

		executeTest(t, model, nil)
	}
}
