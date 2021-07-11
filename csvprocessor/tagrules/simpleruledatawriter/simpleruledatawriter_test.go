package simpleruledatawriter

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type SimpleRuleDataWriterTestSuite struct {
	suite.Suite
	target *SimpleRuleDataWriter
}

func (suite *SimpleRuleDataWriterTestSuite) SetupTest() {
	suite.target = new(SimpleRuleDataWriter)
}

func (suite *SimpleRuleDataWriterTestSuite) TestStoreRuleDataNormally() {

}

func TestSimpleRuleDataWriterTestSuite(t *testing.T) {
	suite.Run(t, new(SimpleRuleDataWriterTestSuite))
}
