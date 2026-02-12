package identity

import (
	"fmt"
	"regexp"
	"strings"
)

var emailRegexp = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

// Email is a value object representing a validated email address.
type Email struct {
	value string
}

// NewEmail creates a validated Email value object.
func NewEmail(raw string) (Email, error) {
	trimmed := strings.TrimSpace(raw)
	if trimmed == "" {
		return Email{}, fmt.Errorf("email is required")
	}
	if !emailRegexp.MatchString(trimmed) {
		return Email{}, fmt.Errorf("invalid email format: %s", trimmed)
	}
	return Email{value: strings.ToLower(trimmed)}, nil
}

// String returns the email as a string.
func (e Email) String() string { return e.value }

// Equals checks equality with another Email.
func (e Email) Equals(other Email) bool { return e.value == other.value }
