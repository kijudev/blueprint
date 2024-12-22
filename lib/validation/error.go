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
	ErrorCodeEmpty    = "EMPTY"
	ErrorCodeTooShort = "TOO_SHORT"
	ErrorCodeTooLong  = "TOO_LONG"
	ErrorCodeTooBig   = "TOO_BIG"
	ErrorCodeTooSmall = "TOO_SMALL"
	ErrorCodeEmail    = "EMAIL"
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
