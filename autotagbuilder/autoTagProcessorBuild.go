package autotagbuilder

import (
	"github.com/ericlgf819/autotag/tagprocessor"
	"github.com/ericlgf819/autotag/tagprocessor/manifest"
	azuretagcsvreader "github.com/ericlgf819/autotag/tagprocessor/manifest/azuretag"
	"github.com/ericlgf819/autotag/tagprocessor/parser"
)

func BuildProcessorRuler() *tagprocessor.TagProcessor {
	processor := new(tagprocessor.TagProcessor)
	processor.Init(buildCsvReader(), buildCsvWriter(), buildTagParser())
	return processor
}

func buildTagReader() *azuretagcsvreader.AzureTagCSVDecoratorReader {
	tagReader := new(azuretagcsvreader.AzureTagCSVDecoratorReader)
	tagReader.Init(buildCsvReader())
	return tagReader
}

func buildManifestReader() *manifest.SimpleManifestReader {
	manifestReader := new(manifest.SimpleManifestReader)
	// TODO.. complete manifest file path
	manifestReader.Init("", buildTagReader())
	return manifestReader
}

func buildTagParser() *parser.AzureCSVParser {
	tagParser := new(parser.AzureCSVParser)
	tagParser.Init(buildManifestReader())
	return tagParser
}
