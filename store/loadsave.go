package store

import (
	"bytes"
	"encoding/gob"
)

// Save - save data to file
func (s *Store) Save() error {
	x, ok := s.Data.Load(s.Key)
	if !ok {
		return nil
	}
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	if err := encoder.Encode(x); err != nil {
		return err
	}
	b := buf.Bytes()
	s.write(&b)
	return nil
}

// Load - load data from file
func (s *Store) Load() error {
	b, err := s.read()
	if err != nil {
		return err
	}

	var d map[string]interface{}
	reader := bytes.NewReader(b)
	decoder := gob.NewDecoder(reader)
	if err := decoder.Decode(&d); err != nil {
		return err
	}
	s.Data.Store(s.Key, d)
	return nil
}
