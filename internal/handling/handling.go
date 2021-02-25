package handling

import (
	"errors"
	"sync"

	"github.com/theTardigrade/fan"
	"github.com/theTardigrade/golang-validation/internal/data"
)

var (
	ErrUnexpectedType = errors.New("unexpected type")
)

type handler interface {
	Test(*data.Main, *data.Tag) (bool, error)
	FailureMessage(*data.Main, *data.Tag) string
}

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

	if handlers, found := store[tag.Key]; found {
		for _, h := range handlers {
			var ok bool

			if ok, err = h.Test(m, tag); err != nil {
				break
			} else if !ok {
				m.SetFailure(tag, h.FailureMessage(m, tag))
			}
		}
	}

	return
}

func HandleAllTags(m *data.Main) (err error) {
	if tags := m.Tags; tags != nil {
		if l := len(tags); l > 0 {
			fan.HandleRepeated(func(i int) error {
				return HandleTag(m, tags[i])
			}, l)
		}
	}

	return
}
