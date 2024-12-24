package services

import "errors"

const (
	ErrCodeNotFound         = "NOT_FOUND"
	ErrCodeDependencyFailed = "DEPENDENCY_FAILED"
	ErrCodeDatasourceFailed = "DATASOURCE_FAILED"
	ErrCodeInvalidParams    = "INVALID_PARAMS"
	ErrCodeValidationFailed = "VALIDATION_FAILED"
	ErrCodeUnknown          = "UNKNOWN"
)

var (
	ErrNotFound         = errors.New(ErrCodeNotFound)
	ErrDependencyFailed = errors.New(ErrCodeDependencyFailed)
	ErrDatasourceFailed = errors.New(ErrCodeDatasourceFailed)
	ErrInvalidParams    = errors.New(ErrCodeInvalidParams)
	ErrValidationFailed = errors.New(ErrCodeValidationFailed)
	ErrUnknown          = errors.New(ErrCodeUnknown)
)
