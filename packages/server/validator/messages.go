package validator

// ErrorMessages is a map of the error type/name and the message that has to be
// given when that validation error occurrs
var ErrorMessages = map[string]string{
	// Type          Message                             Format Specifiers
	"required":      "%s is required",                   // field
	"invalidEmail":  "%s must be a valid email address", // field
	"contains":      "%s must not contain a(n) '%s'",    // field, character
	"alreadyExists": "A %s with that %s already exists", // object, field
}
