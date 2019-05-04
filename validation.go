package validation

import (
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

func Validate(opts Options) (isValidated bool, failureMessages []string, err error) {
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
			var wg sync.WaitGroup
			var mutex sync.RWMutex
			var pool sync.Pool

			wg.Add(l)

			for i := 0; i < l; i++ {
				go func(i int) {
					field := typ.Field(i)
					fieldValue := value.FieldByName(field.Name)
					d := data.NewMain(&field, &fieldValue, &failureMessages, &mutex)

					if err2 := handling.HandleAllTags(d); err2 != nil {
						pool.Put(err2)
					}

					wg.Done()
				}(i)
			}

			wg.Wait()

			if err2, ok := pool.Get().(error); ok {
				err = err2
			}
		}
	}

	if l := len(failureMessages); l == 0 {
		isValidated = true
	} else if opts.SortFailureMessages {
		sort.Strings(failureMessages)
	}

	return
}
