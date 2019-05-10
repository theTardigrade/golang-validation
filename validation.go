package validation

import (
	"math"
	"reflect"
	"sort"
	"sync"

	"github.com/theTardigrade/validation/internal/data"
	"github.com/theTardigrade/validation/internal/handling"
)

type Options struct {
	Model               interface{}
	SortFailureMessages bool
}

func Validate(model interface{}) (bool, []string, error) {
	return ValidateWithOptions(Options{Model: model})
}

func ValidateWithOptions(opts Options) (isValidated bool, failureMessages []string, err error) {
	model := opts.Model
	typ := reflect.TypeOf(model)
	kind := typ.Kind()
	value := reflect.ValueOf(model)

	for kind == reflect.Ptr || kind == reflect.Interface {
		value = value.Elem()
		kind, typ = value.Kind(), value.Type()
	}

	if kind == reflect.Struct {
		if l := typ.NumField(); l > 0 {
			var failureMessagesMutex sync.Mutex
			var wg sync.WaitGroup
			var errMutex sync.RWMutex
			errIndex := math.MaxInt32

			wg.Add(l)

			for i := 0; i < l; i++ {
				go validateField(i, typ, value, &wg, &err, &errIndex, &errMutex, &failureMessages, &failureMessagesMutex)
			}

			wg.Wait()
		}
	}

	if l := len(failureMessages); l == 0 {
		isValidated = true
	} else if opts.SortFailureMessages {
		sort.Strings(failureMessages)
	}

	return
}

func validateField(
	i int,
	typ reflect.Type,
	value reflect.Value,
	wgPtr *sync.WaitGroup,
	errPtr *error,
	errIndexPtr *int,
	errMutexPtr *sync.RWMutex,
	failureMessagesPtr *[]string,
	failureMessagesMutexPtr *sync.Mutex,
) {
	defer wgPtr.Done()

	field := typ.Field(i)
	fieldValue := value.FieldByName(field.Name)
	main := data.NewMain(&field, &fieldValue, failureMessagesPtr, failureMessagesMutexPtr)

	errMutexPtr.RLock()
	exitEarly := *errIndexPtr < i
	errMutexPtr.RUnlock()

	if exitEarly {
		return
	}

	if err := handling.HandleAllTags(main); err != nil {
		errMutexPtr.Lock()
		if i < *errIndexPtr {
			*errPtr, *errIndexPtr = err, i
		}
		errMutexPtr.Unlock()
	}
}
