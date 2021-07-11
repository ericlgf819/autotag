package simplerawruledatareader

func (reader *SimpleRawRuleDataReader) Init(fileReader CSVFileReader) {
	reader.csvFileReader = fileReader
}

func (reader *SimpleRawRuleDataReader) LoadFile(path string) error {
	var err error = nil
	reader.fileContent, err = reader.csvFileReader.ReadFile(path)
	return err
}
