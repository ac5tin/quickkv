package store

import "sync"

// Store - main store
type Store struct {
	Path     string
	Data     sync.Map
	Password string
}
