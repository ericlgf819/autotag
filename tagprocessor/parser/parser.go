package parser

import "github.com/ericlgf819/autotag/tagprocessor"

func (parser *AzureCSVParser) Init(reader ManifestReader) {
	parser.manifestReader = reader
}

func (parser *AzureCSVParser) getManifestReader() ManifestReader {
	return parser.manifestReader
}

func (parser *AzureCSVParser) LookupTag(columnsAndValues []tagprocessor.ColumnValuePair) (string, error) {
	keyColumnNames := map[string]bool{"testColumn1": true, "testColumn2": true}

	compoundKeyName := ""
	for _, pair := range columnsAndValues {
		if keyColumnNames[pair.Column] {
			compoundKeyName += pair.Value + "|"
		}
	}

	return parser.getManifestReader().Get(compoundKeyName)
}
