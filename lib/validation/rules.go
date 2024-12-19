package validation

import (
	"strconv"
	"strings"
)

func RuleNotEmpty(data string) (bool, *RuleError) {
	if data == "" {
		return false, &RuleError{Code: CodeEmpty}
	}
	return true, nil
}

func RuleMinLength(data string, min int) (bool, *RuleError) {
	if len(data) < min {
		return false, &RuleError{Code: CodeTooShort, Data: strconv.Itoa(min)}
	}
	return true, nil
}

func RuleMaxLength(data string, max int) (bool, *RuleError) {
	if len(data) > max {
		return false, &RuleError{Code: CodeTooLong, Data: strconv.Itoa(max)}
	}
	return true, nil
}

func RuleMinValueInt(data int, min int) (bool, *RuleError) {
	if data < min {
		return false, &RuleError{Code: CodeTooSmall, Data: strconv.Itoa(min)}
	}
	return true, nil
}

func RuleMaxValueInt(data int, max int) (bool, *RuleError) {
	if data > max {
		return false, &RuleError{Code: CodeTooBig, Data: strconv.Itoa(max)}
	}
	return true, nil
}

func RuleMinValueFloat(data float64, min float64) (bool, *RuleError) {
	if data < min {
		return false, &RuleError{Code: CodeTooSmall, Data: strconv.FormatFloat(min, 'f', -1, 64)}
	}
	return true, nil
}

func RuleMaxValueFloat(data float64, max float64) (bool, *RuleError) {
	if data > max {
		return false, &RuleError{Code: CodeTooBig, Data: strconv.FormatFloat(max, 'f', -1, 64)}
	}
	return true, nil
}

func RuleEmail(data string) (bool, *RuleError) {
	if data == "" {
		return false, &RuleError{Code: CodeEmail}
	}
	if len(data) < 3 {
		return false, &RuleError{Code: CodeEmail, Data: "3"}
	}
	if len(data) > 254 {
		return false, &RuleError{Code: CodeEmail, Data: "254"}
	}

	if !strings.Contains(data, "@") {
		return false, &RuleError{Code: CodeEmail, Data: "@"}
	}

	return true, nil
}
