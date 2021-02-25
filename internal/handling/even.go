package handling

import (
	"reflect"

	"github.com/theTardigrade/golang-validation/internal/data"
)

func init() {
	addHandler("even", evenDatum{})
}

type evenDatum struct{}

func (d evenDatum) Test(m *data.Main, t *data.Tag) (success bool, err error) {
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

func (d evenDatum) testInts(m *data.Main, t *data.Tag) (success bool, err error) {
	success = m.FieldValue.Int()%2 == 0
	return
}

func (d evenDatum) testUints(m *data.Main, t *data.Tag) (success bool, err error) {
	success = m.FieldValue.Uint()%2 == 0
	return
}

func (d evenDatum) FailureMessage(m *data.Main, t *data.Tag) string {
	return m.FormattedFieldName + " must be even (divisible by two)."
}
