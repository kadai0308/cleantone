package cleantone

import "github.com/kadai0308/cleantone/persistenceSvc"

type DBConfig struct {
	RotateThreshold int
	DataPath        string
	DataFormat      persistenceSvc.DataFormat
}
