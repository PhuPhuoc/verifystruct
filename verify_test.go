package verifystruct

import (
	"fmt"
	"testing"

	"github.com/PhuPhuoc/verifystruct/common"
)

// Define a struct to test verification
type testModel struct {
	Name      string `json:"name" verify:"required=true,type=string,min=5,max=20"`
	Age       int    `json:"age" verify:"required=true,type=number,min=18,max=50"`
	Email     string `json:"email" verify:"required=true,type=email"`
	Gender    bool   `json:"gender" verify:"type=bool"`
	Birthday  string `json:"birthday" verify:"type=date"`
	StartTime string `json:"start-time" verify:"type=time"`
	Role      string `json:"role" verify:"required=true,type=enum[staff-manager]"`
	Address   string `json:"address" verify:"type=string"`
	Yoe       int    `json:"yoe" verify:"type=number"`
}

type testCase struct {
	name           string
	requestDict    map[string]any
	expectedErrors []error
}

var tc_verifyStruct = []testCase{
	{
		name: "All fields valid",
		requestDict: map[string]any{
			"name":       "Phu Phuoc",
			"age":        23,
			"email":      "phuoc@gmail.com",
			"gender":     true,
			"birthday":   "2001-05-29",
			"start-time": "08:30",
			"role":       "manager",
		},
		expectedErrors: []error{},
	},
	{
		name: "Missing required field: name",
		requestDict: map[string]any{
			"age":        23,
			"email":      "phuoc@gmail.com",
			"gender":     true,
			"birthday":   "2001-05-29",
			"start-time": "08:30",
			"role":       "manager",
		},
		expectedErrors: []error{
			fmt.Errorf("name is required but missing"),
		},
	},
	{
		name: "Invalid type for field: age (should be number)",
		requestDict: map[string]any{
			"name":       "Phu Phuoc",
			"age":        "twenty-three", // Invalid type
			"email":      "phuoc@gmail.com",
			"gender":     true,
			"birthday":   "2001-05-29",
			"start-time": "08:30",
			"role":       "manager",
		},
		expectedErrors: []error{
			fmt.Errorf("age must be a numeric type and have a value from 18 to 50"),
		},
	},
	{
		name: "Field value out of range: age (min 18, max 50)",
		requestDict: map[string]any{
			"name":       "Phu Phuoc",
			"age":        60, // Out of range
			"email":      "phuoc@gmail.com",
			"gender":     true,
			"birthday":   "2001-05-29",
			"start-time": "08:30",
			"role":       "manager",
		},
		expectedErrors: []error{
			fmt.Errorf("age must be a numeric type and have a value from 18 to 50"),
		},
	},
	{
		name: "Invalid email format",
		requestDict: map[string]any{
			"name":       "Phu Phuoc",
			"age":        23,
			"email":      "invalid-email", // Invalid email format
			"gender":     true,
			"birthday":   "2001-05-29",
			"start-time": "08:30",
			"role":       "manager",
		},
		expectedErrors: []error{
			fmt.Errorf("email must be formatted as an email (___@gmail.com)"),
		},
	},
	{
		name: "Invalid enum value for role",
		requestDict: map[string]any{
			"name":       "Phu Phuoc",
			"age":        23,
			"email":      "phuoc@gmail.com",
			"gender":     true,
			"birthday":   "2001-05-29",
			"start-time": "08:30",
			"role":       "administrator", // Invalid enum value
		},
		expectedErrors: []error{
			fmt.Errorf("role must be a string belonging to staff or manager"),
		},
	},
	{
		name: "Field not allowed in model: extra_field",
		requestDict: map[string]any{
			"name":        "Phu Phuoc",
			"age":         23,
			"email":       "phuoc@gmail.com",
			"gender":      true,
			"birthday":    "2001-05-29",
			"start-time":  "08:30",
			"role":        "manager",
			"extra_field": "not_allowed", // Extra field not in the model
		},
		expectedErrors: []error{
			fmt.Errorf("extra_field is invalid field"),
		},
	},
	{
		name: "Multiple errors",
		requestDict: map[string]any{
			"age":        61,              // Out of range
			"email":      "invalid-email", // Invalid email format
			"start-time": "not-a-time",    // Invalid time format
		},
		expectedErrors: []error{
			fmt.Errorf("name is required but missing"),
			fmt.Errorf("age must be a numeric type and have a value from 18 to 50"),
			fmt.Errorf("email must be formatted as an email (___@gmail.com)"),
			fmt.Errorf("start-time must be formatted as a time (HH:MM)"),
			fmt.Errorf("role is required but missing"),
		},
	},
	{
		name: "Invalid date format for birthday",
		requestDict: map[string]any{
			"name":       "Phu Phuoc",
			"age":        23,
			"email":      "phuoc@gmail.com",
			"gender":     true,
			"birthday":   "29-05-2001", // Invalid date format
			"start-time": "08:30",
			"role":       "manager",
		},
		expectedErrors: []error{
			fmt.Errorf("birthday must be formatted as a date (YYYY-MM-DD)"),
		},
	},
	{
		name: "Invalid type yoe",
		requestDict: map[string]any{
			"name":       "Phu Phuoc",
			"age":        23,
			"email":      "phuoc@gmail.com",
			"gender":     true,
			"start-time": "08:30",
			"role":       "manager",
			"address":    "haha hihi huhuh",
			"yoe":        "haha",
		},
		expectedErrors: []error{
			fmt.Errorf("yoe must be a numeric type"),
		},
	},
}

func TestVerifyStruct(t *testing.T) {
	for _, tt := range tc_verifyStruct {
		t.Run(tt.name, func(t *testing.T) {
			ActualErrors := VerifyStruct(tt.requestDict, testModel{})
			common.LogValidationDetails(tt.requestDict, ActualErrors, tt.expectedErrors)
			common.CompareErrorInTestcase(t, tt.name, ActualErrors, tt.expectedErrors)
		})
	}
}
