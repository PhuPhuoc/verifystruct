package verifystruct

import (
	"fmt"
	"testing"
)

// Define a struct to test verification
type testModel struct {
	Name  string `json:"name" verify:"required=true,type=string,min=4,max=20"`
	Age   int    `json:"age" verify:"required=true,type=int,min=4,max=20"`
	Email string `json:"email" verify:"type=email,min=4,max=20"`
}

// Test case structure
type testCase struct {
	input          map[string]any
	shouldFail     bool
	expectedErrors []error
}

func TestValidField(t *testing.T) {
	// Define test cases using a map with the test case name as the key
	tests := map[string]testCase{
		"Valid input": {
			input: map[string]any{
				"name": "phuoc",
				"age":  18,
			},
			shouldFail:     false,
			expectedErrors: nil,
		},
		"Extra field 'extra'": {
			input: map[string]any{
				"name":  "phuoc",
				"age":   18,
				"extra": "extra_field",
			},
			shouldFail:     true,
			expectedErrors: []error{fmt.Errorf("field 'extra' is invalid")},
		},
	}

	// Iterate over the test cases using sub-tests
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			model := testModel{}
			list_err := VerifyStruct(tt.input, model)
			checkErrors(t, list_err, tt.shouldFail, tt.expectedErrors)
		})
	}

}


func TestRequiredField(t *testing.T) {
	tests := map[string]testCase{
		"Valid input": {
			input: map[string]any{
				"name": "phuoc",
				"age":  18,
			},
			shouldFail:     false,
			expectedErrors: nil,
		},
		"Missing Required Field age": {
			input: map[string]any{
				"name":  "phuoc",
				"email": "bless@gmail.com",
			},
			shouldFail:     true,
			expectedErrors: []error{fmt.Errorf("field 'age' is required")},
		},
		"Missing Required Field 'name' vs 'age'": {
			input: map[string]any{
				"email": "bless@gmail.com",
			},
			shouldFail:     true,
			expectedErrors: []error{fmt.Errorf("field 'name' is required"), fmt.Errorf("field 'age' is required")},
		},
	}

	// Iterate over the test cases using sub-tests
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			model := testModel{}
			list_err := VerifyStruct(tt.input, model)
			checkErrors(t, list_err, tt.shouldFail, tt.expectedErrors)
		})
	}
}

// checkErrors is a helper function to verify the errors returned by VerifyStruct
func checkErrors(t *testing.T, actualErrors []error, shouldFail bool, expectedErrors []error) {
	if shouldFail {
		if actualErrors == nil {
			expectedErrorMess := convertSliceErrToAString(expectedErrors)
			t.Errorf("Expected error: %s ~ but got nil", expectedErrorMess)
		} else {
			var count int
			var expectMess string
			for i, expected := range expectedErrors {
				if i > 0 {
					expectMess += ", "
				}
				expectMess += expected.Error()
				for _, actual := range actualErrors {
					if expected.Error() == actual.Error() {
						count++
					}
				}
			}
			if count != len(expectedErrors) {
				t.Errorf("Expected error: %s ~ but got: %s'", expectMess, convertSliceErrToAString(actualErrors))
			}
		}
	} else {
		if actualErrors != nil {
			t.Errorf("Did not expect error, but got one: %v", actualErrors)
		}
	}
}

func convertSliceErrToAString(errs []error) string {
	var mess string
	for i, err := range errs {
		if i > 0 {
			mess += ", "
		}
		mess += err.Error()
	}
	return mess
}

func TestVerifyStruct(t *testing.T) {
	// Run all individual test cases
	t.Run("Valid input", TestValidField)
	t.Run("Required field missing", TestRequiredField)
}
