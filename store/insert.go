package store

import (
	"encoding/json"
	"sync"

	uf "github.com/ac5tin/usefulgo"
)

// Set - sets a value in store
func (s *Store) Set(key string, value interface{}) error {
	s.Mux.Lock()
	defer s.Mux.Unlock()
	s.Data[key] = value

	b, err := json.Marshal(s.Data)
	if err != nil {
		return err
	}
	if err := uf.NewFS().Write(b, s.Path); err != nil {
		return err
	}
	return nil
}

// MSet - multiple set
func (s *Store) MSet(key string, values []interface{}) error {
	var wg sync.WaitGroup
	var er error = nil
	for _, value := range values {
		go func(value interface{}, wg *sync.WaitGroup) {
			defer wg.Done()
			if err := s.Set(key, value); err != nil {
				er = err
			}
		}(value, &wg)
	}
	wg.Wait()
	return er
}

// QMSet - quick multiple set
func (s *Store) QMSet(key string, values []interface{}) {
	for _, value := range values {
		go s.Set(key, value)
	}
}
