package simplerawruledatareader

import (
	"errors"
	"fmt"
)

func (reader *SimpleRawRuleDataReader) Init(fileReader CSVFileReader) {
	reader.csvFileReader = fileReader
}

func (reader *SimpleRawRuleDataReader) LoadFile(path string) error {
	var err error = nil
	reader.fileContent, err = reader.csvFileReader.ReadFile(path)
	return err
}

func (reader *SimpleRawRuleDataReader) FetchRawData() ([]map[string]string, error) {
	if len(reader.fileContent) == 0 {
		return nil, errors.New("no file loaded")
	}

	headerRow := reader.fileContent[0]

	resultDic := []map[string]string{}

	for rowNumber, rowContent := range reader.fileContent[1:] {
		if !isColumnCountEqual(headerRow, rowContent) {
			return nil, fmt.Errorf("column count in row %d doesn't match header's %d", rowNumber, len(headerRow))
		}

		resultDic = append(resultDic, buildRowContent(headerRow, rowContent))
	}
	return resultDic, nil
}

func isColumnCountEqual(headerRow []string, contentRow []string) bool {
	return len(headerRow) == len(contentRow)
}

func buildRowContent(headerRow []string, contentRow []string) map[string]string {
	resultDicRow := map[string]string{}
	for index := range headerRow {
		resultDicRow[headerRow[index]] = contentRow[index]
	}
	return resultDicRow
}
