package tagprocessor

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TagProcessorTestSuite struct {
	suite.Suite
	target *TagProcessor
}

func (suite *TagProcessorTestSuite) SetupTest() {
	suite.target = new(TagProcessor)
}

func (suite *TagProcessorTestSuite) Test() {

}

func TestTagProcessorTestSuite(t *testing.T) {
	suite.Run(t, new(TagProcessorTestSuite))
}
