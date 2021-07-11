package tagprocessor

import "errors"

func (processor *TagProcessor) Init(reader ContentReader, writer ContentWriter, parser TagParser) {
	processor.reader = reader
	processor.writer = writer
	processor.parser = parser
}

func (processor *TagProcessor) Process(inputFilePath string, outputFilePath string) error {
	return errors.New("not implemented")
}
