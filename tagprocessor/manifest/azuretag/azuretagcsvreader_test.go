package azuretagcsvreader

import (
	"testing"

	"github.com/ericlgf819/autotag/tagprocessor/tagrules"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AzureTagCSVReaderTestSuite struct {
	suite.Suite
	testFilePath string
	target       *AzureTagCSVDecoratorReader
}

type MockerCSVReader struct{}

func (*MockerCSVReader) ReadFile(string) ([][]string, error) {
	return [][]string{
			{"TCol0", "TCol1", tagrules.TagKeyName, "TCol2", tagrules.TagValueName},
			{"TRow0Col0", "TRow0Col1", "testColumn1|testColumn2", "TRow0Col2", "TRow0Col3TagValue"},
			{"TRow1Col0", "TRow1Col1", "testColumn2|testColumn1", "TRow1Col2", "TRow1Col3TagValue"},
		},
		nil
}

func (suite *AzureTagCSVReaderTestSuite) SetupTest() {
	suite.testFilePath = "mockpath"
	suite.target = new(AzureTagCSVDecoratorReader)

	suite.target.Init(new(MockerCSVReader))
}

func (suite *AzureTagCSVReaderTestSuite) TestReadFileNormallyWithoutError() {
	res, err := suite.target.ReadFile(suite.testFilePath)
	assert.NotNil(suite.T(), res)
	assert.Nil(suite.T(), err)
}

func (suite *AzureTagCSVReaderTestSuite) TestReadFileResultRowShouldBeInsideSourceRows() {
	res, err := suite.target.ReadFile(suite.testFilePath)
	assert.Nil(suite.T(), err)

	mockCSVReader := new(MockerCSVReader)
	source, _ := mockCSVReader.ReadFile(suite.testFilePath)

	isInSource := false
	inSourceTick := 0
	for _, row := range source[1:] {
		isInSource = false
		for k, v := range res {
			existingTick := 0
			for _, sourceColVal := range row {
				if k == sourceColVal || v == sourceColVal {
					existingTick++
				}
			}
			if existingTick > 1 && existingTick < 2*len(row)+1 {
				isInSource = true
			}
		}
		if isInSource {
			inSourceTick++
		}
	}

	assert.Equal(suite.T(), len(res), inSourceTick)
}

func TestAzureTagCSVReaderTestSuite(t *testing.T) {
	suite.Run(t, new(AzureTagCSVReaderTestSuite))
}
