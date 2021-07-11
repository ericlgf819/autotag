package simplerawruledatareader

import "errors"

func (reader *SimpleRawRuleDataReader) Init(fileReader CSVFileReader) {
	reader.csvFileReader = fileReader
}

func (reader *SimpleRawRuleDataReader) LoadFile(path string) error {

	return errors.New("not implement")
}
