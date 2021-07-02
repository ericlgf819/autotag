package parser

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CSVProcessorTestSuite struct {
	suite.Suite
	inputColumnValuePairsCorrectly []ColumnValuePair
	inputColumnValuePairsWrongly   []ColumnValuePair
	expectedTagName                string
}

type mockReader struct{}

func (*mockReader) Get(key string) (string, error) {
	if key == "test1|test2|" {
		return "value1", nil
	}

	return "", fmt.Errorf("key '%s' does not exist", key)
}

func (suite *CSVProcessorTestSuite) SetupTest() {
	suite.inputColumnValuePairsCorrectly = []ColumnValuePair{
		{"testColumn1", "test1"},
		{"testColumn2", "test2"},
	}

	suite.inputColumnValuePairsWrongly = []ColumnValuePair{
		{"testColumn3", "test1"},
		{"testColumn2", "test2"},
	}

	suite.expectedTagName = "value1"
	Init(new(mockReader))
}

func (suite *CSVProcessorTestSuite) TestLookupTagWithCorrectInputPairs() {
	result, err := LookupTag(suite.inputColumnValuePairsCorrectly)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), suite.expectedTagName, result)
}

func (suite *CSVProcessorTestSuite) TestLookupTagWithWrongInputPairs() {
	result, err := LookupTag(suite.inputColumnValuePairsWrongly)

	assert.Error(suite.T(), err)
	assert.NotEqual(suite.T(), suite.expectedTagName, result)
}

func TestCSVProcessorTestSuite(t *testing.T) {
	suite.Run(t, new(CSVProcessorTestSuite))
}
