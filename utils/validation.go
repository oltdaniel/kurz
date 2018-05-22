package utils

// Load standard libraries
import "regexp"

// Set constant
var EmailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// Valiate email with regex
func ValidateEmail(email string) bool {
  // Check length
  if len(email) < 6 { return false }

  // Return regex result
  return EmailRegex.MatchString(email)
}

// Validate length with statements
func ValidateLength(inp string, min int, max int) bool {
  return (len(inp) >= min && len(inp) <= max)
}
