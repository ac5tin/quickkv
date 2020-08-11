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
