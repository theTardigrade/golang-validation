package handling

import (
	"reflect"

	"github.com/theTardigrade/validation/internal/data"
)

func init() {
	addHandler("even", even)
}

func even(m *data.Main, t *data.Tag) error {
	var failure bool

	switch m.Field.Type.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if m.FieldValue.Int()%2 != 0 {
			failure = true
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if m.FieldValue.Uint()%2 != 0 {
			failure = true
		}
	default:
		return ErrUnexpectedType
	}

	if failure {
		m.SetFailure(t, m.FormattedFieldName+" must be an even number.")
	}

	return nil
}
