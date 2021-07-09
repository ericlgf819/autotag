package azuretagcsvreader

type AzureTagCSVDecoratorReader struct {
	CSVFileReader CSVReader
}

type CSVReader interface {
	ReadFile(path string) ([][]string, error)
}

type AzureTagCSVReader interface {
	ReadFile(path string) (map[string]string, error)
	Init(reader CSVReader)
	getCSVReader() CSVReader
}
