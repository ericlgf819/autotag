package csvwriter

import (
	"encoding/csv"
	"os"
)

func (simpleWriter *SimpleCSVWriter) SetDelimiter(delimeter int32) {
	simpleWriter.delimiter = delimeter
}

func (simpleWriter *SimpleCSVWriter) GetDelimiter() int32 {
	return simpleWriter.delimiter
}

func (simpleWriter *SimpleCSVWriter) WriteFile(path string, content [][]string) error {
	simpleWriter.setDefaultDelimiterIfNil()

	file, err := createFileIfNotExist(path)
	if nil != err {
		return err
	}

	err = simpleWriter.writeAll(content, file)
	if nil != err {
		return err
	}

	closeErr := file.Close()
	return closeErr
}

func (simpleWriter *SimpleCSVWriter) writeAll(content [][]string, file *os.File) error {
	csvWriter := csv.NewWriter(file)
	csvWriter.Comma = simpleWriter.GetDelimiter()

	err := csvWriter.WriteAll(content)
	if nil != err {
		return err
	}

	csvWriter.Flush()
	return nil
}

func createFileIfNotExist(path string) (*os.File, error) {
	return os.OpenFile(path, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0o644)
}

func (simpleReader *SimpleCSVWriter) setDefaultDelimiterIfNil() {
	if simpleReader.delimiter == 0 {
		simpleReader.delimiter = ','
	}
}
