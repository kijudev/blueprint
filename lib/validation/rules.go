package validation

import (
	"regexp"
	"strings"
)

type Rule func(s string, args ...any) *UnitError

func Empty(s string) *UnitError {
	if s == "" {
		return &UnitError{
			Code: CodeEmpty,
		}
	}

	return nil
}

func MaxLength(s string, length int) *UnitError {
	if len(s) > length {
		return &UnitError{
			Code: CodeTooLong,
		}
	}

	return nil
}

func MinLength(s string, length int) *UnitError {
	if len(s) < length {
		return &UnitError{
			Code: CodeTooShort,
		}
	}

	return nil
}

func Email(s string) *UnitError {
	if s == "" {
		return &UnitError{
			Code: CodeEmail,
		}
	}

	if len(s) > 254 {
		return &UnitError{
			Code: CodeEmail,
		}
	}

	parts := strings.Split(s, "@")

	if len(parts) != 2 {

		return &UnitError{
			Code: CodeEmail,
		}
	}

	localPart, domain := parts[0], parts[1]

	if len(localPart) > 64 {
		return &UnitError{
			Code: CodeEmail,
		}
	}

	if len(domain) > 253 {
		return &UnitError{
			Code: CodeEmail,
		}
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	if !emailRegex.MatchString(s) {
		return &UnitError{
			Code: CodeEmail,
		}
	}

	return nil
}
