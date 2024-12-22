package modules

import "errors"

const (
	ErrorCodeInitFailed    = "INIT_FAILED"
	ErrorCodeStopFailed    = "STOP_FAILED"
	ErrorCodeInvalidStatus = "INVALID_STATUS"
)

var (
	ErrorInitFailed    = errors.New(ErrorCodeInitFailed)
	ErrorStopFailed    = errors.New(ErrorCodeStopFailed)
	ErrorInvalidStatus = errors.New(ErrorCodeInvalidStatus)
)
