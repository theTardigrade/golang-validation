package handling

import (
	"reflect"
	"regexp"
	"sync"

	"github.com/theTardigrade/golang-validation/data"
)

func init() {
	addHandler("regexp", regexpDatum{})
}

var (
	regexpCache      = make(map[string]*regexp.Regexp)
	regexpCacheMutex sync.Mutex
)

type regexpDatum struct{}

func (d regexpDatum) Test(m *data.Main, t *data.Tag) (success bool, err error) {
	switch m.FieldKind {
	case reflect.String:
		success, err = d.testString(m, t)
	default:
		err = ErrUnexpectedType
	}

	return
}

func (d regexpDatum) testString(m *data.Main, t *data.Tag) (success bool, err error) {
	s := m.FieldValue.String()

	r := func() *regexp.Regexp {
		v := t.Value

		defer regexpCacheMutex.Unlock()
		regexpCacheMutex.Lock()

		if r, ok := regexpCache[v]; ok {
			return r
		} else {
			r := regexp.MustCompile(v)
			regexpCache[v] = r
			return r
		}
	}()

	success = r.MatchString(s)
	return
}

func (d regexpDatum) FailureMessage(m *data.Main, t *data.Tag) string {
	return m.FormattedFieldName + ` must match a standard format.`
}
