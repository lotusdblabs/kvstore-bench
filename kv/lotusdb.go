package kv

import (
	"github.com/lotusdblabs/lotusdb/v2"
)

func newLotusDB(path string) (Store, error) {
	options := lotusdb.DefaultOptions
	options.DirPath = path
	options.BlockCache = 0
	db, err := lotusdb.Open(options)
	return &lotusdbStore{db: db}, err
}

type lotusdbStore struct {
	db *lotusdb.DB
}

func (l *lotusdbStore) Put(key []byte, value []byte) error {
	return l.db.Put(key, value, nil)
}

func (l *lotusdbStore) Get(key []byte) ([]byte, error) {
	return l.db.Get(key)
}

func (l *lotusdbStore) Delete(key []byte) error {
	return l.db.Delete(key, nil)
}

func (l *lotusdbStore) Close() error {
	return l.db.Close()
}
