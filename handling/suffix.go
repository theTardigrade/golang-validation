package handling

import (
	"reflect"
	"strings"

	"github.com/theTardigrade/golang-validation/data"
)

func init() {
	addHandler("suffix", suffixDatum{})
}

type suffixDatum struct{}

func (d suffixDatum) Test(m *data.Main, t *data.Tag) (success bool, err error) {
	switch m.FieldKind {
	case reflect.String:
		success, err = d.testString(m, t)
	default:
		err = ErrUnexpectedType
	}

	return
}

func (d suffixDatum) testString(m *data.Main, t *data.Tag) (success bool, err error) {
	s := m.FieldValue.String()
	l := len(s)

	success = (l == 0 && !m.ContainsTagKey("required")) || strings.HasSuffix(s, t.Value)
	return
}

func (d suffixDatum) FailureMessage(m *data.Main, t *data.Tag) string {
	return m.FormattedFieldName + ` must end with "` + t.Value + `".`
}
