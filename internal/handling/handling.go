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
	var wg sync.WaitGroup
	var pool sync.Pool

	if tags := m.Tags; tags != nil {
		if l := len(tags); l > 0 {
			wg.Add(l)

			for _, tag := range m.Tags {
				go func(tag *data.Tag) {
					if err2 := HandleTag(m, tag); err2 != nil {
						pool.Put(err2)
					}

					wg.Done()
				}(tag)
			}
		}
	}

	wg.Wait()

	if err2, ok := pool.Get().(error); ok {
		err = err2
	}

	return
}
