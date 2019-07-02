package handling

import (
	"strconv"
	"time"

	"github.com/theTardigrade/validation/internal/data"
	utilTime "github.com/theTardigrade/validation/internal/util/time"
)

func init() {
	addHandler("minage", minageDatum{})
}

type minageDatum struct{}

func (d minageDatum) Test(m *data.Main, t *data.Tag) (success bool, err error) {
	i := m.FieldValue.Interface()

	if tm, ok := i.(time.Time); ok {
		success, err = d.testTime(m, t, tm)
	} else {
		err = ErrUnexpectedType
	}

	return
}

func (d minageDatum) testTime(m *data.Main, t *data.Tag, tm time.Time) (success bool, err error) {
	tagValueInt, err := strconv.ParseInt(t.Value, 10, 64)
	if err != nil {
		return
	}

	if age := utilTime.Age(tm); int64(age) >= tagValueInt {
		success = true
	}

	return
}

func (d minageDatum) FailureMessage(m *data.Main, t *data.Tag) string {
	return m.FormattedFieldName + " cannot be less than " + t.Value + " years ago."
}
