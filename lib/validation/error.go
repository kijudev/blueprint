package validation

import "encoding/json"

type UnitError struct {
	Code string `json:"code"`
	Data string `json:"data"`
	Msg  string `json:"msg"`
}

type FieldError struct {
	Field  string      `json:"field"`
	Errors []UnitError `json:"errors"`
}

type Error struct {
	Data map[string][]UnitError `json:"data"`
}

var (
	CodeEmail    = "EMAIL"
	CodeEmpty    = "EMPTY"
	CodeTooLong  = "TOO_LONG"
	CodeTooShort = "TOO_SHORT"
	CodeMismatch = "MISMATCH"
)

func (unitError UnitError) Error() string {
	bytes, _ := json.Marshal(unitError)
	return string(bytes)
}

func (fieldError FieldError) Error() string {
	bytes, _ := json.Marshal(fieldError)
	return string(bytes)
}

func (error Error) Error() string {
	bytes, _ := json.Marshal(error)
	return string(bytes)
}
