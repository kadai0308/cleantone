package cleantone

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type CsvImpl struct {
	BaseImpl
	Writer *bufio.Writer
}

func NewCsvImpl(format DataFormatImpl, dataPath string, rotateThreshold int) (*CsvImpl, error) {
	file, fileID, err := InitDataFile(dataPath, string(format))
	if err != nil {
		return nil, err
	}

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	write := bufio.NewWriterSize(file, 2*rotateThreshold)
	return &CsvImpl{
		BaseImpl{
			File:            file,
			FileID:          fileID,
			FileSize:        fileInfo.Size(),
			DataPath:        dataPath,
			Format:          format,
			RotateThreshold: rotateThreshold,
		},
		write,
	}, nil
}

func (c *CsvImpl) GenerateCsvRow(key string, value string) string {
	StringBuilder := strings.Builder{}
	StringBuilder.WriteString(key)
	StringBuilder.WriteString(",")
	StringBuilder.WriteString(value)
	StringBuilder.WriteString("\n")
	return StringBuilder.String()
}

func (c *CsvImpl) WriteData(key string, value string) error {
	data := c.GenerateCsvRow(key, value)
	c.Writer.WriteString(data)
	c.FileSize = c.FileSize + int64(len(data))
	err := c.RotateFile()
	if err != nil {
		return err
	}
	return nil
}

func (c *CsvImpl) BuildIndex() (map[string]string, error) {
	index := map[string]string{}
	files, _ := ioutil.ReadDir(c.DataPath)
	for _, fileInfo := range files {
		filePath := filepath.Join(c.DataPath, fileInfo.Name())

		file, err := os.OpenFile(filePath, os.O_RDONLY, 0)
		if err != nil {
			return nil, err
		}

		reader := csv.NewReader(file)
		reader.Comma = ','
		records, err := reader.ReadAll()
		if err != nil {
			return nil, errors.New(fmt.Sprintf("reading %s err: %s", fileInfo.Name(), err.Error()))
		}

		for _, record := range records {
			index[record[0]] = record[1]
		}
	}
	return index, nil
}

func (c *CsvImpl) RotateFile() error {
	if c.FileSize >= int64(c.RotateThreshold) {
		err := c.Writer.Flush()
		if err != nil {
			return err
		}
		c.FileID = c.FileID + 1
		newFile := c.GenerateDataFileName(c.FileID)
		newFilePath := filepath.Join(c.DataPath, newFile)
		file, err := os.OpenFile(newFilePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			return err
		}
		c.File.Close()
		c.File = file
		c.Writer = bufio.NewWriterSize(file, 2*c.RotateThreshold)
		c.FileSize = 0
	}
	return nil
}

func (c *CsvImpl) Flush() error {
	return c.Writer.Flush()
}

func (c *CsvImpl) GetFileID(fileName string) int {
	fileIDStr := strings.Split(strings.Split(fileName, "_")[1], ".")[0]
	fileID, _ := strconv.Atoi(fileIDStr)
	return fileID
}

func (c *CsvImpl) Prune(index map[string]string) error {
	files, _ := ioutil.ReadDir(c.DataPath)
	for _, fileInfo := range files {
		fileID := c.GetFileID(fileInfo.Name())
		if fileID >= c.FileID {
			continue
		}
		filePath := filepath.Join(c.DataPath, fileInfo.Name())
		os.Remove(filePath)
	}

	newDataPath := filepath.Join(c.DataPath, "new_data")
	os.Mkdir(newDataPath, 0755)
	csvPersistSvc, err := NewCsvImpl(DataFormat.CSV, newDataPath, c.RotateThreshold)
	if err != nil {
		return err
	}
	for key, value := range index {
		csvPersistSvc.WriteData(key, value)
	}
	csvPersistSvc.Flush()

	files, _ = ioutil.ReadDir(newDataPath)
	for _, fileInfo := range files {
		os.Rename(filepath.Join(newDataPath, fileInfo.Name()), filepath.Join(c.DataPath, fileInfo.Name()))
	}
	os.RemoveAll(newDataPath)

	return nil
}

func (c *CsvImpl) Close() error {
	return c.File.Close()
}
