package store

import (
	"errors"

	uf "github.com/ac5tin/usefulgo"
)

// Del - delete a key
func (s *Store) Del(key string) error {
	x, err := s.getMap()
	if err != nil {
		return err
	}

	delete(x, key)
	s.Data.Store(s.Key, x)
	if err := s.Save(); err != nil {
		return err
	}
	return nil
}

// ArrRm - remove value from an array
func (s *Store) ArrRm(key string, value interface{}) error {
	x, err := s.getMap()
	if err != nil {
		return err
	}
	if arr, ok := x[key].([]interface{}); ok {
		arr1 := arr // clone arr
		for i, v := range arr1 {
			if v == value {
				// arr.remove(i)
				uf.NewArrRmiF().Any(&arr, uint32(i))
			}
		}
		x[key] = arr
		s.Data.Store(s.Key, x)
		if err := s.Save(); err != nil {
			return err
		}
	} else {
		return errors.New("Key not of array type")
	}

	return nil
}

// Pop - removes last element from array and returns it
func (s *Store) Pop(key string) (interface{}, error) {
	x, err := s.getMap()
	if err != nil {
		return nil, err
	}
	var retme interface{}
	if _, ok := x[key]; !ok {
		x[key] = make([]interface{}, 0)
	}

	if arr, ok := x[key].([]interface{}); ok {
		if len(arr) < 1 {
			return nil, errors.New("Slice empty")
		}
		retme = arr[0]
		arr = arr[1:]
		x[key] = arr
		s.Data.Store(s.Key, x)
		if err := s.Save(); err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("Key not of array type")
	}

	return retme, nil
}

// Shift - shift array
func (s *Store) Shift(key string, shifts int) error {
	x, err := s.getMap()
	if err != nil {
		return err
	}
	if _, ok := x[key]; !ok {
		x[key] = make([]interface{}, 0)
	}

	if arr, ok := x[key].([]interface{}); ok {
		if len(arr) < shifts {
			return errors.New("Slice bounds out of range")
		}
		arr = arr[:len(arr)-shifts]
		x[key] = arr
		s.Data.Store(s.Key, x)
		if err := s.Save(); err != nil {
			return err
		}
	} else {
		return errors.New("Key not of array type")
	}

	return nil
}

// Reset - resets the store
func (s *Store) Reset() error {
	x, err := s.getMap()
	if err != nil {
		return err
	}
	for k := range x {
		delete(x, k)
	}
	s.Data.Store(s.Key, x)
	if err := s.Save(); err != nil {
		return err
	}
	return nil
}
