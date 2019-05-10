package handling

import (
	"reflect"
	"strings"

	"github.com/theTardigrade/validation/internal/data"
)

func init() {
	addHandler("suffix", suffix)
}

func suffix(m *data.Main, t *data.Tag) error {
	switch kind := m.Field.Type.Kind(); kind {
	case reflect.String:
		{
			s := m.FieldValue.String()
			l := len(s)

			if (l > 0 || m.ContainsTagKey("required")) && !strings.HasSuffix(s, t.Value) {
				m.SetFailure(t, m.FormattedFieldName+` must end with "`+t.Value+`".`)
			}
		}
	default:
		return ErrUnexpectedType
	}

	return nil
}
