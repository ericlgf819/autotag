package csvwriter

import "os"

func (simpleWriter *SimpleCSVWriter) SetDelimiter(delimeter int32) {
	simpleWriter.delimiter = delimeter
}

func (simpleWriter *SimpleCSVWriter) GetDelimiter() int32 {
	return simpleWriter.delimiter
}

func (simpleWriter *SimpleCSVWriter) AppendColumn(path string, value string, lineNum int) error {
	simpleWriter.setDefaultDelimiterIfNil()

	err := checkFileError(path)
	if err != nil {
		return err
	}

	return nil
}

func checkFileError(path string) error {
	_, err := os.Stat(path)
	return err
}

func (simpleReader *SimpleCSVWriter) setDefaultDelimiterIfNil() {
	if simpleReader.delimiter == 0 {
		simpleReader.delimiter = ','
	}
}
