package csvreader

import (
	"encoding/csv"
	"os"
)

func (simpleReader *SimpleCSVReader) SetDelimiter(delimeter int32) {
	simpleReader.delimiter = delimeter
}

func (simpleReader *SimpleCSVReader) GetDelimiter() int32 {
	return simpleReader.delimiter
}

func (simpleReader *SimpleCSVReader) ReadFile(path string) ([][]string, error) {
	simpleReader.setDefaultDelimiterIfNil()

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	csvReader.Comma = simpleReader.GetDelimiter()

	return csvReader.ReadAll()
}

func (simpleReader *SimpleCSVReader) setDefaultDelimiterIfNil() {
	if simpleReader.delimiter == 0 {
		simpleReader.delimiter = ','
	}
}
