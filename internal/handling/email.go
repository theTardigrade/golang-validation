package handling

import (
	"reflect"
	"regexp"

	"github.com/theTardigrade/validation/internal/data"
)

func init() {
	addHandler("email", email)
}

var (
	emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

func email(m *data.Main, t *data.Tag) error {
	switch m.Field.Type.Kind() {
	case reflect.String:
		if !emailRegexp.MatchString(m.FieldValue.String()) {
			m.SetFailure(t, m.FormattedFieldName+" not recognised as valid.")
		}
	default:
		return ErrUnexpectedType
	}

	return nil
}
