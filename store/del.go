package store

import (
	"errors"

	uf "github.com/ac5tin/usefulgo"
)

// Del - delete a key
func (s Store) Del(key string) error {
	s.Mux.Lock()
	defer s.Mux.Unlock()
	delete(s.Data, key)
	if err := s.Save(); err != nil {
		return err
	}
	return nil
}

// ArrRm - remove value from an array
func (s Store) ArrRm(key string, value interface{}) error {
	s.Mux.Lock()
	defer s.Mux.Unlock()
	if arr, ok := s.Data[key].([]interface{}); ok {
		arr1 := arr // clone arr
		for i, v := range arr1 {
			if v == value {
				// arr.remove(i)
				uf.NewArrRmiF().Any(&arr, uint32(i))
			}
		}
		s.Data[key] = arr

		if err := s.Save(); err != nil {
			return err
		}
	} else {
		return errors.New("Key not of array type")
	}

	return nil
}

// Pop - removes last element from array and returns it
func (s Store) Pop(key string) (interface{}, error) {
	s.Mux.Lock()
	defer s.Mux.Unlock()
	var retme interface{}
	if _, ok := s.Data[key]; !ok {
		s.Data[key] = make([]interface{}, 0)
	}

	if arr, ok := s.Data[key].([]interface{}); ok {
		retme = arr[0]
		arr = arr[1:]
		s.Data[key] = arr

		if err := s.Save(); err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("Key not of array type")
	}

	return retme, nil
}

// Reset - resets the store
func (s Store) Reset() error {
	s.Mux.Lock()
	defer s.Mux.Unlock()
	for k := range s.Data {
		delete(s.Data, k)
	}
	if err := s.Save(); err != nil {
		return err
	}
	return nil
}
