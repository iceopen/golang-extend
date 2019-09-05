package kvbench

import (
	"sync"

	"github.com/dgraph-io/badger"
)

type badgerStore struct {
	mu sync.RWMutex
	db *badger.DB
}

func badgerKey(key []byte) []byte {
	r := make([]byte, len(key)+1)
	r[0] = 'k'
	copy(r[1:], key)
	return r
}

func NewBadgerStore(path string, fsync bool) (Store, error) {
	if path == ":memory:" {
		return nil, errMemoryNotAllowed
	}

	opts := badger.DefaultOptions(path)

	opts.Dir = path
	opts.ValueDir = path
	opts.SyncWrites = fsync
	db, err := badger.Open(opts)
	if err != nil {
		return nil, err
	}

	return &badgerStore{
		db: db,
	}, nil
}

func (s *badgerStore) Close() error {
	s.db.Close()
	return nil
}

func (s *badgerStore) PSet(keys, vals [][]byte) error {
	txn := s.db.NewTransaction(true)
	for i, k := range keys {
		v := vals[i]
		if err := txn.Set(k, v); err == badger.ErrTxnTooBig {
			_ = txn.Commit()
			txn = s.db.NewTransaction(true)
			_ = txn.Set(k, v)
		}
	}
	return txn.Commit()
}

func (s *badgerStore) PGet(keys [][]byte) ([][]byte, []bool, error) {
	var vals = make([][]byte, len(keys))
	var oks = make([]bool, len(keys))

	err := s.db.View(func(txn *badger.Txn) error {
		for i, k := range keys {
			item, err := txn.Get(k)
			if err == nil {
				v, err := item.ValueCopy(nil)
				if err == nil {
					vals[i] = v
					oks[i] = true
				}
			}
		}
		return nil
	})

	return vals, oks, err
}

func (s *badgerStore) Set(key, value []byte) error {
	return s.db.Update(func(txn *badger.Txn) error {
		return txn.Set(key, value)
	})
}

func (s *badgerStore) Get(key []byte) ([]byte, bool, error) {
	var v []byte

	err := s.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err == nil {
			err = item.Value(func(value []byte) error {
				v = value
				return nil
			})
			// v, err = item.ValueCopy(nil)
		}
		return err
	})

	return v, v != nil, err
}

func (s *badgerStore) Del(key []byte) (bool, error) {
	err := s.db.Update(func(txn *badger.Txn) error {
		return txn.Delete(key)
	})
	return err == nil, err
}

func (s *badgerStore) Keys(pattern []byte, limit int, withvals bool) ([][]byte, [][]byte, error) {
	var keys [][]byte
	var vals [][]byte

	err := s.db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		prefix := pattern
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			k := item.Key()
			keys = append(keys, k)
			if withvals {
				v, err := item.ValueCopy(nil)
				if err != nil {
					continue
				}
				vals = append(vals, v)
			}
		}
		return nil
	})

	return keys, vals, err
}

func (s *badgerStore) FlushDB() error {
	return s.db.DropAll()
}
