package csvreader

import (
	"os"
)

func (simpleReader *SimpleCSVReader) SetDelimiter(delimeter string) {
	simpleReader.delimiter = delimeter
}

func (simpleReader *SimpleCSVReader) GetDelimiter() string {
	return simpleReader.delimiter
}

func (simpleReader *SimpleCSVReader) ReadFile(path string) ([][]string, error) {
	simpleReader.setDefaultDelimiterIfNil()

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	//csvReader := csv.NewReader(file)

	return nil, nil
}

func (simpleReader *SimpleCSVReader) setDefaultDelimiterIfNil() {
	if simpleReader.delimiter == "" {
		simpleReader.delimiter = ","
	}
}
