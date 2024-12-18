package validation

import (
	"regexp"
	"strings"
)

func NotEmpty(s string) *RuleResult {
	if s == "" {
		return &RuleResult{
			Code: CodeEmpty,
		}
	}

	return nil
}

func MaxLength(s string, length int) *RuleResult {
	if len(s) > length {
		return &RuleResult{
			Code: CodeTooLong,
		}
	}

	return nil
}

func MinLength(s string, length int) *RuleResult {
	if len(s) < length {
		return &RuleResult{
			Code: CodeTooShort,
		}
	}

	return nil
}

func Email(s string) *RuleResult {
	if s == "" {
		return &RuleResult{
			Code: CodeEmail,
		}
	}

	if len(s) > 254 {
		return &RuleResult{
			Code: CodeEmail,
		}
	}

	parts := strings.Split(s, "@")

	if len(parts) != 2 {
		return &RuleResult{
			Code: CodeEmail,
		}
	}

	localPart, domain := parts[0], parts[1]

	if len(localPart) > 64 {
		return &RuleResult{
			Code: CodeEmail,
		}
	}

	if len(domain) > 253 {
		return &RuleResult{
			Code: CodeEmail,
		}
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	if !emailRegex.MatchString(s) {
		return &RuleResult{
			Code: CodeEmail,
		}
	} else {

		return nil
	}
}
