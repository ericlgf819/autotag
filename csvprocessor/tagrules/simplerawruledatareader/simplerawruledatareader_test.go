package simplerawruledatareader

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SimpleRawRuleDataReaderTestSuite struct {
	suite.Suite
	target          *SimpleRawRuleDataReader
	mockFileContent [][]string
}

type MockCSVFileReader struct {
	testSuite *SimpleRawRuleDataReaderTestSuite
}

func (reader *MockCSVFileReader) ReadFile(string) ([][]string, error) {
	return reader.testSuite.mockFileContent, nil
}

func (suite *SimpleRawRuleDataReaderTestSuite) SetupTest() {
	suite.mockFileContent = [][]string{
		{"col1", "col2", "col3", "colOther", "col4", "col5"},
		{"hello", "world", "hi", "nono", "there", "yes"},
		{"today", "weather", "is", "no", "good", "fine"},
	}
	suite.target = new(SimpleRawRuleDataReader)
	suite.target.Init(&MockCSVFileReader{testSuite: suite})
}

func (suite *SimpleRawRuleDataReaderTestSuite) TestLoadFile() {
	err := suite.target.LoadFile("mockFilePath")

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), len(suite.mockFileContent), len(suite.target.fileContent))
}

func TestSimpleRawRuleDataReaderTestSuite(t *testing.T) {
	suite.Run(t, new(SimpleRawRuleDataReaderTestSuite))
}
