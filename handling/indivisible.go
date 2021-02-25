package handling

import (
	"reflect"
	"strconv"

	"github.com/theTardigrade/golang-validation/data"
)

func init() {
	addHandler("indivisible", indivisibleDatum{})
}

type indivisibleDatum struct{}

func (d indivisibleDatum) Test(m *data.Main, t *data.Tag) (success bool, err error) {
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

func (d indivisibleDatum) testInts(m *data.Main, t *data.Tag) (success bool, err error) {
	tagValueInt, err := strconv.ParseInt(t.Value, 10, 64)
	if err == nil && m.FieldValue.Int()%tagValueInt != 0 {
		success = true
	}

	return
}

func (d indivisibleDatum) testUints(m *data.Main, t *data.Tag) (success bool, err error) {
	tagValueUint, err := strconv.ParseUint(t.Value, 10, 64)
	if err == nil && m.FieldValue.Uint()%tagValueUint != 0 {
		success = true
	}

	return
}

func (d indivisibleDatum) FailureMessage(m *data.Main, t *data.Tag) (value string) {
	return m.FormattedFieldName + " cannot be divisible by " + t.Value + "."
}
