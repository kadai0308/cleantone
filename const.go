package cleantone

const (
	b = 1 << (10 * iota)
	kb
	mb
)

var FileSize = struct {
	B  int
	KB int
	MB int
}{
	b,
	kb,
	mb,
}
