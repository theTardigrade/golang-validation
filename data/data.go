package data

import (
	"bytes"
	"reflect"
	"unicode"
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
	FailureMessages    []string
}

func NewMain(field *reflect.StructField, fieldValue *reflect.Value, failureMessages []string) *Main {
	return &Main{
		Field:              field,
		FieldValue:         fieldValue,
		FormattedFieldName: formatFieldName(field.Name),
		FailureMessages:    failureMessages,
	}
}

func formatFieldName(name string) string {
	var newName bytes.Buffer

	for i, r := range name {
		if unicode.IsUpper(r) {
			if i == 0 {
				newName.WriteRune(r)
			} else {
				newName.WriteRune(' ')
				newName.WriteRune(unicode.ToLower(r))
			}
		} else if unicode.IsLetter(r) {
			newName.WriteRune(r)
		}
	}

	return newName.String()
}

func (m *Main) SetFailure(message string) {
	if m.CurrentTag != nil {
		m.CurrentTag.HasFailed = true
	}

	m.FailureMessages = append(m.FailureMessages, message)
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
