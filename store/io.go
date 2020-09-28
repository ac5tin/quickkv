package store

import (
	"log"

	uf "github.com/ac5tin/usefulgo"
)

func (s *Store) write(b *[]byte) error {
	bin := *b
	if s.Password != "" {
		bb, err := uf.NewCrypto().Enc(b, s.Password)
		if err != nil {
			return err
		}
		bin = bb
	}
	log.Println("writing")
	if err := uf.NewFS().Write(bin, s.Path); err != nil {
		log.Println(err.Error())
		return err
	}
	log.Println("finished writing, start replication")
	go func() {
		if err := s.Replicate(); err != nil {
			log.Println(err.Error())
		}
	}()
	return nil
}

func (s *Store) read() ([]byte, error) {
	if s.Password == "" {

		return uf.NewFS().Read(s.Path)
	}
	b, err := uf.NewCrypto().DecryptFile(s.Path, s.Password)
	if b == nil {
		return nil, err
	}
	return *b, err

}
