package tagrules

// TODO...
func (tagRuler *TagRuler) InitRulesByDefault() {
	tagRuler.ruleColumnNames = []string{
		"col1",
		"col2",
		"col3",
	}

	tagRuler.ruleResultName = "col4"

	tagRuler.ruleOutputColumnName = "outputColumn"
	tagRuler.ruleOutputResultName = "outputValue"
}

func (tagRuler *TagRuler) InitRules(colNames []string, tagetColName string, outputColumnName string, outputResultName string) {
	tagRuler.ruleColumnNames = colNames
	tagRuler.ruleResultName = tagetColName
	tagRuler.ruleOutputColumnName = outputColumnName
	tagRuler.ruleOutputResultName = outputResultName
}

func (tagRuler *TagRuler) Init(reader RawRuleDataReader, writer RuleDataWriter) {
	tagRuler.rawRuleDataReader = reader
	tagRuler.ruleDataWriter = writer
}

func (tagRuler *TagRuler) MakeRules() error {
	err := tagRuler.loadRawData()
	if err != nil {
		return err
	}

	err = tagRuler.makeRulesInternally()
	return err
}

func (tagRuler *TagRuler) loadRawData() error {
	if tagRuler.rawRuleData != nil {
		return nil
	}

	var err error
	tagRuler.rawRuleData, err = tagRuler.rawRuleDataReader.FetchRawData()

	return err
}

func (tagRuler *TagRuler) makeRulesInternally() error {
	return nil
}
