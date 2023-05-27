package src

import "github.com/kadai0308/cleantone/src/persistenceSvc"

type DBConfig struct {
	RotateThreshold int
	DataPath        string
	DataFormat      persistenceSvc.DataFormat
}
