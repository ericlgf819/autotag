package tagrules

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TagRulerTestSuite struct {
	suite.Suite
	target         *TagRuler
	calculatedData [][]string
}

type MockRawRuleDataReader struct{}

func (*MockRawRuleDataReader) FetchRawData() ([]map[string]string, error) {
	return []map[string]string{
		{"col1": "hello", "col2": "world", "col3": "i'm row0 value", "col4": "don't care"},
		{"col1": "good", "col2": "morning", "col3": "i'm row1 value", "col4": "don't care"},
		{"col1": "look", "col2": "there", "col3": "doesnt' care"},
		{"a": "b"},
		{"a": "b"},
		{"a": "b"},
		{"a": "b"},
	}, nil
}

type MockRuleDataWriter struct {
	suite *TagRulerTestSuite
}

func (writer *MockRuleDataWriter) StoreRuleData(finalData [][]string) error {
	writer.suite.calculatedData = finalData
	return nil
}

func (suite *TagRulerTestSuite) SetupTest() {
	suite.target = new(TagRuler)

	writer := new(MockRuleDataWriter)
	writer.suite = suite

	reader := new(MockRawRuleDataReader)

	suite.target.Init(reader, writer)
	suite.target.InitRules(
		[]string{"col1", "col2"}, "col3",
		"TagName", "TagValue")
}

func (suite *TagRulerTestSuite) TestMakeRulesWithNoError() {
	err := suite.target.MakeRules()
	assert.Nil(suite.T(), err)
}

func (suite *TagRulerTestSuite) TestMakeRulesWithNormalOutputCount() {
	suite.target.MakeRules()
	suite.target.StoreRules()

	assert.Equal(suite.T(), 4, len(suite.calculatedData))
	assert.Equal(suite.T(), 2, len(suite.calculatedData[1]))
}

func (suite *TagRulerTestSuite) TestMakeRulesWithNormalOutputRow() {
	suite.target.MakeRules()
	suite.target.StoreRules()

	assert.Equal(suite.T(), "hello|world", suite.calculatedData[1][0])
	assert.Equal(suite.T(), "i'm row0 value", suite.calculatedData[1][1])
}

func TestTagRulerTestSuite(t *testing.T) {
	suite.Run(t, new(TagRulerTestSuite))
}
