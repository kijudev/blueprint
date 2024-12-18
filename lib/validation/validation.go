package validation

func ValidateField(field string, units ...*UnitError) *FieldError {
	error := new(FieldError)

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
