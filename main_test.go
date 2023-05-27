package cleantone

import (
	"fmt"
	"github.com/kadai0308/cleantone/persistenceSvc"
	"math/rand"
	"os"
	"testing"
	"time"
)

var config = DBConfig{
	DataFormat:      persistenceSvc.CSV,
	DataPath:        "/Users/davy/davy/go_playground/kv_db/data",
	RotateThreshold: 100 * MB,
}

var TestDB = NewDB(config)

var INDEX = map[string]string{}
var CsvFile, _ = os.OpenFile("./test.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

var KEY = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
var VALUE = "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"

func BenchmarkSetSingleKey(b *testing.B) {
	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())

	for j := 0; j < b.N; j++ {
		TestDB.Set(KEY, VALUE)
	}

}

func BenchmarkReadKey(b *testing.B) {
	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())

	for j := 0; j < b.N; j++ {
		TestDB.Get("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	}
}

func BenchmarkWriteCSV(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for j := 0; j < b.N; j++ {
		INDEX[KEY] = VALUE
		CsvFile.WriteString(fmt.Sprintf("%s,%s\n", KEY, VALUE))
	}
}

func main() {

	testing.Benchmark(BenchmarkSetSingleKey)

}
