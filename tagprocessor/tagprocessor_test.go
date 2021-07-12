package tagprocessor

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TagProcessorTestSuite struct {
	suite.Suite
	target                *TagProcessor
	inputMockContent      [][]string
	expectedOutputContent [][]string
	actualOutputContent   [][]string
}

type MockReaderWriterParser struct {
	suite *TagProcessorTestSuite
}

func (mock *MockReaderWriterParser) ReadFile(path string) ([][]string, error) {
	return mock.suite.inputMockContent, nil
}

func (mock *MockReaderWriterParser) WriteFile(path string, content [][]string) error {
	mock.suite.actualOutputContent = content
	return nil
}

func (mock *MockReaderWriterParser) LookupTag(columnsAndValues []ColumnValuePair) (string, error) {
	if columnsAndValues[1].Column != "col1" || columnsAndValues[3].Column != "col3" {
		return "", errors.New("not found")
	}

	compoundValue := fmt.Sprintf("%s|%s", columnsAndValues[1].Value, columnsAndValues[3].Value)

	if compoundValue == "world|morning" {
		return "good", nil
	}

	if compoundValue == "to|you" {
		return "job", nil
	}

	return "", errors.New("not found")
}

func (suite *TagProcessorTestSuite) SetupTest() {
	suite.target = new(TagProcessor)
	mockReaderWriterParser := &MockReaderWriterParser{suite: suite}
	suite.target.Init(mockReaderWriterParser, mockReaderWriterParser, mockReaderWriterParser)

	suite.inputMockContent = [][]string{
		{"col0", "col1", "col2", "col3"},
		{"hello", "world", "good", "morning"},
		{"nice", "to", "see", "you"},
	}
	suite.expectedOutputContent = [][]string{
		{"col0", "col1", "col2", "col3", "Tag"},
		{"hello", "world", "good", "morning", "good"},
		{"nice", "to", "see", "you", "job"},
	}
}

func (suite *TagProcessorTestSuite) TestProcessNormally() {
	err := suite.target.Process("mockInputPath", "mockOutputPath")

	assert.Nil(suite.T(), err)
	assert.True(suite.T(), arrEqual(suite.actualOutputContent, suite.expectedOutputContent))
}

func arrEqual(source, dest [][]string) bool {
	if len(source) != len(dest) {
		return false
	}

	for rowIndex := range source {
		if len(source[rowIndex]) != len(dest[rowIndex]) {
			return false
		}

		for colIndex := range source[rowIndex] {
			if len(source[rowIndex][colIndex]) != len(dest[rowIndex][colIndex]) {
				return false
			}
		}
	}
	return true
}

func TestTagProcessorTestSuite(t *testing.T) {
	suite.Run(t, new(TagProcessorTestSuite))
}
