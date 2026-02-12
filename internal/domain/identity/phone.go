package identity

import (
	"fmt"
	"regexp"
	"strings"
)

var phoneRegexp = regexp.MustCompile(`^\+?[0-9]{7,15}$`)

// Phone is a value object representing a validated phone number.
type Phone struct {
	value string
}

// NewPhone creates a validated Phone value object. Empty string is allowed (optional field).
func NewPhone(raw string) (Phone, error) {
	trimmed := strings.TrimSpace(raw)
	if trimmed == "" {
		return Phone{value: ""}, nil
	}
	if !phoneRegexp.MatchString(trimmed) {
		return Phone{}, fmt.Errorf("invalid phone format: %s", trimmed)
	}
	return Phone{value: trimmed}, nil
}

// String returns the phone number as a string.
func (p Phone) String() string { return p.value }

// IsEmpty returns true if no phone was provided.
func (p Phone) IsEmpty() bool { return p.value == "" }

// Equals checks equality with another Phone.
func (p Phone) Equals(other Phone) bool { return p.value == other.value }
