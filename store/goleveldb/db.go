package goleveldb

import (
	"bytes"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

// DB is a LevelDB database.
type DB struct {
	db *leveldb.DB
}

// NewDB creates a new DB instance.
func NewDB(path string) (*DB, error) {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return nil, err
	}

	return &DB{
		db: db,
	}, nil
}

// Close closes the database.
func (d *DB) Close() error {
	return d.db.Close()
}

// Put puts a key-value pair into the database.
func (d *DB) Put(key, value []byte) error {
	return d.db.Put(key, value, &opt.WriteOptions{Sync: true})
}

// Get gets the value of a key from the database.
func (d *DB) Get(key []byte) ([]byte, error) {
	return d.db.Get(key, nil)
}

// Next returns the next key-value pair in the database.
func (d *DB) Next(key []byte) ([]byte, []byte, error) {
	iter := d.db.NewIterator(nil, nil)
	defer iter.Release()

	iter.Seek(key)
	if !iter.Valid() {
		return nil, nil, nil
	}

	return iter.Key(), iter.Value(), nil
}

// Prev returns the previous key-value pair in the database.
func (d *DB) Prev(key []byte) ([]byte, []byte, error) {
	iter := d.db.NewIterator(nil, nil)
	defer iter.Release()

	iter.Seek(key)
	iter.Prev()
	if !iter.Valid() {
		return nil, nil, nil
	}

	return iter.Key(), iter.Value(), nil
}

// Iterate iterates over the key-value pairs in the database.
func (d *DB) Iterate(prefix []byte, fn func(key, value []byte) error) error {
	iter := d.db.NewIterator(nil, nil)
	defer iter.Release()

	for iter.Seek(prefix); iter.Valid(); iter.Next() {
		key := iter.Key()
		value := iter.Value()

		if !bytes.HasPrefix(key, prefix) {
			break
		}

		if err := fn(key, value); err != nil {
			return err
		}
	}

	return nil
}

// Prune prunes the database.
func (d *DB) Prune(prefix []byte) error {
	iter := d.db.NewIterator(nil, nil)
	defer iter.Release()

	iter.Seek(prefix)
	iter.Prev() // Skip the last key-value pair with the prefix.
	for iter.Prev(); iter.Valid(); iter.Prev() {
		key := iter.Key()

		if err := d.db.Delete(key, &opt.WriteOptions{Sync: true}); err != nil {
			return err
		}
	}

	return nil
}
