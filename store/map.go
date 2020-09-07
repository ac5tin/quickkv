package store

func (s *Store) getMap() (map[string]interface{}, error) {
	x := make(map[string]interface{})
	s.Data.Range(func(key, value interface{}) bool {
		x[key.(string)] = value
		return true
	})
	return x, nil
}

func (s *Store) getArr(key string) ([]interface{}, error) {
	y, ok := s.Data.Load(key)
	if !ok {
		// key doesnt exist, initialise empty array
		y = make([]interface{}, 0)
	}
	x := make([]interface{}, len(y.([]interface{})))
	copy(x, y.([]interface{}))

	return x, nil
}
