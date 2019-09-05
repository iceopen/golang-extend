package cache

import (
	"github.com/dgraph-io/badger"
)

type inBadgerCache struct {
	// db 数据库
	db *badger.DB
	Stat
}

func (b *inBadgerCache) Set(k string, v []byte) error {
	return b.db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(k), v)
	})
}

func (b *inBadgerCache) Get(k string) ([]byte, error) {
	var v []byte

	err := b.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(k))
		if err == nil {
			err = item.Value(func(value []byte) error {
				v = value
				return nil
			})
		}
		return err
	})
	return v, err
}

func (b *inBadgerCache) Del(k string) error {
	err := b.db.Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte(k))
	})
	return err
}

func (b *inBadgerCache) GetStat() Stat {
	b.KeySize,b.ValueSize = b.db.Size()
	return b.Stat
}

func newInBadgerCache() *inBadgerCache {
	dirPath := "dir.db"
	opts := badger.DefaultOptions(dirPath)
	opts.Dir = dirPath
	opts.ValueDir = dirPath
	opts.SyncWrites = true
	db, _ := badger.Open(opts)
	return &inBadgerCache{db, Stat{}}
}
