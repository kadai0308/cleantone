package src

import "github.com/davy/kv_db/src/persistenceSvc"

type DBConfig struct {
	RotateThreshold int
	DataPath        string
	DataFormat      persistenceSvc.DataFormat
}
