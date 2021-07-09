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

/*
type AzureTagReader interface {
	ReadFile(path string) (map[string]string, error)
}
*/
type MockAzureTagReader struct{}

func (*MockAzureTagReader) ReadFile(string) (map[string]string, error) {
	return map[string]string{
		"test0": "result0",
		"test1": "result1",
		"test2": "result2",
	}, nil
}

func (suite *SimpleManifestReaderTestSuite) SetupTest() {
	suite.target = new(SimpleManifestReader)
	suite.target.Init("mockpath", new(MockAzureTagReader))
}

func (suite *SimpleManifestReaderTestSuite) TestSimpleManifestReaderInit() {
	err := suite.target.Init("mockpath", new(MockAzureTagReader))
	assert.Nil(suite.T(), err)
}

func (suite *SimpleManifestReaderTestSuite) TestSimpleManifestReaderGetWithResult() {
	res, _ := suite.target.Get("test1")

	assert.Equal(suite.T(), "result1", res)
}

func TestSimpleManifestReaderTestSuite(t *testing.T) {
	suite.Run(t, new(SimpleManifestReaderTestSuite))
}
