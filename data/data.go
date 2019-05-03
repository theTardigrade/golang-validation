package data

import (
	"reflect"
	"strings"
	"unicode"
)

const (
	formattedFieldNameTagName = "name"
)

type Tag struct {
	Key, Value string
	HasFailed  bool
}

type Main struct {
	Field              *reflect.StructField
	FieldValue         *reflect.Value
	FormattedFieldName string
	Tags               []*Tag
	CurrentTag         *Tag
	FailureMessages    *[]string
}

func NewMain(field *reflect.StructField, fieldValue *reflect.Value, failureMessages *[]string) *Main {
	main := Main{
		Field:           field,
		FieldValue:      fieldValue,
		FailureMessages: failureMessages,
	}

	if formattedFieldName, found := field.Tag.Lookup(formattedFieldNameTagName); found {
		main.FormattedFieldName = formattedFieldName
	} else {
		main.FormattedFieldName = formatFieldName(field.Name)
	}

	return &main
}

func formatFieldName(name string) string {
	var builder strings.Builder
	var i int

	for _, r := range name {
		if unicode.IsLetter(r) {
			if unicode.IsUpper(r) {
				if i == 0 {
					builder.WriteRune(r)
				} else {
					builder.WriteRune(' ')
					builder.WriteRune(unicode.ToLower(r))
				}
			} else if unicode.IsLower(r) {
				if i == 0 {
					builder.WriteRune(unicode.ToUpper(r))
				} else {
					builder.WriteRune(r)
				}
			}

			i++
		}
	}

	return builder.String()
}

func (m *Main) SetFailure(message string) {
	if m.CurrentTag != nil {
		m.CurrentTag.HasFailed = true
	}

	*m.FailureMessages = append(*m.FailureMessages, message)
}

func (m *Main) ContainsTagKey(key string) bool {
	for _, t := range m.Tags {
		if t.Key == key {
			return true
		}
	}

	return false
}

func (m Main) TagFromKey(key string) *Tag {
	for _, t := range m.Tags {
		if t.Key == key {
			return t
		}
	}

	return nil
}
