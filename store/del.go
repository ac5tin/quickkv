package store

import (
	"errors"

	uf "github.com/ac5tin/usefulgo"
)

// Del - delete a key
func (s *Store) Del(key string) error {
	x, ok := s.Data.Load(s.Key)
	if !ok {
		return errors.New("Unable to load data")
	}

	delete(x.(map[string]interface{}), key)
	s.Data.Store(s.Key, x.(map[string]interface{}))
	if err := s.Save(); err != nil {
		return err
	}
	return nil
}

// ArrRm - remove value from an array
func (s *Store) ArrRm(key string, value interface{}) error {
	x, ok := s.Data.Load(s.Key)
	if !ok {
		return errors.New("Unable to load data")
	}
	if arr, ok := x.(map[string]interface{})[key].([]interface{}); ok {
		arr1 := arr // clone arr
		for i, v := range arr1 {
			if v == value {
				// arr.remove(i)
				uf.NewArrRmiF().Any(&arr, uint32(i))
			}
		}
		x.(map[string]interface{})[key] = arr
		s.Data.Store(s.Key, x.(map[string]interface{}))
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
	x, ok := s.Data.Load(s.Key)
	if !ok {
		return nil, errors.New("Unable to load data")
	}
	var retme interface{}
	if _, ok := x.(map[string]interface{})[key]; !ok {
		x.(map[string]interface{})[key] = make([]interface{}, 0)
	}

	if arr, ok := x.(map[string]interface{})[key].([]interface{}); ok {
		if len(arr) < 1 {
			return nil, errors.New("Slice empty")
		}
		retme = arr[0]
		arr = arr[1:]
		x.(map[string]interface{})[key] = arr
		s.Data.Store(s.Key, x.(map[string]interface{}))
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
	x, ok := s.Data.Load(s.Key)
	if !ok {
		return errors.New("Unable to load data")
	}
	if _, ok := x.(map[string]interface{})[key]; !ok {
		x.(map[string]interface{})[key] = make([]interface{}, 0)
	}

	if arr, ok := x.(map[string]interface{})[key].([]interface{}); ok {
		if len(arr) < shifts {
			return errors.New("Slice bounds out of range")
		}
		arr = arr[:len(arr)-shifts]
		x.(map[string]interface{})[key] = arr
		s.Data.Store(s.Key, x.(map[string]interface{}))
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
	x, ok := s.Data.Load(s.Key)
	if !ok {
		return errors.New("Unable to load data")
	}
	for k := range x.(map[string]interface{}) {
		delete(x.(map[string]interface{}), k)
	}
	s.Data.Store(s.Key, x.(map[string]interface{}))
	if err := s.Save(); err != nil {
		return err
	}
	return nil
}
