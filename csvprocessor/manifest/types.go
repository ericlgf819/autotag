package manifest

type SimpleManifestReader struct {
	sourceFilePath string
}

type AzureTagReader interface {
	ReadFile(path string) (map[string]string, error)
}
