package handling

import (
	"testing"

	"github.com/theTardigrade/validation/internal/data"
)

type emailStringDummyModel struct {
	x string `validation:"email"`
	y string `validation:"required,email"`
}

func TestEmail_stringInvalid(t *testing.T) {
	model := emailStringDummyModel{}
	emailDatum := emailDatum{}
	requiredDatum := requiredDatum{}

	for _, v := range [][2]string{
		[...]string{"", ""},
		[...]string{"", "x.com"},
		[...]string{"x.com", "x.com"},
	} {
		model.x = v[0]
		model.y = v[1]

		executeTest(t, model, func(m *data.Main, t *data.Tag) (s []string) {
			switch fieldValue := m.FieldValue.String(); t.Key {
			case "email":
				if len(fieldValue) > 0 || m.ContainsTagKey("required") {
					s = append(s, emailDatum.FailureMessage(m, t))
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

func TestEmail_stringValid(t *testing.T) {
	model := emailStringDummyModel{}

	for _, v := range [][2]string{
		[...]string{"", "x@x.com"},
		[...]string{"x@x.com", "x.test@test.com"},
	} {
		model.x = v[0]
		model.y = v[1]

		executeTest(t, model, nil)
	}
}
