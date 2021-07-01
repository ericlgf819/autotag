package parser

type ManifestReader interface {
	Get(key string) (string, error)
}

var manifestReader ManifestReader

func Init(reader ManifestReader) {
	manifestReader = reader
}

func LookupTag(columnsAndValues map[string]string) (string, error) {
	keyColumnNames := map[string]bool{"testColumn1": true, "testColumn2": true}

	compoundKeyName := ""
	for k, v := range columnsAndValues {
		if keyColumnNames[k] {
			compoundKeyName += v + "|"
		}
	}

	return manifestReader.Get(compoundKeyName)
}
