package validation

type FloatValidator struct {
	data float64
	err  ValidationError
}

func Float(data float64) *FloatValidator {
	return &FloatValidator{
		data: data,
		err:  ValidationError{},
	}
}

func (v *FloatValidator) Validate() error {
	if len(v.err) > 0 {
		return &v.err
	}

	return nil
}

func (v *FloatValidator) getError() ValidationError {
	return v.err
}

func (v *FloatValidator) MinValue(min float64) *FloatValidator {
	valid, err := RuleMinValueFloat(v.data, min)

	if !valid {
		v.err = append(v.err, *err)
	}

	return v
}

func (v *FloatValidator) MaxValue(max float64) *FloatValidator {
	valid, err := RuleMaxValueFloat(v.data, max)

	if !valid {
		v.err = append(v.err, *err)
	}

	return v
}
