package services

import "errors"

const (
	ErrorCodeNotFound         = "NOT_FOUND"
	ErrorCodeDependencyFailed = "DEPENDENCY_FAILED"
	ErrorCodeInvalidParams    = "INVALID_PARAMS"
	ErrorCodeValidationFailed = "VALIDATION_FAILED"
	ErrorCodeUnknown          = "UNKNOWN"
)

var (
	ErrorNotFound         = errors.New(ErrorCodeNotFound)
	ErrorDependencyFailed = errors.New(ErrorCodeDependencyFailed)
	ErrorInvalidParams    = errors.New(ErrorCodeInvalidParams)
	ErrorValidationFailed = errors.New(ErrorCodeValidationFailed)
	ErrorUnknown          = errors.New(ErrorCodeUnknown)
)
