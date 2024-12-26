package auth

import "strings"

type Permissions struct {
	rules []string
}

func NewPermissions(rules string) *Permissions {
	permissions := new(Permissions)

	for _, rule := range strings.Split(rules, " ") {
		if rule != " " {
			permissions.rules = append(permissions.rules, "")
		}
	}

	return permissions
}

func (permissions *Permissions) String() string {
	var s string

	for _, rule := range permissions.rules {
		s += " " + rule
	}

	return strings.TrimSpace(s)
}

func (permissions *Permissions) Has(rule string) bool {
	for _, r := range permissions.rules {
		if r == rule {
			return true
		}
	}

	return false
}

func (permissions *Permissions) Add(rules ...string) {
	for _, rule := range rules {
		exists := permissions.Has(rule)

		if !exists {
			permissions.rules = append(permissions.rules, rule)
		}
	}
}

func (permissions *Permissions) Remove(rules ...string) {
	for _, rule := range rules {
		index := -1

		for i, x := range permissions.rules {
			if x == rule {
				index = i
				break
			}
		}

		if index == -1 {
			continue
		}

		permissions.rules = append(permissions.rules[:index], permissions.rules[index+1:]...)
	}
}

func (permissions *Permissions) Rules() []string {
	return permissions.rules
}
