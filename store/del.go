package store

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
