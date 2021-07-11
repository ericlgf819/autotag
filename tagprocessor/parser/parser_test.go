package parser

import (
	"fmt"
	"testing"

	"github.com/ericlgf819/autotag/tagprocessor"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CSVProcessorTestSuite struct {
	suite.Suite
	inputColumnValuePairsCorrectly []tagprocessor.ColumnValuePair
	inputColumnValuePairsWrongly   []tagprocessor.ColumnValuePair
	expectedTagName                string
	target                         *AzureCSVParser
}

type mockReader struct{}

func (*mockReader) Get(key string) (string, error) {
	if key == "test1|test2|" {
		return "value1", nil
	}

	return "", fmt.Errorf("key '%s' does not exist", key)
}

func (suite *CSVProcessorTestSuite) SetupTest() {
	suite.inputColumnValuePairsCorrectly = []tagprocessor.ColumnValuePair{

		{Column: "testColumn1", Value: "test1"},
		{Column: "testColumn2", Value: "test2"},
	}

	suite.inputColumnValuePairsWrongly = []tagprocessor.ColumnValuePair{
		{Column: "testColumn3", Value: "test1"},
		{Column: "testColumn2", Value: "test2"},
	}

	suite.expectedTagName = "value1"
	suite.target = new(AzureCSVParser)

	suite.target.Init(new(mockReader))
}

func (suite *CSVProcessorTestSuite) TestLookupTagWithCorrectInputPairs() {
	result, err := suite.target.LookupTag(suite.inputColumnValuePairsCorrectly)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), suite.expectedTagName, result)
}

func (suite *CSVProcessorTestSuite) TestLookupTagWithWrongInputPairs() {
	result, err := suite.target.LookupTag(suite.inputColumnValuePairsWrongly)

	assert.Error(suite.T(), err)
	assert.NotEqual(suite.T(), suite.expectedTagName, result)
}

func TestCSVProcessorTestSuite(t *testing.T) {
	suite.Run(t, new(CSVProcessorTestSuite))
}
