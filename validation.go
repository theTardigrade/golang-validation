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

			var i int
			for i = 0; i < l; i++ {
				go func(i int) {
					defer wg.Done()

					field := typ.Field(i)
					fieldValue := value.FieldByName(field.Name)
					d := data.NewMain(&field, &fieldValue, &failureMessages, &failureMessagesMutex)

					errMutex.RLock()
					exitEarly := errIndex < i
					errMutex.RUnlock()

					if exitEarly {
						return
					}

					if err2 := handling.HandleAllTags(d); err2 != nil {
						errMutex.Lock()
						if i < errIndex {
							err, errIndex = err2, i
						}
						errMutex.Unlock()
					}
				}(i)
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
