package handling

import (
	"strconv"
	"time"

	"github.com/theTardigrade/age"
	"github.com/theTardigrade/validation/internal/data"
)

func init() {
	addHandler("minage", minageDatum{})
}

type minageDatum struct{}

func (d minageDatum) Test(m *data.Main, t *data.Tag) (success bool, err error) {
	i := m.FieldValue.Interface()

	if date, ok := i.(time.Time); ok {
		success, err = d.testTime(m, t, date)
	} else {
		err = ErrUnexpectedType
	}

	return
}

func (d minageDatum) testTime(m *data.Main, t *data.Tag, date time.Time) (success bool, err error) {
	if date.IsZero() && !m.ContainsTagKey("required") {
		success = true
		return
	}

	tagValueInt, err := strconv.ParseInt(t.Value, 10, 64)
	if err == nil && int64(age.Calculate(date)) >= tagValueInt {
		success = true
	}

	return
}

func (d minageDatum) FailureMessage(m *data.Main, t *data.Tag) string {
	return m.FormattedFieldName + " cannot be less than " + t.Value + " years of age."
}
