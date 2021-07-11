package parser

func (parser *AzureCSVParser) Init(reader ManifestReader) {
	parser.manifestReader = reader
}

func (parser *AzureCSVParser) getManifestReader() ManifestReader {
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

	return parser.getManifestReader().Get(compoundKeyName)
}
