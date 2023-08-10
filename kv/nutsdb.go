package kv

import (
	"github.com/nutsdb/nutsdb"
)

type nutsdbStore struct {
	db *nutsdb.DB
}

func newNutsDB(path string) (Store, error) {
	options := nutsdb.DefaultOptions
	options.Dir = path
	options.EntryIdxMode = nutsdb.HintKeyAndRAMIdxMode
	options.SyncEnable = false

	db, err := nutsdb.Open(options)
	return &nutsdbStore{db: db}, err
}

func (n nutsdbStore) Put(key []byte, value []byte) error {
	return n.db.Update(func(tx *nutsdb.Tx) error {
		return tx.Put("bucket", key, value, nutsdb.Persistent)
	})
}

func (n nutsdbStore) Get(key []byte) ([]byte, error) {
	var (
		value []byte
	)
	err := n.db.View(func(tx *nutsdb.Tx) error {
		e, _ := tx.Get("bucket", key)
		value = e.Value
		return nil
	})
	return value, err
}

func (n nutsdbStore) Delete(key []byte) error {
	return n.db.Update(func(tx *nutsdb.Tx) error {
		return tx.Delete("bucket", key)
	})
}

func (n nutsdbStore) Close() error {
	return n.db.Close()
}
