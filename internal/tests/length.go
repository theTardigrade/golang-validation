package tests

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/theTardigrade/validation/internal/data"
)

func MinLen(m *data.Main) error {
	tagValueInt, err := strconv.Atoi(m.CurrentTag.Value)
	if err != nil {
		return err
	}

	fmt.Println(m.FormattedFieldName, m.CurrentTag.Value, tagValueInt)

	switch m.Field.Type.Kind() {
	case reflect.Slice, reflect.Array, reflect.Map:
		if m.FieldValue.Len() < tagValueInt {
			m.SetFailure(m.FormattedFieldName + " must contain at least " + m.CurrentTag.Value + " values.")
		}
	case reflect.String:
		if l := len(m.FieldValue.String()); (l > 0 || m.ContainsTagKey("required")) && l < tagValueInt {
			m.SetFailure(m.FormattedFieldName + " must be at least " + m.CurrentTag.Value + " characters long.")
		}
	default:
		return ErrUnexpectedType
	}

	return nil
}

func MaxLen(m *data.Main) error {
	tagValueInt, err := strconv.Atoi(m.CurrentTag.Value)
	if err != nil {
		return err
	}

	switch m.Field.Type.Kind() {
	case reflect.Slice, reflect.Array, reflect.Map:
		if m.FieldValue.Len() > tagValueInt {
			m.SetFailure(m.FormattedFieldName + " cannot contain more than " + m.CurrentTag.Value + " values.")
		}
	case reflect.String:
		if len(m.FieldValue.String()) > tagValueInt {
			m.SetFailure(m.FormattedFieldName + " cannot be more than " + m.CurrentTag.Value + " characters long.")
		}
	default:
		return ErrUnexpectedType
	}

	return nil
}
