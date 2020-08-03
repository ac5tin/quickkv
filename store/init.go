package store

import (
	"encoding/json"
	"log"
	"sync"

	uf "github.com/ac5tin/usefulgo"
)

// Init - initialise store
func Init(path string) *Store {
	d := make(map[string]interface{})
	if !uf.NewFS().FileExist(path) {
		// file doesnt exist, create new file
		if err := uf.NewFS().CreateFile(path); err != nil {
			log.Panic(err.Error())
		}
	} else {
		b, err := uf.NewFS().Read(path)
		if err != nil {
			log.Panic(err.Error())
		}
		if err := json.Unmarshal(b, &d); err != nil {
			log.Panic(err.Error())
		}
	}
	s := Store{
		Path: path,
		Data: d,
		Mux:  &sync.RWMutex{},
	}
	return &s

}
