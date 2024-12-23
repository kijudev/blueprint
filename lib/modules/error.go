package modules

import "errors"

const (
	ErrorCodeInitFailed        = "INIT_FAILED"
	ErrorCodeStopFailed        = "STOP_FAILED"
	ErrorCodeInvalidStatus     = "INVALID_STATUS"
	ErrorCodeUnknown           = "UNKNOWN"
	ErrorCodeMissingDependency = "MISSING_DEPS"
)

var (
	ErrorInitFailed        = errors.New(ErrorCodeInitFailed)
	ErrorStopFailed        = errors.New(ErrorCodeStopFailed)
	ErrorInvalidStatus     = errors.New(ErrorCodeInvalidStatus)
	ErrorUnknown           = errors.New(ErrorCodeUnknown)
	ErrorMissingDependency = errors.New(ErrorCodeMissingDependency)
)
