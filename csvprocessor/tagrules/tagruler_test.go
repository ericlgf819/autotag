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

func (suite *TagRulerTestSuite) SetupTest() {
	suite.target = new(TagRuler)
	suite.target.Init()
}

func (suite *TagRulerTestSuite) TestMakeRulesWithNoError() {
	err := suite.target.MakeRules()
	assert.NotNil(suite.T(), err)
}

func TestTagRulerTestSuite(t *testing.T) {
	suite.Run(t, new(TagRulerTestSuite))
}
