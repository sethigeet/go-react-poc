package validator

import (
	"fmt"
	"regexp"
)

func validateString(val, fieldName, flag string) string {
	switch flag {
	case "required":
		if val == "" {
			return fmt.Sprintf(ErrorMessages["required"], fieldName)
		}
	case "email":
		emailPattern := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
		matches := emailPattern.MatchString(val)
		if !matches {
			return fmt.Sprintf(ErrorMessages["invalidEmail"], fieldName)
		}
	default:
		panic("Incorrect flag provided for validation!")
	}

	return ""
}
