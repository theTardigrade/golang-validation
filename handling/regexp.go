package handling

import (
	"reflect"
	"regexp"

	"github.com/theTardigrade/golang-validation/data"
)

func init() {
	addHandler("regexp", regexpDatum{})
}

type regexpDatum struct{}

func (d regexpDatum) Test(m *data.Main, t *data.Tag) (success bool, err error) {
	switch m.FieldKind {
	case reflect.String:
		success, err = d.testString(m, t)
	default:
		err = ErrUnexpectedType
	}

	return
}

func (d regexpDatum) testString(m *data.Main, t *data.Tag) (success bool, err error) {
	s := m.FieldValue.String()
	v := t.Value
	r, err := regexp.Compile(v)
	if err != nil {
		return
	}

	success = r.MatchString(s)
	return
}

func (d regexpDatum) FailureMessage(m *data.Main, t *data.Tag) string {
	return m.FormattedFieldName + ` must match a standard format.`
}
