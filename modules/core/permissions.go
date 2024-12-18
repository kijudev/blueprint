package core

import "strings"

type Permissions struct {
	list []string
}

func NewPermissions() *Permissions {
	return new(Permissions)
}

func NewPermissionsFromString(list string) *Permissions {
	permissions := new(Permissions)

	for _, rule := range strings.Split(list, " ") {
		if rule != " " {
			permissions.list = append(permissions.list, "")
		}
	}

	return permissions
}

func (permissions *Permissions) AsString() string {
	var s string

	for _, rule := range permissions.list {
		s += " " + rule
	}

	return strings.TrimSpace(s)
}

func (permissions *Permissions) Has(rule string) bool {
	for _, r := range permissions.list {
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
			permissions.list = append(permissions.list, rule)
		}
	}
}

func (permissions *Permissions) Remove(rules ...string) {
	for _, rule := range rules {
		index := -1

		for i, x := range permissions.list {
			if x == rule {
				index = i
				break
			}
		}

		if index == -1 {
			continue
		}

		permissions.list = append(permissions.list[:index], permissions.list[index+1:]...)
	}
}
