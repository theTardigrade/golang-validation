package handling

import (
	"reflect"

	"github.com/theTardigrade/golang-validation/data"
)

func init() {
	addHandler("odd", oddDatum{})
}

type oddDatum struct{}

func (d oddDatum) Test(m *data.Main, t *data.Tag) (success bool, err error) {
	switch m.FieldKind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		success, err = d.testInts(m, t)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		success, err = d.testUints(m, t)
	default:
		err = ErrUnexpectedType
	}

	return
}

func (d oddDatum) testInts(m *data.Main, t *data.Tag) (success bool, err error) {
	success = m.FieldValue.Int()%2 != 0
	return
}

func (d oddDatum) testUints(m *data.Main, t *data.Tag) (success bool, err error) {
	success = m.FieldValue.Uint()%2 != 0
	return
}

func (d oddDatum) FailureMessage(m *data.Main, t *data.Tag) string {
	return m.FormattedFieldName + " must be odd (indivisible by two)."
}
