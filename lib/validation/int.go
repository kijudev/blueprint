package validation

type IntValidator struct {
	data int
	err  ValidationError
}

func Int(data int) *IntValidator {
	return &IntValidator{
		data: data,
		err:  ValidationError{},
	}
}

func (v *IntValidator) Validate() error {
	if len(v.err) > 0 {
		return &v.err
	}

	return nil
}

func (v *IntValidator) getError() ValidationError {
	return v.err
}

func (v *IntValidator) MinValue(min int) *IntValidator {
	valid, err := RuleMinValueInt(v.data, min)

	if !valid {
		v.err = append(v.err, *err)
	}

	return v
}

func (v *IntValidator) MaxValue(max int) *IntValidator {
	valid, err := RuleMaxValueInt(v.data, max)

	if !valid {
		v.err = append(v.err, *err)
	}

	return v
}
