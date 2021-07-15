package autotagbuilder

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type AutoTagProcessorBuilderTestSuite struct {
	suite.Suite
}

func (suite *AutoTagProcessorBuilderTestSuite) SetupTest() {

}

func (suite *AutoTagProcessorBuilderTestSuite) TestBuilderWithNoNil() {
	//tagProcessor := BuildProcessorRuler()
	//assert.NotNil(suite.T(), tagProcessor)
}

func TestAutoTagProcessorBuilderTestSuite(t *testing.T) {
	suite.Run(t, new(AutoTagProcessorBuilderTestSuite))
}
