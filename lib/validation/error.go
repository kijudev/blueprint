package validation

import "encoding/json"

type RuleError struct {
	Code string `json:"code"`
	Data string `json:"data"`
	Msg  string `json:"msg"`
}

type ValidatorError []RuleError

var (
	CodeEmpty    = "EMPTY"
	CodeTooShort = "TOO_SHORT"
	CodeTooLong  = "TOO_LONG"
	CodeTooBig   = "TOO_BIG"
	CodeTooSmall = "TOO_SMALL"
	CodeEmail    = "EMAIL"
)

func (r *RuleError) Error() string {
	bytes, _ := json.Marshal(r)
	return string(bytes)
}

func (v *ValidatorError) Error() string {
	bytes, _ := json.Marshal(v)
	return string(bytes)
}
