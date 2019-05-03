package handling

import (
	"errors"

	"github.com/theTardigrade/validation/internal/data"
)

var (
	ErrUnexpectedType = errors.New("unexpected type")
)

type Handler (func(m *data.Main) error)

func HandleTag(m *data.Main, tag *data.Tag) (err error) {
	m.CurrentTag = tag

	var handler Handler

	switch tag.Key {
	case "required":
		handler = Required
	case "email":
		handler = Email
	case "minlen":
		handler = MinLen
	case "maxlen":
		handler = MaxLen
	case "min":
		handler = Min
	case "max":
		handler = Max
	}

	if handler != nil {
		err = handler(m)
	}

	return
}

func HandleAllTags(m *data.Main) (err error) {
	for _, tag := range m.Tags {
		if err = HandleTag(m, tag); err != nil {
			break
		}
	}

	return
}
