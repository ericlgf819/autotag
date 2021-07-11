package csvwriter

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SimpleCSVWriterTestSuite struct {
	suite.Suite
	testFilePath string
	target       *SimpleCSVWriter
	testContent  [][]string
}

func (suite *SimpleCSVWriterTestSuite) SetupTest() {
	suite.testFilePath = "./csvtestingfile.csv"
	suite.testContent = [][]string{
		{"col1", "col2", "col3"},
		{"how", "are", "you?"},
		{"i", "am", "fine"},
	}
	suite.target = new(SimpleCSVWriter)
}

func (suite *SimpleCSVWriterTestSuite) TearDownSuite() {
	os.Remove(suite.testFilePath)
}

func (suite *SimpleCSVWriterTestSuite) TestWriteFileAndFileExist() {
	err := suite.target.WriteFile(suite.testFilePath, suite.testContent)
	fileInfo, statErr := os.Stat(suite.testFilePath)

	assert.Nil(suite.T(), err)
	assert.Nil(suite.T(), statErr)
	assert.Contains(suite.T(), suite.testFilePath, fileInfo.Name())
	assert.Greater(suite.T(), fileInfo.Size(), int64(0))
}

func TestSimpleCSVWriterTestSuite(t *testing.T) {
	suite.Run(t, new(SimpleCSVWriterTestSuite))
}
