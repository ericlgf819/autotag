package autotagbuilder

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AutoTagRulerBuilderTestSuite struct {
	suite.Suite
}

func (suite *AutoTagRulerBuilderTestSuite) SetupTest() {

}

func (suite *AutoTagRulerBuilderTestSuite) TestBuilderWithNoNil() {
	tagRuler := BuildTagRuler()
	assert.NotNil(suite.T(), tagRuler)
}

func TestAutoTagRulerBuilderTestSuite(t *testing.T) {
	suite.Run(t, new(AutoTagRulerBuilderTestSuite))
}
