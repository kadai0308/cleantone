package src

import (
	"errors"
	"fmt"
	"github.com/davy/kv_db/src/persistenceSvc"
	"log"
)

type DB struct {
	Index          map[string]string
	PersistenceSvc persistenceSvc.PersistenceSvc
}

func NewDB(config DBConfig) *DB {
	persistenceSvc, err := persistenceSvc.NewPersistenceSvc(config.DataFormat, config.DataPath, config.RotateThreshold)
	if err != nil {
		log.Fatal(err)
	}
	index, err := persistenceSvc.BuildIndex()
	if err != nil {
		log.Fatal(err)
	}
	db := &DB{
		Index:          index,
		PersistenceSvc: persistenceSvc,
	}
	return db
}

func (db *DB) Persist(key string, value string) error {
	return db.PersistenceSvc.WriteData(key, value)
}

func (db *DB) Set(key string, value string) {
	db.Index[key] = value
	if err := db.Persist(key, value); err != nil {
		log.Fatal(err)
	}
}

func (db *DB) Get(key string) (string, error) {
	value, ok := db.Index[key]

	if !ok {
		errMsg := fmt.Sprintf("Key %s not exist.", key)
		return "", errors.New(errMsg)
	}

	return value, nil
}
