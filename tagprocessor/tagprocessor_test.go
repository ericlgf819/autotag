package tagprocessor

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TagProcessorTestSuite struct {
	suite.Suite
	target         *TagProcessor
	inputFilePath  string
	outputFilePath string
}

func (suite *TagProcessorTestSuite) SetupTest() {
	suite.target = new(TagProcessor)
}

func (suite *TagProcessorTestSuite) TestProcessNormally() {
	err := suite.target.Process(suite.inputFilePath, suite.outputFilePath)

	assert.Nil(suite.T(), err)
}

func TestTagProcessorTestSuite(t *testing.T) {
	suite.Run(t, new(TagProcessorTestSuite))
}
