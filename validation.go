package validation

import (
	"reflect"

	"github.com/theTardigrade/validation/internal/data"
	"github.com/theTardigrade/validation/internal/tests"
)

const (
	tagName = "validation"
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
		for i, l := 0, t.NumField(); i < l; i++ {
			field := t.Field(i)
			fieldValue := value.FieldByName(field.Name)
			d := data.NewMain(&field, &fieldValue, &failureMessages)

			tests.HandleAllTags(d)
		}
	}

	isValidated = len(failureMessages) == 0
	return
}
