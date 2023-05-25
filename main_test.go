package main

import (
	"fmt"
	"github.com/davy/kv_db/src"
	"github.com/davy/kv_db/src/persistenceSvc"
	"math/rand"
	"os"
	"testing"
	"time"
)

var config = src.DBConfig{
	DataFormat:      persistenceSvc.CSV,
	DataPath:        "/Users/davy/davy/go_playground/kv_db/data",
	RotateThreshold: 100 * src.MB,
}

var DB = src.NewDB(config)

var INDEX = map[string]string{}
var CsvFile, _ = os.OpenFile("./test.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

var KEY = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
var VALUE = "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"

func BenchmarkSetSingleKey(b *testing.B) {
	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())

	for j := 0; j < b.N; j++ {
		DB.Set(KEY, VALUE)
	}
	// 创建性能图表的数据点
	//var xvalue []float64
	//var yvalue []float64
	//
	//for i := 1; i <= 10; i++ {
	//	start := time.Now()
	//
	//	// 执行基准测试循环
	//	for j := 0; j < b.N; j++ {
	//		DB.Set("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb")
	//	}
	//
	//	elapsed := time.Since(start)
	//	// 将结果添加到数据点
	//	xvalue = append(xvalue, float64(i))
	//	yvalue = append(yvalue, float64(elapsed.Milliseconds()))
	//}

	//fmt.Println(xvalue)
	//fmt.Println(yvalue)

}

//func BenchmarkReadKey(b *testing.B) {
//	// 设置随机数种子
//	rand.Seed(time.Now().UnixNano())
//
//	for j := 0; j < b.N; j++ {
//		DB.Get("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
//	}
//}

func BenchmarkWriteCSV(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for j := 0; j < b.N; j++ {
		INDEX[KEY] = VALUE
		CsvFile.WriteString(fmt.Sprintf("%s,%s\n", KEY, VALUE))
	}
}

func main() {

	testing.Benchmark(BenchmarkWriteCSV)
	//testing.Benchmark(BenchmarkReadKey)
	//defer DB.PersistenceSvc.Flush()
	//
	//for i := 0; i < 1000000; i++ {
	//	//uuid1 := uuid.New()
	//	uuid2 := uuid.New()
	//	DB.Set("a", uuid2.String())
	//	//fmt.Println(i, uuid2)
	//}
	//
	//fmt.Println(DB.PersistenceSvc.Prune(DB.Index))

	return

}
