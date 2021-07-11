package simpleruledatawriter

type SimpleRuleDataWriter struct {
	fileWriter CSVFileWriter
	filePath   string
}

type CSVFileWriter interface {
	WriteFile(path string, content [][]string) error
}
