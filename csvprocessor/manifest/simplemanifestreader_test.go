package manifest

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SimpleManifestReaderTestSuite struct {
	suite.Suite
	target *SimpleManifestReader
}

func (suite *SimpleManifestReaderTestSuite) SetupTest() {
	suite.target = new(SimpleManifestReader)
	suite.target.Init("./manifesttestingfile.csv")
}

func (suite *SimpleManifestReaderTestSuite) TestSimpleManifestReaderInit() {
	assert.NotNil(suite.T(), nil)
}

// func (suite *SimpleManifestReaderTestSuite) TestSimpleManifestReaderGetWithResult() {
// 	assert.NotNil(suite.T(), nil)
// }

func TestSimpleManifestReaderTestSuite(t *testing.T) {
	suite.Run(t, new(SimpleManifestReaderTestSuite))
}
