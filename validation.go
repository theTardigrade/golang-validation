package validation

import (
	"reflect"
	"strings"

	"github.com/theTardigrade/validation/data"
	"github.com/theTardigrade/validation/tests"
)

const (
	tagName           = "validation"
	tagSeparator      = ","
	tagValueSeparator = "="
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
			tag := field.Tag.Get(tagName)
			splitTags := strings.Split(tag, tagSeparator)
			d := data.NewMain(&field, &fieldValue, &failureMessages)
			d.Tags = make([]*data.Tag, 0, len(splitTags))

			var tagKey string
			var tagValue string

			for _, tag = range splitTags {
				if tagValueSeparatorIndex := strings.Index(tag, tagValueSeparator); tagValueSeparatorIndex != -1 {
					tagValue = tag[tagValueSeparatorIndex+1:]
					tagKey = tag[:tagValueSeparatorIndex]
				} else {
					tagKey = tag
				}

				d.Tags = append(d.Tags, &data.Tag{
					Key:   tagKey,
					Value: tagValue,
				})
			}

			for _, tagPtr := range d.Tags {
				d.CurrentTag = tagPtr

				var handler tests.Handler

				switch tagPtr.Key {
				case "required":
					handler = tests.Required
				case "email":
					handler = tests.Email
				case "minlen":
					handler = tests.MinLen
				case "maxlen":
					handler = tests.MaxLen
				case "min":
					handler = tests.Min
				case "max":
					handler = tests.Max
				}

				if handler != nil {
					if err = handler(d); err != nil {
						return
					}
				}
			}
		}
	}

	isValidated = len(failureMessages) == 0
	return
}
