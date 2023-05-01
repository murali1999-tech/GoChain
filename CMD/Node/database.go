
package main

import (
	"github.com/dgraph-io/badger/v3"
)

type Database struct {
	db *badger.DB
}

func NewDatabase(dataDir string) (*Database, error) {
	opts := badger.DefaultOptions(dataDir)
	db, err := badger.Open(opts)
	if err != nil {
		return nil, err
	}
	return &Database{db: db}, nil
}

func (d *Database) Close() error {
	return d.db.Close()
}

func (d *Database) Put(key, value []byte) error {
	err := d.db.Update(func(txn *badger.Txn) error {
		return txn.Set(key, value)
	})
	return err
}

func (d *Database) Get(key []byte) ([]byte, error) {
	var value []byte
	err := d.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			return err
		}
		value, err = item.ValueCopy(nil)
		return err
	})
	return value, err
}
