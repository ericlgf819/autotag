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
	suite.testFilePath = ""
	suite.target = new(SimpleCSVReader)
}

func (suite *SimpleCSVReaderTestSuite) TestReadFileNormally() {
	dic, err := suite.target.ReadFile(suite.testFilePath)
	assert.NotNil(suite.T(), dic)
	assert.Nil(suite.T(), err)
}

func TestSimpleCSVReaderTestSuite(t *testing.T) {
	suite.Run(t, new(SimpleCSVReaderTestSuite))
}
