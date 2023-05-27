package cleantone

type DataFormat string

var (
	CSV  DataFormat = "csv"
	JSON DataFormat = "json"
)

const (
	B = 1 << (10 * iota)
	KB
	MB
)
