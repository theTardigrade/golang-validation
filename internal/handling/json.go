package handling

import (
	"encoding/json"
	"reflect"

	"github.com/theTardigrade/validation/internal/data"
)

func init() {
	addHandler("json", jsonDatum{})
}

type jsonDatum struct{}

func (d jsonDatum) Test(m *data.Main, t *data.Tag) (success bool, err error) {
	switch m.FieldKind {
	case reflect.String:
		success, err = d.testString(m, t)
	default:
		err = ErrUnexpectedType
	}

	return
}

func (d jsonDatum) testString(m *data.Main, t *data.Tag) (success bool, err error) {
	var x map[string]interface{}
	s := m.FieldValue.String()
	l := len(s)

	success = (l == 0 && !m.ContainsTagKey("required")) || json.Unmarshal([]byte(s), &x) == nil
	return
}

func (d jsonDatum) FailureMessage(m *data.Main, t *data.Tag) string {
	return m.FormattedFieldName + " not recognised as JSON."
}
