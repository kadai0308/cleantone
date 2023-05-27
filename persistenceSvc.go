package cleantone

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type PersistenceSvc interface {
	WriteData(key string, value string) error
	BuildIndex() (map[string]string, error)
	Prune(index map[string]string) error
	RotateFile() error
	Flush() error
	Close() error
}

func NewPersistenceSvc(format DataFormatImpl, dataPath string, rotateThreshold int) (PersistenceSvc, error) {
	if format == DataFormat.CSV {
		impl, err := NewCsvImpl(format, dataPath, rotateThreshold)
		if err != nil {
			return nil, err
		}

		return impl, err
	} else if format == DataFormat.JSON {

	}
	errMsg := fmt.Sprintf("Format %s not supported", format)
	return nil, errors.New(errMsg)
}

type BaseImpl struct {
	File            *os.File
	FileID          int
	FileSize        int64
	DataPath        string
	Format          DataFormatImpl
	RotateThreshold int
}

func InitDataFile(dataPath string, extension string) (*os.File, int, error) {
	files, _ := ioutil.ReadDir(dataPath)
	fileID := 0
	var currDbFile *os.File
	var err error
	if len(files) == 0 {
		filePath := fmt.Sprintf("%s/data_%d.%s", dataPath, fileID, extension)
		currDbFile, err = os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	} else {
		latestFileName := files[len(files)-1].Name()
		fileID, _ = strconv.Atoi(strings.Split(latestFileName, "_")[1])
		filePath := filepath.Join(dataPath, latestFileName)
		currDbFile, err = os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	}
	if err != nil {
		return nil, 0, err
	}
	return currDbFile, fileID, nil
}

func (c *BaseImpl) GenerateDataFileName(id int) string {
	name := fmt.Sprintf("data_%d.%s", id, c.Format)
	return name
}
