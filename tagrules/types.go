package tagrules

type TagRuler struct {
	rawRuleData        []map[string]string
	calculatedRuleData [][]string

	ruleColumnNames []string
	ruleResultName  string

	ruleOutputColumnName string
	ruleOutputResultName string

	rawRuleDataReader RawRuleDataReader
	ruleDataWriter    RuleDataWriter
}

type RawRuleDataReader interface {
	FetchRawData() ([]map[string]string, error)
}

type RuleDataWriter interface {
	StoreRuleData([][]string) error
}

var TagKeyName = "TagKeys"
var TagValueName = "TagValues"
