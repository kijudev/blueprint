package validation

import "encoding/json"

type RuleResult struct {
	Code string `json:"code"`
	Data string `json:"data"`
	Msg  string `json:"msg"`
}

type FieldResult struct {
	Field  string       `json:"field"`
	Errors []RuleResult `json:"errors"`
}

type ObjectResult struct {
	Data map[string][]RuleResult
}

var (
	CodeEmail    = "EMAIL"
	CodeEmpty    = "EMPTY"
	CodeTooLong  = "TOO_LONG"
	CodeTooShort = "TOO_SHORT"
	CodeMismatch = "MISMATCH"
)

func (ruleResult *RuleResult) JSON() string {
	bytes, _ := json.Marshal(ruleResult)
	return string(bytes)
}

func (fieldResult *FieldResult) JSON() string {
	bytes, _ := json.Marshal(fieldResult)
	return string(bytes)
}

func (objectResult *ObjectResult) JSON() string {
	bytes, _ := json.Marshal(objectResult.Data)
	return string(bytes)
}

func (ruleResult RuleResult) Error() string {
	bytes, _ := json.Marshal(ruleResult)
	return string(bytes)
}

func (fieldResult FieldResult) Error() string {
	bytes, _ := json.Marshal(fieldResult)
	return string(bytes)
}

func (objectResult ObjectResult) Error() string {
	bytes, _ := json.Marshal(objectResult.Data)
	return string(bytes)
}

func NewField(field string) *FieldResult {
	return &FieldResult{
		Field:  field,
		Errors: []RuleResult{},
	}
}

func (fieldResult *FieldResult) AddRule(ruleResult *RuleResult) {
	if ruleResult == nil {
		return
	}

	fieldResult.Errors = append(fieldResult.Errors, *ruleResult)
}

func NewObject() ObjectResult {
	return ObjectResult{
		Data: make(map[string][]RuleResult),
	}
}

func (objectResult *ObjectResult) AddField(fieldResult *FieldResult) {
	if fieldResult == nil {
		return
	}

	if len(fieldResult.Errors) == 0 {
		return
	}

	objectResult.Data[fieldResult.Field] = fieldResult.Errors
}

func (objectResult *ObjectResult) Resolve() bool {
	if len(objectResult.Data) == 0 {
		return true
	}

	return false
}
