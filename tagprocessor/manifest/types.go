package manifest

type AzureTagReader interface {
	ReadFile(path string) (map[string]string, error)
}

type SimpleManifestReader struct {
	sourceFilePath string
	azureTagReader AzureTagReader
	tagDataStore   map[string]string
}
