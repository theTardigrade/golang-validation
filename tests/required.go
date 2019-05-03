package tests

import (
	"reflect"

	"github.com/theTardigrade/validation/data"
)

func Required(m *data.Main) error {
	switch kind := m.Field.Type.Kind(); kind {
	case reflect.String:
		if len(m.FieldValue.String()) == 0 {
			m.SetFailure(m.FormattedFieldName + " required.")
		}
		/*
			case reflect.Ptr, reflect.Interface:
						for {
							value := m.FieldValue.Elem()
							m.FieldValue = &value
							kind = value.Kind()

							if kind != reflect.Ptr && kind != reflect.Interface {
								break
							}
						}
						fallthrough
					case reflect.Struct:
		*/
	default:
		return ErrUnexpectedType
	}

	return nil
}
