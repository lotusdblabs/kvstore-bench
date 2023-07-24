package kv

import (
	"github.com/cockroachdb/pebble"
)

type pebbleStore struct {
	db *pebble.DB
}

func newPebble(path string) (Store, error) {
	db, err := pebble.Open(path, &pebble.Options{})
	if err != nil {
		return nil, err
	}

	return &pebbleStore{db: db}, nil
}

func (p *pebbleStore) Put(key []byte, value []byte) error {
	if err := p.db.Set(key, value, pebble.NoSync); err != nil {
		return err
	}
	return nil
}

func (p *pebbleStore) Get(key []byte) ([]byte, error) {
	value, closer, err := p.db.Get(key)
	if err != nil {
		return nil, err
	}

	defer closer.Close()
	return value, nil
}

func (p *pebbleStore) Delete(key []byte) error {
	return p.db.Delete(key, pebble.NoSync)
}

func (p *pebbleStore) Close() error {
	return p.db.Close()
}
