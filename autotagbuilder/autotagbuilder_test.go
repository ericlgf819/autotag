package autotagbuilder

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AutoTagBuilderTestSuite struct {
	suite.Suite
}

func (suite *AutoTagBuilderTestSuite) SetupTest() {

}

func (suite *AutoTagBuilderTestSuite) TestBuilderWithNoNil() {
	tagRuler := BuildTagRuler()
	assert.NotNil(suite.T(), tagRuler)
}

func TestAutoTagBuilderTestSuite(t *testing.T) {
	suite.Run(t, new(AutoTagBuilderTestSuite))
}
