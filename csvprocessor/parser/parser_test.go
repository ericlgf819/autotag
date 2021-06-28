package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CSVProcessorTestSuite struct {
	suite.Suite
	inputColumnValuePairs map[string]string
	expectedTagName       string
}

func (suite *CSVProcessorTestSuite) SetupTest() {
	suite.inputColumnValuePairs = map[string]string{
		"": "",
	}
	suite.expectedTagName = ""
}

func (suite *CSVProcessorTestSuite) TestLookupTagWithCorrectInputPairs() {
	result, err := LookupTag(suite.inputColumnValuePairs)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), suite.expectedTagName, result)
}

func TestCSVProcessorTestSuite(t *testing.T) {
	suite.Run(t, new(CSVProcessorTestSuite))
}
