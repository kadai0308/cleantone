package cleantone

type DBConfig struct {
	RotateThreshold int
	DataPath        string
	DataFormat      DataFormatImpl
}
