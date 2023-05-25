package main

import (
	"fmt"
	"github.com/davy/kv_db/src"
	"github.com/davy/kv_db/src/persistenceSvc"
	"github.com/google/uuid"
)

func main() {
	config := src.DBConfig{
		DataFormat:      persistenceSvc.CSV,
		DataPath:        "/Users/davy/davy/go_playground/kv_db/data",
		RotateThreshold: 1 * src.MB,
	}

	DB := src.NewDB(config)
	defer DB.PersistenceSvc.Flush()

	for i := 0; i < 1000000; i++ {
		//uuid1 := uuid.New()
		uuid2 := uuid.New()
		DB.Set("a", uuid2.String())
		//fmt.Println(i, uuid2)
	}

	fmt.Println(DB.PersistenceSvc.Prune(DB.Index))

	return

}
