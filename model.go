package cleantone

type DataFormatImpl string

var DataFormat = struct {
	CSV  DataFormatImpl
	JSON DataFormatImpl
}{
	"csv",
	"json",
}
