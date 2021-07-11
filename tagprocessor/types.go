package tagprocessor

type TagProcessor struct {
	csvReader CSVReader
	csvWriter CSVWriter
	parser    Parser
}

type ColumnValuePair struct {
	Column string
	Value  string
}

type CSVReader interface {
	ReadFile(path string) ([][]string, error)
}

type CSVWriter interface {
	WriteFile(path string, content [][]string) error
}

type Parser interface {
	LookupTag(columnsAndValues []ColumnValuePair) (string, error)
}
