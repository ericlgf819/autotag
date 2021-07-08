package parser

type ManifestReader interface {
	Get(key string) (string, error)
}

type ColumnValuePair struct {
	column string
	value  string
}

type Parser interface {
	Init(reader ManifestReader)
	LookupTag(columnsAndValues []ColumnValuePair) (string, error)
	getManifestReader() ManifestReader
}

type AzureCSVParser struct {
	manifestReader ManifestReader
}
