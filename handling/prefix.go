package handling

import (
	"reflect"
	"strings"

	"github.com/theTardigrade/golang-validation/data"
)

func init() {
	addHandler("prefix", prefixDatum{})
}

type prefixDatum struct{}

func (d prefixDatum) Test(m *data.Main, t *data.Tag) (success bool, err error) {
	switch m.FieldKind {
	case reflect.String:
		success, err = d.testString(m, t)
	default:
		err = ErrUnexpectedType
	}

	return
}

func (d prefixDatum) testString(m *data.Main, t *data.Tag) (success bool, err error) {
	s := m.FieldValue.String()
	l := len(s)

	success = (l == 0 && !m.ContainsTagKey("required")) || strings.HasPrefix(s, t.Value)
	return
}

func (d prefixDatum) FailureMessage(m *data.Main, t *data.Tag) string {
	return m.FormattedFieldName + ` must begin with "` + t.Value + `".`
}
