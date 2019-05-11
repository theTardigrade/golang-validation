package handling

import (
	"reflect"
	"regexp"

	"github.com/theTardigrade/validation/internal/data"
)

func init() {
	addHandler("email", emailDatum{})
}

type emailDatum struct{}

var (
	emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

func (d emailDatum) Test(m *data.Main, t *data.Tag) (success bool, err error) {
	switch m.FieldKind {
	case reflect.String:
		success, err = d.testString(m, t)
	default:
		err = ErrUnexpectedType
	}

	return
}

func (d emailDatum) testString(m *data.Main, t *data.Tag) (success bool, err error) {
	l := len(m.FieldValue.String())
	success = (l == 0 && !m.ContainsTagKey("required")) || emailRegexp.MatchString(m.FieldValue.String())
	return
}

func (d emailDatum) FailureMessage(m *data.Main, t *data.Tag) string {
	return m.FormattedFieldName + " not recognised as valid."
}
