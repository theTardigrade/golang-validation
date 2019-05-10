package handling

import (
	"reflect"

	"github.com/theTardigrade/validation/internal/data"
)

func init() {
	addHandler("required", required)
}

func required(m *data.Main, t *data.Tag) error {
	switch kind := m.Field.Type.Kind(); kind {
	case reflect.String:
		if len(m.FieldValue.String()) == 0 {
			m.SetFailure(t, m.FormattedFieldName+" required.")
		}
	case reflect.Ptr:
		for value := *m.FieldValue; value.Kind() == reflect.Ptr; value = value.Elem() {
			if value.IsNil() {
				m.SetFailure(t, m.FormattedFieldName+" required.")
				break
			}
		}
	default:
		return ErrUnexpectedType
	}

	return nil
}
