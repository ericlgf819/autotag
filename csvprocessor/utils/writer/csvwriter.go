package csvwriter

func (simpleReader *SimpleCSVWriter) SetDelimiter(delimeter int32) {
	simpleReader.delimiter = delimeter
}

func (simpleReader *SimpleCSVWriter) GetDelimiter() int32 {
	return simpleReader.delimiter
}
