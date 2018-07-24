package main

import (
	"log"

	"github.com/dgraph-io/badger"
	"fmt"
)

func main() {
	// Open the Badger database located in the /tmp/badger directory.
	// It will be created if it doesn't exist.
	opts := badger.DefaultOptions
	opts.Dir = "./data"
	opts.ValueDir = "./data"
	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	//Read-write transactions
	err = db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte("answer"), []byte("42"))
		return err
	})
	if err != nil {
		log.Fatal(err)
	}
	// Read-only transactions
	err = db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("answer"))
		if err != nil {
			return err
		}
		val, err := item.Value()
		if err != nil {
			return err
		}
		fmt.Printf("The answer is: %s\n", val)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	// delete test date
	err = db.Update(func(txn *badger.Txn) error {
		txn.Delete([]byte("answer"))
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
