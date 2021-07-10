package csvreader

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SimpleCSVReaderTestSuite struct {
	suite.Suite
	testFilePath string
	target       *SimpleCSVReader
}

func (suite *SimpleCSVReaderTestSuite) SetupTest() {
	suite.testFilePath = "./csvtestingfile.csv"
	suite.target = new(SimpleCSVReader)
}

func (suite *SimpleCSVReaderTestSuite) TestReadFileWithoutError() {
	results, err := suite.target.ReadFile(suite.testFilePath)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), results)
}

func (suite *SimpleCSVReaderTestSuite) TestReadFileWithCorrectLines() {
	results, err := suite.target.ReadFile(suite.testFilePath)

	assert.Nil(suite.T(), err)

	// the number 3 should be equal to the lines number in file ./csvtestingfile.csv
	assert.Equal(suite.T(), 3, len(results))
}

func (suite *SimpleCSVReaderTestSuite) TestReadFileWithCorrectColumns() {
	results, err := suite.target.ReadFile(suite.testFilePath)

	assert.Nil(suite.T(), err)

	// the number 5 should be equal to the columns number in file ./csvtestingfile.csv
	assert.Equal(suite.T(), 5, len(results[0]))
}

func TestSimpleCSVReaderTestSuite(t *testing.T) {
	suite.Run(t, new(SimpleCSVReaderTestSuite))
}
