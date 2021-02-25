package validation

import (
	"reflect"
	"sort"
	"sync"

	fan "github.com/theTardigrade/golang-fan"
	"github.com/theTardigrade/golang-validation/data"
	"github.com/theTardigrade/golang-validation/handling"
)

type Options struct {
	Model               interface{}
	SortFailureMessages bool
}

func Validate(model interface{}) (bool, []string, error) {
	return ValidateWithOptions(Options{
		Model:               model,
		SortFailureMessages: true,
	})
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

	if value == reflect.Zero(typ) {
		return
	}

	if kind == reflect.Struct {
		if l := typ.NumField(); l > 0 {
			var failureMessagesMutex sync.Mutex

			fan.HandleRepeated(func(i int) error {
				return validateField(i, typ, value, &failureMessages, &failureMessagesMutex)
			}, l)
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
	failureMessagesPtr *[]string,
	failureMessagesMutexPtr *sync.Mutex,
) error {
	field := typ.Field(i)
	fieldValue := value.FieldByName(field.Name)
	main := data.NewMain(&field, &fieldValue, failureMessagesPtr, failureMessagesMutexPtr)

	return handling.HandleAllTags(main)
}
