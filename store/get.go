package store

import (
	"errors"
	"sync"
)

// Get - retrieve data from store
func (s *Store) Get(key string) (interface{}, error) {
	v, ok := s.Data.Load(key)
	if !ok {
		return nil, errors.New("Failed to find key")
	}
	return v, nil
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
