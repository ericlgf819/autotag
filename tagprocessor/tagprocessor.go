package tagprocessor

import "errors"

func (processor *TagProcessor) Init(reader ContentReader, writer ContentWriter, parser TagParser) {
	processor.reader = reader
	processor.writer = writer
	processor.parser = parser
	processor.tagColumnName = "Tag"
}

func (processor *TagProcessor) Process(inputFilePath string, outputFilePath string) error {
	contentTable, err := processor.reader.ReadFile(inputFilePath)
	if nil != err {
		return err
	}

	err = checkContentTableEmpty(contentTable)
	if nil != err {
		return err
	}

	headerRow := contentTable[0]
	contentRows := contentTable[1:]

	outputContent := processor.buildOutputContentHeader(headerRow)

	colValPairs := convertArrayRowToColumnValuePairs(headerRow, contentRows)

	outputContent, err = processor.buildOutputContent(colValPairs, contentRows, outputContent)

	if nil != err {
		return err
	}

	err = processor.writer.WriteFile(outputFilePath, outputContent)
	return err
}

func checkContentTableEmpty(contentTable [][]string) error {
	if len(contentTable) < 2 {
		return errors.New("no content")
	}
	return nil
}

func convertArrayRowToColumnValuePairs(header []string, contentRows [][]string) [][]ColumnValuePair {
	result := [][]ColumnValuePair{}
	row := []ColumnValuePair{}

	for _, contentRow := range contentRows {
		for index := range header {
			cell := ColumnValuePair{Column: header[index], Value: contentRow[index]}
			row = append(row, cell)
		}
		result = append(result, row)
	}

	return result
}

func (processor *TagProcessor) buildOutputContentHeader(headerRow []string) [][]string {
	result := [][]string{}
	resultHeaderRow := append(headerRow, processor.tagColumnName)
	result = append(result, resultHeaderRow)
	return result
}

func (processor *TagProcessor) buildOutputContent(colValPairs [][]ColumnValuePair, contentRows [][]string, outputContent [][]string) ([][]string, error) {
	var err error
	var tagVal string

	for index, colValPair := range colValPairs {
		tagVal, err = processor.parser.LookupTag(colValPair)
		if nil != err {
			return nil, err
		}

		outputRow := append(contentRows[index], tagVal)
		outputContent = append(outputContent, outputRow)
	}

	return outputContent, nil
}
