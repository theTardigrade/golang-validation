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

	if len(tag.Key) < 6 {
		return errors.New(tag.Key)
	}

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
			errCollection := make([]error, l)
			var errCollectionMutex sync.RWMutex

			wg.Add(l)

			var i int
			for i = l - 1; i >= 0; i-- {
				go func(i int) {
					defer wg.Done()

					tag := tags[i]

					var exitEarly bool

					errCollectionMutex.RLock()
					for j := 0; j < i; j++ {
						if errCollection[j] != nil {
							exitEarly = true
						}
					}
					errCollectionMutex.RUnlock()

					if exitEarly {
						return
					}

					if err2 := HandleTag(m, tag); err2 != nil {
						errCollectionMutex.Lock()
						errCollection[i] = err2
						errCollectionMutex.Unlock()
					}
				}(i)
			}

			wg.Wait()

			for i = 0; i < l; i++ {
				if err2 := errCollection[i]; err2 != nil {
					err = err2
					break
				}
			}
		}
	}

	return
}
