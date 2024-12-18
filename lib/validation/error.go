package validation

type UnitError struct {
	Code string `json:"code"`
	Data string `json:"data"`
	Msg  string `json:"msg"`
}

type FieldError struct {
	Field  string      `json:"field"`
	Errors []UnitError `json:"errors"`
}

type Error map[string]UnitError

var (
	CodeEmail    = "EMAIL"
	CodeEmpty    = "EMPTY"
	CodeTooLong  = "TOO_LONG"
	CodeTooShort = "TOO_SHORT"
	CodeMismatch = "MISMATCH"
)
