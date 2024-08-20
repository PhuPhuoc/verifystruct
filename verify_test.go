package verifystruct

import "testing"

// Define a struct to test verification
type testModel struct {
	Name string `json:"name" verify:"required=true,type=string,min=4,max=20"`
	Age  int    `json:"age" verify:"required=true,type=int,min=4,max=20"`
}

func TestVerifyStruct(t *testing.T) {
	// Define test cases using a slice of structs
	tests := []struct {
		name       string
		input      map[string]any
		shouldFail bool
	}{
		{
			name: "Valid input",
			input: map[string]any{
				"name": "phuoc",
				"age":  18,
			},
			shouldFail: false,
		},
		{
			name: "Extra fields",
			input: map[string]any{
				"name":  "phuoc",
				"age":   18,
				"email": "extra_field",
			},
			shouldFail: true,
		},
		{
			name: "Extra fields",
			input: map[string]any{
				"name":  "phuoc",
				"age":   18,
				"phone": "extra_field",
			},
			shouldFail: true,
		},
	}

	// Iterate over the test cases using sub-tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			model := testModel{}
			err := VerifyStruct(tt.input, model)

			if tt.shouldFail && err == nil {
				t.Errorf("Expected an error, but got nil")
			}

			if !tt.shouldFail && err != nil {
				t.Errorf("Did not expect an error, but got one: %v", err)
			}
		})
	}
}
