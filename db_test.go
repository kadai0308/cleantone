package cleantone

import (
	"fmt"
	"github.com/kadai0308/cleantone/persistenceSvc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
	"testing"
)

type DBTestSuite struct {
	suite.Suite
	DBSvc    *DB
	DBConfig *DBConfig
}

//func (suite *DBTestSuite) SetupSuite() {
//dataPath := "/tmp/cleantone_test_data"
//err := os.Mkdir(dataPath, 0755)
//if err != nil {
//	log.Panic(err)
//}
//
//dbConfig := DBConfig{
//	RotateThreshold: 1 * KB,
//	DataPath:        dataPath,
//	DataFormat:      persistenceSvc.CSV,
//}
//suite.DBConfig = &dbConfig
//suite.DBSvc = NewDB(dbConfig)
//}

//func (suite *DBTestSuite) TearDownSuite() {
//	err := os.RemoveAll(suite.DBConfig.DataPath)
//	if err != nil {
//		log.Panic(err)
//		return
//	}
//}

func (suite *DBTestSuite) BeforeTest(_, _ string) {
	dataPath := "/tmp/cleantone_test_data"
	err := os.Mkdir(dataPath, 0755)
	if err != nil {
		log.Panic(err)
	}

	dbConfig := DBConfig{
		RotateThreshold: 1 * KB,
		DataPath:        dataPath,
		DataFormat:      persistenceSvc.CSV,
	}
	suite.DBConfig = &dbConfig
	suite.DBSvc = NewDB(dbConfig)
}

func (suite *DBTestSuite) AfterTest(_, _ string) {
	err := os.RemoveAll(suite.DBConfig.DataPath)
	if err != nil {
		log.Panic(err)
		return
	}
}

func (suite *DBTestSuite) TestDBSet() {
	key := "a"
	value := "b"
	suite.DBSvc.Set(key, value)

	assert.Equal(suite.T(), suite.DBSvc.Index[key], value)
}

func (suite *DBTestSuite) TestDBRead() {
	key := "a"
	value := "b"
	suite.DBSvc.Set(key, value)

	result, err := suite.DBSvc.Get(key)
	if err != nil {
		assert.Equal(suite.T(), err.Error(), fmt.Sprintf("Key %s not exist.", key))
	}
	assert.Equal(suite.T(), result, value)
}

func (suite *DBTestSuite) TestDBPersistence() {
	totalDataSize := 0
	key := "a"
	for i := 0; i < 1*KB; i++ {
		value := strconv.FormatInt(int64(i), 10)
		suite.DBSvc.Set(key, value)

		totalDataSize = totalDataSize + len(fmt.Sprintf("%s,%s\n", key, value))
	}
	files, _ := ioutil.ReadDir(suite.DBConfig.DataPath)
	fileAmount := math.Ceil(float64(totalDataSize) / float64(suite.DBConfig.RotateThreshold))

	assert.Equal(suite.T(), int(fileAmount), len(files))
}

func (suite *DBTestSuite) TestDBInitBuildIndex() {
	key := "a"
	value := "666"
	for i := 0; i < KB; i++ {
		suite.DBSvc.Set(key, value)
	}
	newValue := "76786838"
	suite.DBSvc.Set(key, newValue)
	suite.DBSvc.Close()
	suite.DBSvc = NewDB(*suite.DBConfig)

	result, err := suite.DBSvc.Get(key)
	if err != nil {
		log.Panic(err)
	}
	assert.Equal(suite.T(), result, newValue)
}

func TestDBTestSuite(t *testing.T) {
	suite.Run(t, new(DBTestSuite))
}
