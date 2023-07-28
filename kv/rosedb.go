package kv

import "github.com/rosedblabs/rosedb/v2"

func newRoseDB(path string) (Store, error) {
	options := rosedb.DefaultOptions
	options.DirPath = path
	db, err := rosedb.Open(options)
	return &rosedbStore{db: db}, err
}

type rosedbStore struct {
	db *rosedb.DB
}

func (r *rosedbStore) Put(key []byte, value []byte) error {
	return r.db.Put(key, value)
}

func (r *rosedbStore) Get(key []byte) ([]byte, error) {
	return r.db.Get(key)
}

func (r *rosedbStore) Delete(key []byte) error {
	return r.db.Delete(key)
}

func (r *rosedbStore) Close() error {
	return r.db.Close()
}
