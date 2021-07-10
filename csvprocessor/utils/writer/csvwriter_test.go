package csvwriter

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SimpleCSVWriterTestSuite struct {
	suite.Suite
	testFilePath string
	target       *SimpleCSVWriter
}

func (suite *SimpleCSVWriterTestSuite) SetupTest() {
	suite.testFilePath = "./csvtestingfile.csv"
	suite.target = new(SimpleCSVWriter)
}

func (suite *SimpleCSVWriterTestSuite) TestWriteFileWithoutError() {

	assert.NotNil(suite.T(), nil)
}

func TestSimpleCSVWriterTestSuite(t *testing.T) {
	suite.Run(t, new(SimpleCSVWriterTestSuite))
}
