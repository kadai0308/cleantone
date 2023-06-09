package cleantone

import (
	"errors"
	"fmt"
	"log"
)

type DB struct {
	Index          map[string]string
	PersistenceSvc PersistenceSvc
}

func NewDB(config DBConfig) *DB {
	persistenceSvc, err := newPersistenceSvc(config.DataFormat, config.DataPath, config.RotateThreshold)
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

func (db *DB) persist(key string, value string) error {
	return db.PersistenceSvc.WriteData(key, value)
}

func (db *DB) Set(key string, value string) {
	db.Index[key] = value
	if err := db.persist(key, value); err != nil {
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

func (db *DB) Close() error {
	if err := db.PersistenceSvc.Flush(); err != nil {
		log.Fatal(err)
	}

	if err := db.PersistenceSvc.Close(); err != nil {
		log.Fatal(err)
	}
	return nil
}
