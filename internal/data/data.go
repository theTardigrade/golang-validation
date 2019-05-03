package data

import (
	"reflect"
	"strings"
	"unicode"
)

const (
	validationTagName           = "validation"
	validationTagSeparator      = ","
	validationTagValueSeparator = "="
	formattedFieldNameTagName   = "name"
)

type Tag struct {
	Key, Value string
	HasFailed  bool
}

type TagCollection []*Tag

type Main struct {
	Field              *reflect.StructField
	FieldValue         *reflect.Value
	FormattedFieldName string
	Tags               TagCollection
	CurrentTag         *Tag
	FailureMessages    *[]string
}

func NewMain(field *reflect.StructField, fieldValue *reflect.Value, failureMessages *[]string) *Main {
	main := Main{
		Field:           field,
		FieldValue:      fieldValue,
		FailureMessages: failureMessages,
	}

	main.loadTags()
	main.loadFormattedFieldName()

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

func (m *Main) loadTags() {
	tag := m.Field.Tag.Get(validationTagName)
	splitTags := strings.Split(tag, validationTagSeparator)
	m.Tags = make(TagCollection, 0, len(splitTags))

	var tagKey string
	var tagValue string

	for _, tag = range splitTags {
		if tagValueSeparatorIndex := strings.Index(tag, validationTagValueSeparator); tagValueSeparatorIndex != -1 {
			tagValue = tag[tagValueSeparatorIndex+1:]
			tagKey = tag[:tagValueSeparatorIndex]
		} else {
			tagKey = tag
		}

		m.Tags = append(m.Tags, &Tag{
			Key:   tagKey,
			Value: tagValue,
		})
	}
}

func (m *Main) loadFormattedFieldName() {
	if formattedFieldName, found := m.Field.Tag.Lookup(formattedFieldNameTagName); found {
		m.FormattedFieldName = formattedFieldName
	} else {
		m.FormattedFieldName = formatFieldName(m.Field.Name)
	}
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
