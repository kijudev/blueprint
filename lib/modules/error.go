package modules

import "errors"

const (
	ErrCodeInitFailed        = "INIT_FAILED"
	ErrCodeStopFailed        = "STOP_FAILED"
	ErrCodeInvalidStatus     = "INVALID_STATUS"
	ErrCodeUnknown           = "UNKNOWN"
	ErrCodeMissingDependency = "MISSING_DEPENDENCY"
)

var (
	ErrInitFailed        = errors.New(ErrCodeInitFailed)
	ErrStopFailed        = errors.New(ErrCodeStopFailed)
	ErrInvalidStatus     = errors.New(ErrCodeInvalidStatus)
	ErrUnknown           = errors.New(ErrCodeUnknown)
	ErrMissingDependency = errors.New(ErrCodeMissingDependency)
)
