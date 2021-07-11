package simpleruledatawriter

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SimpleRuleDataWriterTestSuite struct {
	suite.Suite
	target         *SimpleRuleDataWriter
	contentToStore [][]string
}

type MockCSVFileWriter struct{}

func (*MockCSVFileWriter) WriteFile(path string, content [][]string) error {
	return nil
}

func (suite *SimpleRuleDataWriterTestSuite) SetupTest() {
	suite.target = new(SimpleRuleDataWriter)
	suite.target.Init("mockpath", new(MockCSVFileWriter))
}

func (suite *SimpleRuleDataWriterTestSuite) TestStoreRuleDataNormally() {
	err := suite.target.StoreRuleData(suite.contentToStore)
	assert.Nil(suite.T(), err)
}

func TestSimpleRuleDataWriterTestSuite(t *testing.T) {
	suite.Run(t, new(SimpleRuleDataWriterTestSuite))
}
