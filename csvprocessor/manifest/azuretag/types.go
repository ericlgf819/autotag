package azuretagcsvreader

type AzureTagCSVDecoratorReader struct {
	CSVFileReader CSVReader
}

type CSVReader interface {
	ReadFile(path string) ([][]string, error)
}
