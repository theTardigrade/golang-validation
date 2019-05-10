package handling

import (
	"reflect"
	"strconv"

	"github.com/theTardigrade/validation/internal/data"
)

func init() {
	addHandler("min", min)
}

func min(m *data.Main, t *data.Tag) error {
	var failure bool

	switch m.Field.Type.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		tagValueInt, err := strconv.ParseInt(t.Value, 10, 64)
		if err != nil {
			return err
		}

		if m.FieldValue.Int() < tagValueInt {
			failure = true
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		tagValueUint, err := strconv.ParseUint(t.Value, 10, 64)
		if err != nil {
			return err
		}

		if m.FieldValue.Uint() < tagValueUint {
			failure = true
		}
	case reflect.Float32, reflect.Float64:
		tagValueFloat, err := strconv.ParseFloat(t.Value, 64)
		if err != nil {
			return err
		}

		if m.FieldValue.Float() < tagValueFloat {
			failure = true
		}
	default:
		return ErrUnexpectedType
	}

	if failure {
		m.SetFailure(t, m.FormattedFieldName+" must be at least "+t.Value+".")
	}

	return nil
}
