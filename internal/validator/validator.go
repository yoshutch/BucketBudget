package validator

import (
	"strings"
	"unicode/utf8"
)

type Validator struct {
	FieldErrors map[string][]string
}

func (v *Validator) Valid() bool {
	return len(v.FieldErrors) == 0
}

func (v *Validator) CheckField(ok bool, key, message string) {
	if v.FieldErrors == nil {
		v.FieldErrors = make(map[string][]string)
	}
	if !ok {
		FieldErrors, exists := v.FieldErrors[key]
		if !exists {
			FieldErrors = []string{}
		}
		v.FieldErrors[key] = append(FieldErrors, message)
	}
}

func IsNotBlank(val string) bool {
	return strings.TrimSpace(val) != ""
}

func MaxChars(value string, n int) bool {
	return utf8.RuneCountInString(value) <= n
}
