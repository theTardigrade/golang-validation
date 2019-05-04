package handling

import (
	"reflect"
	"sync"
	"testing"

	"github.com/theTardigrade/tests"
	"github.com/theTardigrade/validation/internal/data"
)

func executeTest(t *testing.T, model interface{}, expectedFailureMsgsLen int) {
	typ := reflect.TypeOf(model)
	value := reflect.ValueOf(model)

	var failureMsgs []string
	var mutex sync.RWMutex

	for i, l := 0, value.NumField(); i < l; i++ {
		field := typ.Field(i)
		fieldValue := value.FieldByName(field.Name)

		m := data.NewMain(&field, &fieldValue, &failureMsgs, &mutex)

		if err := HandleAllTags(m); err != nil {
			panic(err)
		}
	}

	if msg, fail := tests.Message(expectedFailureMsgsLen, len(failureMsgs)); fail {
		t.Errorf("%s\n%v", msg, model)
	}
}
