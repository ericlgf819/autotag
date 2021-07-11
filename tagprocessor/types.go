package tagprocessor

type TagProcessor struct {
	reader ContentReader
	writer ContentWriter
	parser TagParser
}

type ColumnValuePair struct {
	Column string
	Value  string
}

type ContentReader interface {
	ReadFile(path string) ([][]string, error)
}

type ContentWriter interface {
	WriteFile(path string, content [][]string) error
}

type TagParser interface {
	LookupTag(columnsAndValues []ColumnValuePair) (string, error)
}
