package parser

type ManifestReader interface {
	Get(key string) (string, error)
}

type AzureCSVParser struct {
	manifestReader ManifestReader
}
