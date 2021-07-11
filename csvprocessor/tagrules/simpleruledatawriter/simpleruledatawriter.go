package simpleruledatawriter

func (simpleWriter *SimpleRuleDataWriter) Init(filePath string, writer CSVFileWriter) {
	simpleWriter.filePath = filePath
	simpleWriter.fileWriter = writer
}

func (simpleWriter *SimpleRuleDataWriter) StoreRuleData(content [][]string) error {
	return simpleWriter.fileWriter.WriteFile(simpleWriter.filePath, content)
}
