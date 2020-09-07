package store

import "errors"

func (s *Store) getMap() (map[string]interface{}, error) {
	y, ok := s.Data.Load(s.Key)
	x := make(map[string]interface{})
	for k, v := range y.(map[string]interface{}) {
		x[k] = v
	}
	if !ok {
		return nil, errors.New("Unable to load data")
	}
	return x, nil
}
