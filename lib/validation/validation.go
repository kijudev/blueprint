package validation

func ValidateField(field string, units ...*UnitError) *FieldError {
	error := new(FieldError)
	error.Field = field

	for _, unit := range units {
		if unit != nil {
			error.Errors = append(error.Errors, *unit)
		}
	}

	if len(error.Errors) == 0 {
		return nil
	}

	return error
}

func Combine(results ...*FieldError) *Error {
	error := new(Error)
	error.Data = make(map[string][]UnitError)

	for _, result := range results {
		if result != nil {
			error.Data[result.Field] = result.Errors
		}
	}

	return error
}
