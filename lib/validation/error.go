package validation

import "encoding/json"

type RuleError struct {
	Code string `json:"code"`
	Data string `json:"data"`
	Msg  string `json:"msg"`
}

type ValidationError []RuleError

type Error map[string]ValidationError

var (
	CodeEmpty    = "EMPTY"
	CodeTooShort = "TOO_SHORT"
	CodeTooLong  = "TOO_LONG"
	CodeTooBig   = "TOO_BIG"
	CodeTooSmall = "TOO_SMALL"
	CodeEmail    = "EMAIL"
)

func (r RuleError) Error() string {
	bytes, _ := json.Marshal(r)
	return string(bytes)
}

func (v ValidationError) Error() string {
	bytes, _ := json.Marshal(v)
	return string(bytes)
}

func (c Error) Error() string {
	bytes, _ := json.Marshal(c)
	return string(bytes)
}
