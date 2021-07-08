package csvreader

func (simpleReader *SimpleCSVReader) SetDelimeter(delimeter string) {
	simpleReader.delimeter = delimeter
}

func (simpleReader *SimpleCSVReader) GetDelimeter() string {
	return simpleReader.delimeter
}

func (simpleReader *SimpleCSVReader) ReadFile(path string) ([][]string, error) {
	simpleReader.setDefaultDelimeterIfNil()
	return nil, nil
}

func (simpleReader *SimpleCSVReader) setDefaultDelimeterIfNil() {
	if simpleReader.delimeter == "" {
		simpleReader.delimeter = ","
	}
}
