package simplerawruledatareader

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type SimpleRawRuleDataReaderTestSuite struct {
	suite.Suite
}

func (suite *SimpleRawRuleDataReaderTestSuite) SetupTest() {

}

func (suite *SimpleRawRuleDataReaderTestSuite) Test() {

}

func TestSimpleRawRuleDataReaderTestSuite(t *testing.T) {
	suite.Run(t, new(SimpleRawRuleDataReaderTestSuite))
}
