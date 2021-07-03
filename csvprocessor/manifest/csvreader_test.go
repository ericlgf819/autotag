package csvreader

import (
	//"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ManifestCSVReaderImplTestSuite struct {
	suite.Suite
	testFilePath string
}

func (suite *ManifestCSVReaderImplTestSuite) SetupTest() {
	suite.testFilePath = ""
}

func (suite *ManifestCSVReaderImplTestSuite) TestReadFileNormally() {
	dic, err := ReadFile(suite.testFilePath)
	assert.NotNil(suite.T(), dic)
	assert.Nil(suite.T(), err)
}

func TestManifestReaderImplTestSuite(t *testing.T) {
	suite.Run(t, new(ManifestCSVReaderImplTestSuite))
}
