package store

import (
	"errors"

	uf "github.com/ac5tin/usefulgo"
)

// Del - delete a key
func (s *Store) Del(key string) error {
	s.Data.Delete(key)
	if err := s.Save(); err != nil {
		return err
	}
	return nil
}

// ArrRm - remove value from an array
func (s *Store) ArrRm(key string, value interface{}) error {
	arr, err := s.getArr(key)
	if err != nil {
		return err
	}

	for i, v := range arr {
		if v == value {
			// arr.remove(i)
			uf.NewArrRmiF().Any(&arr, uint32(i))
		}
	}
	s.Data.Store(key, arr)
	if err := s.Save(); err != nil {
		return err
	}

	return nil
}

// Pop - removes last element from array and returns it
func (s *Store) Pop(key string) (interface{}, error) {
	arr, err := s.getArr(key)
	if err != nil {
		return nil, err
	}
	if len(arr) == 0 {
		return nil, errors.New("Empty slice")
	}
	retme := arr[0]
	arr = arr[1:]
	s.Data.Store(key, arr)
	if err := s.Save(); err != nil {
		return nil, err
	}

	return retme, nil
}

// Shift - shift array
func (s *Store) Shift(key string, shifts int) error {
	arr, err := s.getArr(key)
	if err != nil {
		return err
	}
	if len(arr) < shifts {
		return errors.New("Slice bounds out of range")
	}
	arr = arr[:len(arr)-shifts]
	s.Data.Store(key, arr)
	if err := s.Save(); err != nil {
		return err
	}
	return nil
}

// Reset - resets the store
func (s *Store) Reset() error {
	s.Data.Range(func(k, v interface{}) bool {
		s.Data.Delete(k)
		return true
	})
	if err := s.Save(); err != nil {
		return err
	}
	return nil
}
