package lib

import (
	"encoding/json"
	"errors"
	"unsafe"
)

type ValRuleError struct {
	Code string `json:"code"`
	Data string `json:"data"`
	Msg  string `json:"msg"`
}

type ValFieldError []ValRuleError

type ValCollectionError map[string]ValFieldError

var (
	ErrorCodeEmpty    = "EMPTY"
	ErrorCodeTooShort = "TOO_SHORT"
	ErrorCodeTooLong  = "TOO_LONG"
	ErrorCodeTooBig   = "TOO_BIG"
	ErrorCodeTooSmall = "TOO_SMALL"
	ErrorCodeEmail    = "EMAIL"
)

func (r ValRuleError) Error() string {
	bytes, _ := json.Marshal(r)
	return string(bytes)
}

func (v ValFieldError) Error() string {
	bytes, _ := json.Marshal(v)
	return string(bytes)
}

func (c ValCollectionError) Error() string {
	bytes, _ := json.Marshal(c)
	return string(bytes)
}

const (
	ErrCodeNotFound             = "NOT_FOUND"
	ErrCodeDependencyFailed     = "DEPENDENCY_FAILED"
	ErrCodeValidationFailed     = "VALIDATION_FAILED"
	ErrCodeDatasourceFailed     = "DATASOURCE_FAILED"
	ErrCodeAuthorizationFailed  = "AUTHORIZATION_FAILED"
	ErrCodeAuthenticationFailed = "AUTHENTICATION_FAILED"
	ErrCodeModuleInitFailed     = "MODULE_INIT_FAILED"
	ErrCodeModuleStopFailed     = "MODULE_STOP_FAILED"
	ErrCodeModuleNotInitialized = "MODULE_NOT_INITIALIZED"
	ErrCodeModuleAlreadyRunning = "MODULE_ALREADY_RUNNING"
	ErrCodeMissingDependency    = "MISSING_DEPENDENCY"
	ErrCodeUnknown              = "UNKNOWN"
)

var (
	ErrNotFound             = errors.New(ErrCodeNotFound)
	ErrDependencyFailed     = errors.New(ErrCodeDependencyFailed)
	ErrValidationFailed     = errors.New(ErrCodeValidationFailed)
	ErrDatasourceFailed     = errors.New(ErrCodeDatasourceFailed)
	ErrAuthorizationFailed  = errors.New(ErrCodeAuthorizationFailed)
	ErrAuthenticationFailed = errors.New(ErrCodeAuthenticationFailed)
	ErrModuleInitFailed     = errors.New(ErrCodeModuleInitFailed)
	ErrModuleStopFailed     = errors.New(ErrCodeModuleStopFailed)
	ErrModuleNotInitialized = errors.New(ErrCodeModuleNotInitialized)
	ErrModuleAlreadyRunning = errors.New(ErrCodeModuleAlreadyRunning)
	ErrMissingDependency    = errors.New(ErrCodeMissingDependency)
	ErrUnknown              = errors.New(ErrCodeUnknown)
)

func JoinErrors(errs ...error) error {
	n := 0

	for _, err := range errs {
		if err != nil {
			n++
		}
	}

	if n == 0 {
		return nil
	}

	e := &joinError{
		errs: make([]error, 0, n),
	}

	for _, err := range errs {
		if err != nil {
			e.errs = append(e.errs, err)
		}
	}
	return e
}

type joinError struct {
	errs []error
}

func (e *joinError) Error() string {
	// Since Join returns nil if every value in errs is nil,
	// e.errs cannot be empty.
	if len(e.errs) == 1 {
		return e.errs[0].Error()
	}

	b := []byte(e.errs[0].Error())

	for _, err := range e.errs[1:] {
		b = append(b, ';')
		b = append(b, ' ')
		b = append(b, err.Error()...)
	}

	// At this point, b has at least one byte '\n'.
	return "=> " + unsafe.String(&b[0], len(b))
}

func (e *joinError) Unwrap() []error {
	return e.errs
}
