package store

import (
	"log"
	"sync"

	uf "github.com/ac5tin/usefulgo"
)

// Init - initialise store
func Init(path, password string) *Store {
	f := uf.NewFS()

	s := Store{
		Path:     path,
		Data:     sync.Map{},
		Password: password,
		Key:      "DATA",
	}
	d := make(map[string]interface{})
	s.Data.Store(s.Key, d)

	if f.FileExist(path) {
		// file exist, load data
		if err := s.Load(); err != nil {
			log.Panic(err.Error())
		}
	} else {
		// file doesn't exist
		if err := s.Save(); err != nil {
			log.Panic(err.Error())
		}
	}

	STORE = &s
	return STORE

}
