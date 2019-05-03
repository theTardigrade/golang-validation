package data

import (
	"reflect"
	"strings"
	"sync"
	"unicode"
)

const (
	validationStructTagName           = "validation"
	validationStructTagSeparator      = ","
	validationStructTagValueSeparator = "="
	formattedFieldNameStructTagName   = "name"
	formattedFieldNameTagKey          = formattedFieldNameStructTagName
)

type Tag struct {
	Key, Value string
	HasFailed  bool
}

type TagCollection []*Tag

type Main struct {
	Field                *reflect.StructField
	FieldValue           *reflect.Value
	FormattedFieldName   string
	Tags                 TagCollection
	CurrentTag           *Tag
	FailureMessages      *[]string
	failureMessagesMutex *sync.Mutex
}

func NewMain(
	field *reflect.StructField,
	fieldValue *reflect.Value,
	failureMessages *[]string,
	mutex *sync.Mutex,
) *Main {
	main := Main{
		Field:                field,
		FieldValue:           fieldValue,
		FailureMessages:      failureMessages,
		failureMessagesMutex: mutex,
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
	tag := m.Field.Tag.Get(validationStructTagName)
	splitTags := strings.Split(tag, validationStructTagSeparator)
	m.Tags = make(TagCollection, 0, len(splitTags))

	var tagKey string
	var tagValue string

	for _, tag = range splitTags {
		if separatorIndex := strings.Index(tag, validationStructTagValueSeparator); separatorIndex != -1 {
			tagValue = tag[separatorIndex+1:]
			tagKey = tag[:separatorIndex]
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
	var n string

	if tag := m.TagFromKey(formattedFieldNameTagKey); tag != nil {
		n = tag.Value
	} else if formattedFieldName, found := m.Field.Tag.Lookup(formattedFieldNameStructTagName); found {
		n = formattedFieldName
	} else {
		n = formatFieldName(m.Field.Name)
	}

	m.FormattedFieldName = n
}

func (m *Main) SetFailure(message string) {
	if m.CurrentTag != nil {
		m.CurrentTag.HasFailed = true
	}

	m.failureMessagesMutex.Lock()
	*m.FailureMessages = append(*m.FailureMessages, message)
	m.failureMessagesMutex.Unlock()
}

func (m *Main) ContainsTagKey(key string) bool {
	return m.TagFromKey(key) != nil
}

func (m Main) TagFromKey(key string) *Tag {
	for _, t := range m.Tags {
		if t.Key == key {
			return t
		}
	}

	return nil
}
