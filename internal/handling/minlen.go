package handling

import (
	"reflect"
	"strconv"

	"github.com/theTardigrade/validation/internal/data"
)

func init() {
	addHandler("minlen", minlenDatum{})
}

type minlenDatum struct{}

func (d minlenDatum) Test(m *data.Main, t *data.Tag) (success bool, err error) {
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

func (d minlenDatum) testCollections(m *data.Main, t *data.Tag) (success bool, err error) {
	tagValueInt, err := strconv.Atoi(t.Value)
	if err == nil {
		success = m.FieldValue.Len() >= tagValueInt
	}

	return
}

func (d minlenDatum) testString(m *data.Main, t *data.Tag) (success bool, err error) {
	tagValueInt, err := strconv.Atoi(t.Value)
	if err == nil {
		l := len(m.FieldValue.String())
		success = (l == 0 && !m.ContainsTagKey("required")) || l >= tagValueInt
	}

	return
}

func (d minlenDatum) FailureMessage(m *data.Main, t *data.Tag) string {
	switch m.FieldKind {
	case reflect.Slice, reflect.Array, reflect.Map:
		return d.failureMessageCollections(m, t)
	case reflect.String:
		return d.failureMessageString(m, t)
	default:
		panic(ErrUnexpectedType)
	}
}

func (d minlenDatum) failureMessageCollections(m *data.Main, t *data.Tag) string {
	return m.FormattedFieldName + " must contain at least " + t.Value + " values."
}

func (d minlenDatum) failureMessageString(m *data.Main, t *data.Tag) string {
	return m.FormattedFieldName + " must be at least " + t.Value + " characters long."
}
