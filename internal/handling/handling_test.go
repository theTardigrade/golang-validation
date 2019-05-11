package handling

import (
	"reflect"
	"sort"
	"strings"
	"sync"
	"testing"

	"github.com/theTardigrade/tests"
	"github.com/theTardigrade/validation/internal/data"
)

type executeTestExpectedFailureMessagesCallback func(*data.Main, *data.Tag) []string

const (
	expectedTestComparatorSeparator = "\n\t"
)

func executeTest(t *testing.T, model interface{}, callback executeTestExpectedFailureMessagesCallback) {
	typ := reflect.TypeOf(model)
	value := reflect.ValueOf(model)

	var expectedFailureMsgs []string
	var failureMsgs []string
	var mutex sync.Mutex

	for i, l := 0, value.NumField(); i < l; i++ {
		field := typ.Field(i)
		fieldValue := value.FieldByName(field.Name)

		main := data.NewMain(&field, &fieldValue, &failureMsgs, &mutex)

		if err := HandleAllTags(main); err != nil {
			panic(err)
		}

		if tags := main.Tags; callback != nil && tags != nil {
			for _, tag := range tags {
				expectedFailureMsgs = append(expectedFailureMsgs, callback(main, tag)...)
			}
		}
	}

	sort.Strings(failureMsgs)
	sort.Strings(expectedFailureMsgs)

	executeTestRemoveDuplicatesHelper(&failureMsgs)
	executeTestRemoveDuplicatesHelper(&expectedFailureMsgs)

	failureMsgsComparator := strings.Join(failureMsgs, expectedTestComparatorSeparator)
	expectedFailureMsgsComparator := strings.Join(expectedFailureMsgs, expectedTestComparatorSeparator)

	if msg, fail := tests.Message(expectedFailureMsgsComparator, failureMsgsComparator); fail {
		t.Errorf("%s\n%v", msg, model)
	}
}

func executeTestRemoveDuplicatesHelper(a *[]string) {
	var prev string

	for i, l := 0, len(*a); i < l; i++ {
		curr := (*a)[i]

		if prev != "" && curr == prev {
			*a = append((*a)[:i], (*a)[i+1:]...)
			i--
			l--
			continue
		}

		prev = curr
	}
}
