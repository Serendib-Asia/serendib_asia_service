package validator

import (
	"regexp"
	"slices"
	"strings"
)

// Validate the filed value is valid for expected regexp
func containsOnly(field, regExpStr string) bool {
	reg, _ := regexp.Compile(regExpStr)
	return !reg.MatchString(field)
}

// Validate the field value is one of the valid values
func validateOneOf(field string, validValues []string) bool {
	return slices.Contains(validValues, field)
}

// Validate the field value is a non-empty string
func trimAndCheckLength(field string, length int) bool {
	return len(strings.TrimSpace(field)) > length
}
