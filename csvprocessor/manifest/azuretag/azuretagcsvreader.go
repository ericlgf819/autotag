package azuretagcsvreader

import "errors"

func (azReader *AzureTagCSVDecoratorReader) Init(reader CSVReader) {
	azReader.CSVFileReader = reader
}

func (azReader *AzureTagCSVDecoratorReader) getCSVReader() CSVReader {
	return azReader.CSVFileReader
}

func (azReader *AzureTagCSVDecoratorReader) ReadFile(path string) (map[string]string, error) {
	tagKeyColumnName := "TagKeys"
	tagValueColumnName := "TagValues"

	fileCsvTableContent, err := getCsvfileInTable(path, azReader.getCSVReader())

	if err != validateFileHasContent(fileCsvTableContent, err) {
		return nil, err
	}

	tagKeyColumnIndex, err := findTargetColumnIndex(tagKeyColumnName, fileCsvTableContent[0])

	if err != nil {
		return nil, err
	}

	tagValueColumnIndex, err := findTargetColumnIndex(tagValueColumnName, fileCsvTableContent[0])

	if err != nil {
		return nil, err
	}

	tagDicionary := createDictionaryByTable(tagKeyColumnIndex, tagValueColumnIndex, fileCsvTableContent[1:])

	return tagDicionary, nil
}

func validateFileHasContent(content [][]string, err error) error {
	if err != nil {
		return err
	}

	if len(content) == 0 {
		return errors.New("no content")
	}

	return nil
}

func findTargetColumnIndex(targetColName string, headerRow []string) (int, error) {
	for columnIndex, columnName := range headerRow {
		if columnName == targetColName {
			return columnIndex, nil
		}
	}
	return -1, errors.New("no target column")
}

func createDictionaryByTable(keyColumnIndex, valueColumnIndex int, table [][]string) map[string]string {
	resultDictionary := make(map[string]string)

	for _, rowInFileCsvTableContent := range table[1:] {
		resultDictionary[rowInFileCsvTableContent[keyColumnIndex]] = rowInFileCsvTableContent[valueColumnIndex]
	}

	return resultDictionary
}

func getCsvfileInTable(path string, reader CSVReader) ([][]string, error) {
	return reader.ReadFile(path)
}
