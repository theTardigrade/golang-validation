package handling

import (
	"reflect"

	"github.com/theTardigrade/validation/internal/data"
)

func required(m *data.Main) error {
	switch kind := m.Field.Type.Kind(); kind {
	case reflect.String:
		if len(m.FieldValue.String()) == 0 {
			m.SetFailure(m.FormattedFieldName + " required.")
		}
	case reflect.Ptr:
		for value := *m.FieldValue; value.Kind() == reflect.Ptr; value = value.Elem() {
			if value.IsNil() {
				m.SetFailure(m.FormattedFieldName + " required.")
				break
			}
		}
	default:
		return ErrUnexpectedType
	}

	return nil
}

func init() {
	addHandler("required", required)
}
