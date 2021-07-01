package parser

type ManifestReader interface {
	Get(key string) (string, error)
}

var manifestReader ManifestReader

func Init(reader ManifestReader) {
	manifestReader = reader
}

func LookupTag(colValPairs map[string]string) (string, error) {
	keyColumnNames := map[string]bool{"testColumn1": true, "testColumn2": true}

	compoundKeyName := ""
	for _, k := range colValPairs {
		colName := k
		colValue := colValPairs[colName]

		if keyColumnNames[colName] {
			compoundKeyName += colValue + "|"
		}
	}

	return manifestReader.Get(compoundKeyName)
}
