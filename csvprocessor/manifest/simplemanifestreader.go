package manifest

func (reader *SimpleManifestReader) Init(path string, concreteReader AzureTagReader) error {
	reader.sourceFilePath = path
	reader.azureTagReader = concreteReader

	var err error
	reader.tagDataStore, err = concreteReader.ReadFile(path)

	return err
}

func (reader *SimpleManifestReader) Get(key string) (string, error) {
	return reader.tagDataStore[key], nil
}
