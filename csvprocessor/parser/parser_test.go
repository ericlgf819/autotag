package parser

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CSVProcessorTestSuite struct {
	suite.Suite
	inputColumnValuePairs map[string]string
	expectedTagName       string
}

type mockReader struct{}

func (r *mockReader) Get(key string) (string, error) {
	if key == "test1|test2|" {
		return "value1", nil
	}

	return "", fmt.Errorf("key '%s' does not exist", key)
}

func (suite *CSVProcessorTestSuite) SetupTest() {
	suite.inputColumnValuePairs = map[string]string{
		"testColumn1": "test1",
		"testColumn2": "test2",
	}
	suite.expectedTagName = "value1"
	Init(new(mockReader))
}

func (suite *CSVProcessorTestSuite) TestLookupTagWithCorrectInputPairs() {
	result, err := LookupTag(suite.inputColumnValuePairs)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), suite.expectedTagName, result)
}

func TestCSVProcessorTestSuite(t *testing.T) {
	suite.Run(t, new(CSVProcessorTestSuite))
}
