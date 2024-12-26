package lib

import (
	"strconv"
	"strings"
)

type Validator interface {
	Validate() error
	getError() ValFieldError
}

type ValCollection struct {
	err ValCollectionError
}

func NewValCollection() *ValCollection {
	return &ValCollection{
		err: make(ValCollectionError),
	}
}

func (c *ValCollection) Add(key string, validator Validator) {
	err := validator.getError()

	if len(err) > 0 {
		c.err[key] = err
	}
}

func (c *ValCollection) Resolve() error {
	if len(c.err) == 0 {
		return nil
	}

	return c.err
}

type StringValidator struct {
	data string
	err  ValFieldError
}

func ValString(data string) *StringValidator {
	return &StringValidator{
		data: data,
		err:  ValFieldError{},
	}
}

func (v *StringValidator) Validate() error {
	if len(v.err) > 0 {
		return &v.err
	}

	return nil
}

func (v *StringValidator) getError() ValFieldError {
	return v.err
}

func (v *StringValidator) NotEmpty() *StringValidator {
	valid, err := ValRuleNotEmpty(v.data)

	if !valid {
		v.err = append(v.err, *err)
	}

	return v
}

func (v *StringValidator) MinLength(min int) *StringValidator {
	valid, err := ValRuleMinLength(v.data, min)

	if !valid {
		v.err = append(v.err, *err)
	}

	return v
}

func (v *StringValidator) MaxLength(max int) *StringValidator {
	valid, err := ValRuleMaxLength(v.data, max)

	if !valid {
		v.err = append(v.err, *err)
	}

	return v
}

func (v *StringValidator) Email() *StringValidator {
	valid, err := ValRuleEmail(v.data)

	if !valid {
		v.err = append(v.err, *err)
	}

	return v
}

type IntValidator struct {
	data int
	err  ValFieldError
}

func ValInt(data int) *IntValidator {
	return &IntValidator{
		data: data,
		err:  ValFieldError{},
	}
}

func (v *IntValidator) Validate() error {
	if len(v.err) > 0 {
		return &v.err
	}

	return nil
}

func (v *IntValidator) getError() ValFieldError {
	return v.err
}

func (v *IntValidator) MinValue(min int) *IntValidator {
	valid, err := ValRuleMinValueInt(v.data, min)

	if !valid {
		v.err = append(v.err, *err)
	}

	return v
}

func (v *IntValidator) MaxValue(max int) *IntValidator {
	valid, err := ValRuleMaxValueInt(v.data, max)

	if !valid {
		v.err = append(v.err, *err)
	}

	return v
}

type FloatValidator struct {
	data float64
	err  ValFieldError
}

func ValFloat(data float64) *FloatValidator {
	return &FloatValidator{
		data: data,
		err:  ValFieldError{},
	}
}

func (v *FloatValidator) Validate() error {
	if len(v.err) > 0 {
		return &v.err
	}

	return nil
}

func (v *FloatValidator) getError() ValFieldError {
	return v.err
}

func (v *FloatValidator) MinValue(min float64) *FloatValidator {
	valid, err := ValRuleMinValueFloat(v.data, min)

	if !valid {
		v.err = append(v.err, *err)
	}

	return v
}

func (v *FloatValidator) MaxValue(max float64) *FloatValidator {
	valid, err := ValRuleMaxValueFloat(v.data, max)

	if !valid {
		v.err = append(v.err, *err)
	}

	return v
}

func ValRuleNotEmpty(data string) (bool, *ValRuleError) {
	if data == "" {
		return false, &ValRuleError{Code: ErrorCodeEmpty}
	}
	return true, nil
}

func ValRuleMinLength(data string, min int) (bool, *ValRuleError) {
	if len(data) < min {
		return false, &ValRuleError{Code: ErrorCodeTooShort, Data: strconv.Itoa(min)}
	}
	return true, nil
}

func ValRuleMaxLength(data string, max int) (bool, *ValRuleError) {
	if len(data) > max {
		return false, &ValRuleError{Code: ErrorCodeTooLong, Data: strconv.Itoa(max)}
	}
	return true, nil
}

func ValRuleMinValueInt(data int, min int) (bool, *ValRuleError) {
	if data < min {
		return false, &ValRuleError{Code: ErrorCodeTooSmall, Data: strconv.Itoa(min)}
	}
	return true, nil
}

func ValRuleMaxValueInt(data int, max int) (bool, *ValRuleError) {
	if data > max {
		return false, &ValRuleError{Code: ErrorCodeTooBig, Data: strconv.Itoa(max)}
	}
	return true, nil
}

func ValRuleMinValueFloat(data float64, min float64) (bool, *ValRuleError) {
	if data < min {
		return false, &ValRuleError{Code: ErrorCodeTooSmall, Data: strconv.FormatFloat(min, 'f', -1, 64)}
	}
	return true, nil
}

func ValRuleMaxValueFloat(data float64, max float64) (bool, *ValRuleError) {
	if data > max {
		return false, &ValRuleError{Code: ErrorCodeTooBig, Data: strconv.FormatFloat(max, 'f', -1, 64)}
	}
	return true, nil
}

func ValRuleEmail(data string) (bool, *ValRuleError) {
	if data == "" {
		return false, &ValRuleError{Code: ErrorCodeEmail}
	}
	if len(data) < 3 {
		return false, &ValRuleError{Code: ErrorCodeEmail, Data: "3"}
	}
	if len(data) > 254 {
		return false, &ValRuleError{Code: ErrorCodeEmail, Data: "254"}
	}

	if !strings.Contains(data, "@") {
		return false, &ValRuleError{Code: ErrorCodeEmail, Data: "@"}
	}

	return true, nil
}
