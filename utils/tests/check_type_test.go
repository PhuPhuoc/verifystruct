package utils

import (
	"fmt"
	"testing"

	"github.com/PhuPhuoc/verifystruct/common"
	"github.com/PhuPhuoc/verifystruct/utils"
)

type testCase_CheckType struct {
	name           string
	requestDict    map[string]any
	verifyReqMap   map[string]map[string]string
	expectedErrors []error
}

var tc_checkValidType = []testCase_CheckType{
	{
		name: "All fields valid",
		requestDict: map[string]any{
			"field1": "value1",
			"field2": 42,
			"field3": true,
			"field4": "2024-08-23",
			"field5": "15:30",
			"field6": "example@gmail.com",
			"field7": "active",
		},
		verifyReqMap: map[string]map[string]string{
			"field1": {
				"type": "string",
			},
			"field2": {
				"type": "number",
			},
			"field3": {
				"type": "bool",
			},
			"field4": {
				"type": "date",
			},
			"field5": {
				"type": "time",
			},
			"field6": {
				"type": "email",
			},
			"field7": {
				"type": "enum[active,inactive]",
			},
		},
		expectedErrors: []error{},
	},
	{
		name: "Invalid string field",
		requestDict: map[string]any{
			"field1": 123, // Invalid: expected a string
		},
		verifyReqMap: map[string]map[string]string{
			"field1": {
				"type": "string",
			},
		},
		expectedErrors: []error{
			fmt.Errorf("field1 must be formatted as a string"),
		},
	},
	{
		name: "Invalid number field",
		requestDict: map[string]any{
			"field2": "not_a_number", // Invalid: expected a number
		},
		verifyReqMap: map[string]map[string]string{
			"field2": {
				"type": "number",
			},
		},
		expectedErrors: []error{
			fmt.Errorf("field2 must be formatted as a number"),
		},
	},
	{
		name: "Invalid bool field",
		requestDict: map[string]any{
			"field3": "not_a_bool", // Invalid: expected a boolean
		},
		verifyReqMap: map[string]map[string]string{
			"field3": {
				"type": "bool",
			},
		},
		expectedErrors: []error{
			fmt.Errorf("field3 must be formatted as a boolean"),
		},
	},
	{
		name: "Invalid date field",
		requestDict: map[string]any{
			"field4": "23-08-2024", // Invalid: expected format YYYY-MM-DD
		},
		verifyReqMap: map[string]map[string]string{
			"field4": {
				"type": "date",
			},
		},
		expectedErrors: []error{
			fmt.Errorf("field4 must be formatted as a date (YYYY-MM-DD)"),
		},
	},
	{
		name: "Invalid time field",
		requestDict: map[string]any{
			"field5": "3:30 PM", // Invalid: expected format HH:MM
		},
		verifyReqMap: map[string]map[string]string{
			"field5": {
				"type": "time",
			},
		},
		expectedErrors: []error{
			fmt.Errorf("field5 must be formatted as a time (HH:MM)"),
		},
	},
	{
		name: "Invalid email field",
		requestDict: map[string]any{
			"field6": "example.com", // Invalid: missing @domain
		},
		verifyReqMap: map[string]map[string]string{
			"field6": {
				"type": "email",
			},
		},
		expectedErrors: []error{
			fmt.Errorf("field6 must be formatted as an email (___@gmail.com)"),
		},
	},
	{
		name: "Invalid enum field",
		requestDict: map[string]any{
			"field7": "pending", // Invalid: not in enum[active,inactive]
		},
		verifyReqMap: map[string]map[string]string{
			"field7": {
				"type": "enum[active,inactive]",
			},
		},
		expectedErrors: []error{
			fmt.Errorf("field7 must be a string belonging to active,inactive"),
		},
	},
	{
		name: "Unsupported type",
		requestDict: map[string]any{
			"field8": "some_value", // Unsupported type
		},
		verifyReqMap: map[string]map[string]string{
			"field8": {
				"type": "unsupported_type",
			},
		},
		expectedErrors: []error{
			fmt.Errorf("unsupported data type: unsupported_type"),
		},
	},
}

func TestCheckValidType(t *testing.T) {
	for _, tt := range tc_checkValidType {
		t.Run(tt.name, func(t *testing.T) {
			ActualErrors := utils.CheckValidType(tt.requestDict, tt.verifyReqMap)
			common.LogValidationDetails(tt.requestDict, ActualErrors, tt.expectedErrors)
			common.CompareErrorInTestcase(t, tt.name, ActualErrors, tt.expectedErrors)
		})
	}
}
