package store

import "sync"

// Store - main store
type Store struct {
	Path     string
	Data     *map[string]interface{}
	Password string
	Mux      *sync.RWMutex
}
