package store

import "sync"

// Store - main store
type Store struct {
	Path string
	Data map[string]interface{}
	Mux  *sync.RWMutex
}
