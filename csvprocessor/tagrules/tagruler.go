package tagrules

import "errors"

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

func (tagRuler *TagRuler) StoreRules() error {
	if tagRuler.calculatedRuleData == nil {
		return errors.New("call MakeRules() before StoreRules()")
	}

	return tagRuler.ruleDataWriter.StoreRuleData(tagRuler.calculatedRuleData)
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
	tagRuler.makeOutputHeader()

	for _, dic := range tagRuler.rawRuleData {
		if validateRowMapItem(dic, tagRuler.ruleColumnNames, tagRuler.ruleResultName) {
			tagRuler.makeOutputRow(dic)
		}
	}
	return nil
}

func (tagRuler *TagRuler) makeOutputHeader() {
	tagRuler.calculatedRuleData =
		append(
			[][]string{},
			[]string{tagRuler.ruleOutputColumnName, tagRuler.ruleOutputResultName})
}

func (tagRuler *TagRuler) makeOutputRow(rawDic map[string]string) {
	tagKey := ""
	for _, colName := range tagRuler.ruleColumnNames {
		if tagKey == "" {
			tagKey = rawDic[colName]
		} else {
			tagKey = tagKey + "|" + rawDic[colName]
		}
	}
	tagValue := rawDic[tagRuler.ruleResultName]

	tagRuler.calculatedRuleData =
		append(tagRuler.calculatedRuleData,
			[]string{tagKey, tagValue})
}

func validateRowMapItem(dic map[string]string, ruleColumnNames []string, ruleResultName string) bool {
	for _, colName := range ruleColumnNames {
		if dic[colName] == "" {
			return false
		}
	}

	return dic[ruleResultName] != ""
}
