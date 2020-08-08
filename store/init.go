package store

import (
	"encoding/json"
	"log"
	"sync"

	uf "github.com/ac5tin/usefulgo"
)

// Init - initialise store
func Init(path, password string) Store {
	d := make(map[string]interface{})
	f := uf.NewFS()
	if !f.FileExist(path) {
		// file doesnt exist, create new file
		if err := f.CreateFile(path); err != nil {
			log.Panic(err.Error())
		}
		b, err := json.Marshal(d)
		if err != nil {
			log.Panic(err.Error())
		}
		if password != "" {
			encb, err := uf.NewCrypto().Enc(&b, password)
			if err != nil {
				log.Panic(err.Error())
			}
			b = encb
		}
		f.Write(b, path)

	} else {
		b, err := f.Read(path)
		if err != nil {
			log.Panic(err.Error())
		}

		if password != "" {
			decb, err := uf.NewCrypto().Dec(&b, password)
			if err != nil {
				log.Panic(err.Error())
			}
			b = *decb
		}

		if err := json.Unmarshal(b, &d); err != nil {
			log.Panic(err.Error())
		}
	}

	s := Store{
		Path:     path,
		Data:     d,
		Password: password,
		Mux:      &sync.RWMutex{},
	}

	return s

}
