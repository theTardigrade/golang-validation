package handling

import (
	"reflect"

	"github.com/theTardigrade/validation/internal/data"
)

func init() {
	addHandler("required", requiredDatum{})
}

type requiredDatum struct{}

func (d requiredDatum) Test(m *data.Main, t *data.Tag) (success bool, err error) {
	switch m.FieldKind {
	case reflect.String:
		success, err = d.testString(m, t)
	case reflect.Ptr:
		success, err = d.testPointer(m, t)
	default:
		err = ErrUnexpectedType
	}

	return
}

func (d requiredDatum) testString(m *data.Main, t *data.Tag) (success bool, err error) {
	success = len(m.FieldValue.String()) != 0
	return
}

func (d requiredDatum) testPointer(m *data.Main, t *data.Tag) (success bool, err error) {
	success = true

	for value := *m.FieldValue; value.Kind() == reflect.Ptr; value = value.Elem() {
		if value.IsNil() {
			success = false
			break
		}
	}

	return
}

func (d requiredDatum) FailureMessage(m *data.Main, t *data.Tag) string {
	return m.FormattedFieldName + " required."
}
