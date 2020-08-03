package store

import (
	"encoding/json"
	"errors"
	"sync"

	uf "github.com/ac5tin/usefulgo"
)

// Get - retrieve data from store
func (s *Store) Get(key string) (interface{}, error) {
	if v, ok := s.Data[key]; ok {
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

// Load - load data from file
func (s *Store) Load() error {
	b, err := uf.NewFS().Read(s.Path)
	if err != nil {
		return err
	}

	var d map[string]interface{}
	if err := json.Unmarshal(b, &d); err != nil {
		return err
	}
	s.Data = d
	return nil
}
