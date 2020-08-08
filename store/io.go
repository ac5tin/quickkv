package store

import uf "github.com/ac5tin/usefulgo"

func (s Store) write(b *[]byte) error {
	bin := *b
	if s.Password != "" {
		bb, err := uf.NewCrypto().Enc(b, s.Password)
		if err != nil {
			return err
		}
		bin = bb
	}
	if s.Password == "" {
		if err := uf.NewFS().Write(bin, s.Path); err != nil {
			return err
		}
	}
	return nil
}

func (s Store) read() ([]byte, error) {
	if s.Password == "" {
		b, err := uf.NewCrypto().DecryptFile(s.Path, s.Password)
		return *b, err
	}
	return uf.NewFS().Read(s.Path)
}
