package handling

import (
	"strconv"
	"time"

	"github.com/theTardigrade/validation/internal/data"
)

func init() {
	addHandler("maxYearsAgo", maxYearsAgoDatum{})
}

type maxYearsAgoDatum struct{}

func (d maxYearsAgoDatum) Test(m *data.Main, t *data.Tag) (success bool, err error) {
	i := m.FieldValue.Interface()

	if tm, ok := i.(time.Time); ok {
		success, err = d.testTime(m, t, tm)
	} else {
		err = ErrUnexpectedType
	}

	return
}

func (d maxYearsAgoDatum) testTime(m *data.Main, t *data.Tag, tm time.Time) (success bool, err error) {
	currentYear := time.Now().UTC().Year()
	givenYear := tm.Year()

	tagValueInt, err := strconv.Atoi(t.Value)
	if err != nil {
		return
	}

	if givenYear >= currentYear-tagValueInt {
		success = true
	}

	return
}

func (d maxYearsAgoDatum) FailureMessage(m *data.Main, t *data.Tag) string {
	return m.FormattedFieldName + " cannot be more than " + t.Value + " years ago."
}
