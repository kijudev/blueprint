package validation

type StringValidator struct {
	data string
	err  ValidatorError
}

func String(data string) *StringValidator {
	return &StringValidator{
		data: data,
		err:  ValidatorError{},
	}
}

func (v *StringValidator) Validate() error {
	if len(v.err) > 0 {
		return &v.err
	}

	return nil
}

func (v *StringValidator) NotEmpty() *StringValidator {
	valid, err := RuleNotEmpty(v.data)

	if !valid {
		v.err = append(v.err, *err)
	}

	return v
}

func (v *StringValidator) MinLength(min int) *StringValidator {
	valid, err := RuleMinLength(v.data, min)

	if !valid {
		v.err = append(v.err, *err)
	}

	return v
}

func (v *StringValidator) MaxLength(max int) *StringValidator {
	valid, err := RuleMaxLength(v.data, max)

	if !valid {
		v.err = append(v.err, *err)
	}

	return v
}

func (v *StringValidator) Email() *StringValidator {
	valid, err := RuleEmail(v.data)

	if !valid {
		v.err = append(v.err, *err)
	}

	return v
}
