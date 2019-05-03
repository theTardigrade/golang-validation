package handling

import (
	"reflect"
	"strconv"

	"github.com/theTardigrade/validation/internal/data"
)

func Min(m *data.Main) error {
	var failure bool

	switch m.Field.Type.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		tagValueInt, err := strconv.ParseInt(m.CurrentTag.Value, 10, 64)
		if err != nil {
			return err
		}

		if m.FieldValue.Int() < tagValueInt {
			failure = true
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		tagValueUint, err := strconv.ParseUint(m.CurrentTag.Value, 10, 64)
		if err != nil {
			return err
		}

		if m.FieldValue.Uint() < tagValueUint {
			failure = true
		}
	case reflect.Float32, reflect.Float64:
		tagValueFloat, err := strconv.ParseFloat(m.CurrentTag.Value, 64)
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
		m.SetFailure(m.FormattedFieldName + " must be at least " + m.CurrentTag.Value + ".")
	}

	return nil
}

func Max(m *data.Main) error {
	var failure bool

	switch m.Field.Type.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		tagValueInt, err := strconv.ParseInt(m.CurrentTag.Value, 10, 64)
		if err != nil {
			return err
		}

		if m.FieldValue.Int() > tagValueInt {
			failure = true
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		tagValueUint, err := strconv.ParseUint(m.CurrentTag.Value, 10, 64)
		if err != nil {
			return err
		}

		if m.FieldValue.Uint() > tagValueUint {
			failure = true
		}
	case reflect.Float32, reflect.Float64:
		tagValueFloat, err := strconv.ParseFloat(m.CurrentTag.Value, 64)
		if err != nil {
			return err
		}

		if m.FieldValue.Float() > tagValueFloat {
			failure = true
		}
	default:
		return ErrUnexpectedType
	}

	if failure {
		m.SetFailure(m.FormattedFieldName + " cannot be greater than " + m.CurrentTag.Value + ".")
	}

	return nil
}
