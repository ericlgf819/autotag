package azuretagcsvreader

type CSVReader interface {
	ReadFile(path string) ([][]string, error)
	//ReadFileAsync(path string, header chan []string, row chan []string) error
}

type AzureTagReader interface {
	ReadFile(path string) (map[string]string, error)
	Init(reader CSVReader)
	getCSVReader() CSVReader
}

type AzureTagCSVDecoratorReader struct {
	CSVFileReader CSVReader
}
