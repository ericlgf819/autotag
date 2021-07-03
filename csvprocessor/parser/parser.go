package parser

type ManifestReader interface {
	Get(key string) (string, error)
}

type ColumnValuePair struct {
	column string
	value  string
}

func Init(reader ManifestReader) {
	manifestReader = reader
}

func LookupTag(columnsAndValues []ColumnValuePair) (string, error) {
	keyColumnNames := map[string]bool{"testColumn1": true, "testColumn2": true}

	compoundKeyName := ""
	for _, pair := range columnsAndValues {
		if keyColumnNames[pair.column] {
			compoundKeyName += pair.value + "|"
		}
	}

	return manifestReader.Get(compoundKeyName)
}

var manifestReader ManifestReader
