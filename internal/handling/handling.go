package handling

import (
	"errors"
	"sync"

	"github.com/theTardigrade/validation/internal/data"
)

var (
	ErrUnexpectedType = errors.New("unexpected type")
)

type handler (func(m *data.Main) error)
type handlerCollection []handler
type storeMap map[string]handlerCollection

var (
	store      = make(storeMap)
	storeMutex sync.RWMutex
)

func addHandler(key string, h handler) {
	defer storeMutex.Unlock()
	storeMutex.Lock()

	for k, c := range store {
		if k == key {
			c = append(c, h)
			return
		}
	}

	c := make(handlerCollection, 1)
	c[0] = h
	store[key] = c
}

func HandleTag(m *data.Main, tag *data.Tag) (err error) {
	defer storeMutex.RUnlock()
	storeMutex.RLock()

	m.CurrentTag = tag

	if handlers, found := store[tag.Key]; found {
		for _, h := range handlers {
			if err = h(m); err != nil {
				break
			}
		}
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
