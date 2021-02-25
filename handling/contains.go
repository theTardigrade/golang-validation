package handling

import (
	"reflect"
	"strings"

	"github.com/theTardigrade/golang-validation/data"
)

func init() {
	addHandler("contains", containsDatum{})
}

type containsDatum struct{}

func (d containsDatum) Test(m *data.Main, t *data.Tag) (success bool, err error) {
	switch m.FieldKind {
	case reflect.String:
		success, err = d.testString(m, t)
	default:
		err = ErrUnexpectedType
	}

	return
}

func (d containsDatum) testString(m *data.Main, t *data.Tag) (success bool, err error) {
	s := m.FieldValue.String()
	l := len(s)

	success = (l == 0 && !m.ContainsTagKey("required")) || strings.Contains(s, t.Value)
	return
}

func (d containsDatum) FailureMessage(m *data.Main, t *data.Tag) string {
	return m.FormattedFieldName + ` must include "` + t.Value + `".`
}
