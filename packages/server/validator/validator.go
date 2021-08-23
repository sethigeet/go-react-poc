// Package validator provides functions to validate a struct according to the
// tags specified on it
package validator

import (
	"reflect"
	"strings"
)

// StructTagKey is the name of the key that has to be used while specifying the
// validation rules in the tag on the struct field
const StructTagKey = "validation"

// Validate validates the properties of a struct according to the validation
// rules specified in the tags of the struct fields
func Validate(input interface{}) map[string]string {
	var errors = map[string]string{}

	ele := reflect.ValueOf(input)
	typ := ele.Type()
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		tag := field.Tag
		flags, ok := tag.Lookup(StructTagKey)
		if !ok {
			continue
		}

		for _, flag := range strings.Split(flags, ",") {
			// Lowercase the first letter of the key to match with incoming JSON
			fieldName := strings.ToLower(string(field.Name[0])) + string(field.Name[1:])
			if errors[fieldName] == "" {
				if err := validateField(ele.Field(i), fieldName, flag); err != "" {
					errors[fieldName] = err
				}
			}
		}
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

func validateField(val reflect.Value, fieldName, flag string) string {
	switch val.Type().Kind() {
	case reflect.String:
		return validateString(val.String(), fieldName, flag)
	default:
		return ""
	}
}
