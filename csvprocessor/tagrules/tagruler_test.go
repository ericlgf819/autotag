package tagrules

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TagRulerTestSuite struct {
	suite.Suite
	target *TagRuler
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

type MockRuleDataWriter struct{}

func (*MockRuleDataWriter) StoreRuleData([][]string) error {
	return nil
}

func (suite *TagRulerTestSuite) SetupTest() {
	suite.target = new(TagRuler)
	suite.target.Init(new(MockRawRuleDataReader), new(MockRuleDataWriter))
	suite.target.InitRules(
		[]string{"col1", "col2"}, "col3",
		"TagName", "TagValue")
}

func (suite *TagRulerTestSuite) TestMakeRulesWithNoError() {
	err := suite.target.MakeRules()
	assert.NotNil(suite.T(), err)
}

func TestTagRulerTestSuite(t *testing.T) {
	suite.Run(t, new(TagRulerTestSuite))
}
