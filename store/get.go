package store

import (
	"errors"
	"sync"
)

// Get - retrieve data from store
func (s *Store) Get(key string) (interface{}, error) {
	x, err := s.getMap()
	if err != nil {
		return nil, err
	}
	if v, ok := x[key]; ok {
		return v, nil
	}
	return nil, errors.New("Failed to find key")
}

// MGet - multiple get
func (s *Store) MGet(keys []string) []interface{} {
	retme := make([]interface{}, 0)
	var wg sync.WaitGroup
	mux := &sync.RWMutex{}
	for _, key := range keys {
		wg.Add(1)
		go func(key string, wg *sync.WaitGroup) {
			defer wg.Done()
			v, _ := s.Get(key)
			mux.Lock()
			defer mux.Unlock()
			retme = append(retme, v)
		}(key, &wg)
	}
	wg.Wait()
	return retme
}

// GetAll - retrieve whole store
func (s *Store) GetAll() map[string]interface{} {
	x, err := s.getMap()
	if err != nil {
		return nil
	}
	return x
}
