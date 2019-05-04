package handling

import (
	"errors"
	"sync"

	"github.com/theTardigrade/validation/internal/data"
)

var (
	ErrUnexpectedType = errors.New("unexpected type")
)

type handler (func(*data.Main, *data.Tag) error)
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
			if err = h(m, tag); err != nil {
				break
			}
		}
	}

	return
}

func HandleAllTags(m *data.Main) (err error) {
	if tags := m.Tags; tags != nil {
		if l := len(tags); l > 0 {
			var wg sync.WaitGroup
			var pool sync.Pool

			wg.Add(l)

			for i := 0; i < l; i++ {
				go func(tag *data.Tag) {
					if err = HandleTag(m, tag); err != nil {
						pool.Put(err)
					}

					wg.Done()
				}(m.Tags[i])
			}

			wg.Wait()

			if err2, ok := pool.Get().(error); ok {
				err = err2
			}
		}
	}

	return
}
