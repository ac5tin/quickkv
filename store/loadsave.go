package store

import (
	"bytes"
	"encoding/gob"
)

// Save - save data to file
func (s *Store) Save() error {
	x, err := s.getMap()
	if err != nil {
		return err
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

// GetBinary - retrieve store data in binary
func (s *Store) GetBinary() ([]byte, error) {
	x, err := s.getMap()
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	if err := encoder.Encode(x); err != nil {
		return nil, err
	}
	b := buf.Bytes()
	return b, nil
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
	for k, v := range d {
		s.Data.Store(k, v)
	}
	return nil
}

// LoadBinary - load binary into store
func (s *Store) LoadBinary(b []byte) error {
	var d map[string]interface{}
	reader := bytes.NewReader(b)
	decoder := gob.NewDecoder(reader)
	if err := decoder.Decode(&d); err != nil {
		return err
	}
	for k, v := range d {
		s.Data.Store(k, v)
	}
	return nil
}
