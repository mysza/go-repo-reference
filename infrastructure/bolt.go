package infrastructure

import (
	"io/ioutil"
	"log"

	"github.com/boltdb/bolt"
)

type BoltDatabase struct {
	blt *bolt.DB
}

func (db *BoltDatabase) Open() {
	tmpDir, err := ioutil.TempDir("", "boltexample")
	if err != nil {
		log.Fatal(err)
	}
	tmpFile, err := ioutil.TempFile(tmpDir, "boltexample")
	if err != nil {
		log.Fatal(err)
	}
	fName := tmpFile.Name()
	blt, err := bolt.Open(fName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	db.blt = blt
}

func (db *BoltDatabase) Close() {
	if db.blt != nil {
		db.blt.Close()
	}
}

func (db *BoltDatabase) GetEntity(entityKind string, key []byte) ([]byte, error) {
	var retValue []byte
	err := db.blt.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(entityKind))
		retValue = b.Get(key)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return retValue, err
}

func (db *BoltDatabase) PutEntity(entityKind string, key, value []byte) error {
	return db.blt.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(entityKind))
		return b.Put(key, value)
	})
}
