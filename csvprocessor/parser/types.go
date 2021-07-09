package parser

type ManifestReader interface {
	Get(key string) (string, error)
}

type ColumnValuePair struct {
	column string
	value  string
}

type AzureCSVParser struct {
	manifestReader ManifestReader
}
