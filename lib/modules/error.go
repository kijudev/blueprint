package modules

import "errors"

const (
	ErrorCodeInitFailed    = "INIT_FAILED"
	ErrorCodeStopFailed    = "STOP_FAILED"
	ErrorCodeInvalidStatus = "INVALID_STATUS"
	ErrorCodeUnknown       = "UNKNOWN"
)

var (
	ErrorInitFailed    = errors.New(ErrorCodeInitFailed)
	ErrorStopFailed    = errors.New(ErrorCodeStopFailed)
	ErrorInvalidStatus = errors.New(ErrorCodeInvalidStatus)
	ErrorUnknown       = errors.New(ErrorCodeUnknown)
)
