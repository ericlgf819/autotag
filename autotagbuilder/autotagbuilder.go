package autotagbuilder

import (
	"github.com/ericlgf819/autotag/tagprocessor/utils/csvreader"
	"github.com/ericlgf819/autotag/tagprocessor/utils/csvwriter"
	"github.com/ericlgf819/autotag/tagrules"
	"github.com/ericlgf819/autotag/tagrules/simplerawruledatareader"
	"github.com/ericlgf819/autotag/tagrules/simpleruledatawriter"
)

func BuildTagRuler() *tagrules.TagRuler {

	tagRuler := new(tagrules.TagRuler)
	// TODO.. modify rules
	tagRuler.InitRulesConditionsByDefault()
	tagRuler.Init(buildRawRuleDataReader(), buildRuleDataWriter())
	return tagRuler
}

func buildRawRuleDataReader() *simplerawruledatareader.SimpleRawRuleDataReader {
	dataReader := new(simplerawruledatareader.SimpleRawRuleDataReader)
	dataReader.Init(buildCsvReader())
	// TODO..
	dataReader.LoadFile("./rawruledata.csv")
	return dataReader
}

func buildCsvReader() *csvreader.SimpleCSVReader {
	csvReader := new(csvreader.SimpleCSVReader)
	csvReader.SetDelimiter(',')
	return csvReader
}

func buildRuleDataWriter() *simpleruledatawriter.SimpleRuleDataWriter {
	dataWriter := new(simpleruledatawriter.SimpleRuleDataWriter)
	// TODO for the path..
	dataWriter.Init("./ruledataoutput.csv", buildCsvWriter())
	return dataWriter
}

func buildCsvWriter() *csvwriter.SimpleCSVWriter {
	csvWriter := new(csvwriter.SimpleCSVWriter)
	csvWriter.SetDelimiter(',')
	return csvWriter
}
