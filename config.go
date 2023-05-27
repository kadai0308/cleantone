package cleantone

import "github.com/kadai0308/cleantone/PersistenceSvc"

type DBConfig struct {
	RotateThreshold int
	DataPath        string
	DataFormat      PersistenceSvc.DataFormatImpl
}
