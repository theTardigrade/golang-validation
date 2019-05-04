package validation

import (
	"reflect"
	"sync"

	"github.com/theTardigrade/validation/internal/data"
	"github.com/theTardigrade/validation/internal/handling"
)

func Validate(model interface{}) (isValidated bool, failureMessages []string, err error) {
	t := reflect.TypeOf(model)
	kind := t.Kind()
	value := reflect.ValueOf(model)

	for kind == reflect.Ptr || kind == reflect.Interface {
		value = value.Elem()
		kind, t = value.Kind(), value.Type()
	}

	if kind == reflect.Struct {
		if l := t.NumField(); l > 0 {
			var wg sync.WaitGroup
			var mutex sync.RWMutex
			var pool sync.Pool

			wg.Add(l)

			for i := 0; i < l; i++ {
				go func(i int) {
					field := t.Field(i)
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

	isValidated = len(failureMessages) == 0
	return
}
