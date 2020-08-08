package store

import (
	"encoding/json"
	"sync"
)

// Set - sets a value in store
func (s Store) Set(key string, value interface{}) error {
	s.Mux.Lock()
	defer s.Mux.Unlock()
	s.Data[key] = value

	b, err := json.Marshal(s.Data)
	if err != nil {
		return err
	}

	if err := s.write(&b); err != nil {
		return err
	}
	return nil
}

// MSet - multiple set
func (s Store) MSet(key string, values []interface{}) error {
	var wg sync.WaitGroup
	var er error = nil
	for _, value := range values {
		wg.Add(1)
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
func (s Store) QMSet(key string, values []interface{}) {
	for _, value := range values {
		go s.Set(key, value)
	}
}

// MapSet - multiple set using map
func (s Store) MapSet(input map[string]interface{}) error {
	var wg sync.WaitGroup
	var er error = nil
	for key, value := range input {
		wg.Add(1)
		go func(key string, value interface{}, wg *sync.WaitGroup) {
			defer wg.Done()
			if err := s.Set(key, value); err != nil {
				er = err
			}
		}(key, value, &wg)
	}
	wg.Wait()
	return er
}

// QMapSet - quick map set
func (s Store) QMapSet(input map[string]interface{}) {
	for key, value := range input {
		go s.Set(key, value)
	}
}
