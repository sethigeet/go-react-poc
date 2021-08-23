package validator_test

import (
	"fmt"
	"testing"

	"github.com/sethigeet/go-react-poc/packages/server/model"
	"github.com/sethigeet/go-react-poc/packages/server/validator"
)

type validateTest struct {
	input    interface{}
	expected map[string]string
}

var validateTests = []validateTest{
	// Required fields
	{
		input: model.User{
			Username: "",
			Email:    "",
		},
		expected: map[string]string{
			"username": fmt.Sprintf(validator.ErrorMessages["required"], "username"),
			"email":    fmt.Sprintf(validator.ErrorMessages["required"], "email"),
		},
	},

	// Email
	{
		input: model.User{
			Username: "someuser",
			Email:    "abcd",
		},
		expected: map[string]string{
			"email": fmt.Sprintf(validator.ErrorMessages["invalidEmail"], "email"),
		},
	},

	// No errors
	{
		input: model.User{
			Username: "someuser",
			Email:    "some@email.com",
		},
		expected: nil,
	},
}

func TestValidate(t *testing.T) {
	for _, test := range validateTests {
		output := validator.Validate(test.input)
		if output == nil && test.expected != nil {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
		for k, v := range output {
			if test.expected[k] != v {
				t.Errorf("Output %q not equal to expected %q", v, test.expected[k])
			}
		}
	}
}
