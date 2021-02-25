package handling

import (
	"reflect"
	"strconv"

	"github.com/theTardigrade/golang-validation/data"
)

func init() {
	addHandler("maxlen", maxlenDatum{})
}

type maxlenDatum struct{}

func (d maxlenDatum) Test(m *data.Main, t *data.Tag) (success bool, err error) {
	switch m.FieldKind {
	case reflect.Slice, reflect.Array, reflect.Map:
		success, err = d.testCollections(m, t)
	case reflect.String:
		success, err = d.testString(m, t)
	default:
		err = ErrUnexpectedType
	}

	return
}

func (d maxlenDatum) testCollections(m *data.Main, t *data.Tag) (success bool, err error) {
	tagValueInt, err := strconv.Atoi(t.Value)
	if err == nil {
		success = m.FieldValue.Len() <= tagValueInt
	}

	return
}

func (d maxlenDatum) testString(m *data.Main, t *data.Tag) (success bool, err error) {
	tagValueInt, err := strconv.Atoi(t.Value)
	if err == nil {
		success = len(m.FieldValue.String()) <= tagValueInt
	}

	return
}

func (d maxlenDatum) FailureMessage(m *data.Main, t *data.Tag) string {
	switch m.FieldKind {
	case reflect.Slice, reflect.Array, reflect.Map:
		return d.failureMessageCollections(m, t)
	case reflect.String:
		return d.failureMessageString(m, t)
	default:
		panic(ErrUnexpectedType)
	}
}

func (d maxlenDatum) failureMessageCollections(m *data.Main, t *data.Tag) string {
	return m.FormattedFieldName + " cannot contain more than " + t.Value + " values."
}

func (d maxlenDatum) failureMessageString(m *data.Main, t *data.Tag) string {
	return m.FormattedFieldName + " cannot be more than " + t.Value + " characters long."
}
