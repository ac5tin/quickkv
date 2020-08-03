package store

import (
	"encoding/json"
	"log"
	"os"
	"sync"

	uf "github.com/ac5tin/usefulgo"
)

// Init - initialise store
func Init(path string) *Store {
	d := make(map[string]interface{})
	if !fileExist(path) {
		// file doesnt exist, create new file
		if err := createFile(path); err != nil {
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

func fileExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}

func createFile(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}
