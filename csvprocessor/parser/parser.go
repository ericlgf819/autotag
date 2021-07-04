package parser

type ManifestReader interface {
	Get(key string) (string, error)
}

type ColumnValuePair struct {
	column string
	value  string
}

type Parser interface {
	SetManifestReader(reader ManifestReader)
	GetManifestReader() ManifestReader
	LookupTag(columnsAndValues []ColumnValuePair) (string, error)
}

type AzureCSVParser struct {
	manifestReader ManifestReader
}

func (parser *AzureCSVParser) SetManifestReader(reader ManifestReader) {
	parser.manifestReader = reader
}

func (parser *AzureCSVParser) GetManifestReader() ManifestReader {
	return parser.manifestReader
}

func (parser *AzureCSVParser) LookupTag(columnsAndValues []ColumnValuePair) (string, error) {
	keyColumnNames := map[string]bool{"testColumn1": true, "testColumn2": true}

	compoundKeyName := ""
	for _, pair := range columnsAndValues {
		if keyColumnNames[pair.column] {
			compoundKeyName += pair.value + "|"
		}
	}

	return parser.GetManifestReader().Get(compoundKeyName)
}
