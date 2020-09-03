package store

import (
	"bytes"
	"encoding/gob"
)

// Save - save data to file
func (s *Store) Save() error {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	if err := encoder.Encode(s.Data); err != nil {
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
	// s.Data = d <- this line doesnt work
	// need to run loop to reassign keyvalue
	for key, value := range d {
		s.Data[key] = value
	}
	return nil
}
