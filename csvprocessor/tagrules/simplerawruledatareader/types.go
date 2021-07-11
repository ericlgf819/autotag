package simplerawruledatareader

type SimpleRawRuleDataReader struct {
	csvFileReader CSVFileReader
	fileContent   [][]string
}

type CSVFileReader interface {
	ReadFile(path string) ([][]string, error)
}
