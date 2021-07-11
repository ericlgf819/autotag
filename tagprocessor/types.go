package tagprocessor

type TagProcessor struct {
	csvReader CSVReader
}

type ColumnValuePair struct {
	Column string
	Value  string
}

type CSVReader interface {
	ReadFile(path string) ([][]string, error)
}
