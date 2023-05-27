package gobUtils

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type DataRow struct {
	Key   string
	Value string
}

var lengthDigit int = 10

type GobPersistence struct {
	File     *os.File
	FileID   int
	DataPath string
}

func NewGobPersistence(dataPath string) (*GobPersistence, error) {
	files, _ := ioutil.ReadDir(dataPath)
	fileID := 0
	var currDbFile *os.File
	var err error
	if len(files) == 0 {
		filePath := fmt.Sprintf("%s/data_%d.gob", dataPath, fileID)
		currDbFile, err = os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	} else {
		latestFileName := files[len(files)-1].Name()
		fileID, _ = strconv.Atoi(strings.Split(latestFileName, "_")[1])
		filePath := filepath.Join(dataPath, latestFileName)
		currDbFile, err = os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	}

	if err != nil {
		return nil, err
	}
	return &GobPersistence{
		File:     currDbFile,
		FileID:   fileID,
		DataPath: dataPath,
	}, nil
}

func (g *GobPersistence) WriteData(key string, value string) error {
	data := DataRow{
		Key: key, Value: value,
	}
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(data)
	if err != nil {
		return err
	}

	hexStr := strconv.FormatInt(int64(buffer.Len()), 16)
	hexStrFixLen := fmt.Sprintf("%010s", hexStr)

	g.File.WriteString(hexStrFixLen)
	g.File.Write(buffer.Bytes())
	return nil
}

func (g *GobPersistence) RotateFile() error {
	return nil
}

func (g *GobPersistence) BuildIndex() (map[string]string, error) {
	files, _ := ioutil.ReadDir(g.DataPath)
	for _, fileInfo := range files {
		fmt.Println(fileInfo.Name())
		filePath := filepath.Join(g.DataPath, fileInfo.Name())

		file, err := os.OpenFile(filePath, os.O_RDONLY, 0)
		if err != nil {
			return nil, err
		}
		// read data length
		dataLength := make([]byte, lengthDigit)
		_, err = io.ReadFull(file, dataLength)
		if err != nil {
			return nil, err
		}
		dataSize, err := strconv.ParseInt(string(dataLength), 16, 64)
		if err != nil {
			fmt.Println("Error converting hex string to int:", err)
			return nil, err
		}

		fmt.Println("data size:", dataSize)

		byteData := make([]byte, dataSize)
		_, err = io.ReadFull(file, byteData)
		if err != nil {
			return nil, err
		}
		reader := bytes.NewReader(byteData)
		decoder := gob.NewDecoder(reader)
		var data DataRow
		err = decoder.Decode(&data)
		if err != nil {
			fmt.Println("Error decoding Gob data:", err)
			return nil, err
		}

		fmt.Println("Decoded data:", data)

	}
	return nil, nil
}

func (g *GobPersistence) MergeFiles() error {
	return nil
}

func (g *GobPersistence) Close() {
	g.File.Close()
}
