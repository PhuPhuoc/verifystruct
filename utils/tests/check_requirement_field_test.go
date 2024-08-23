package utils

import (
	"fmt"
	"testing"

	"github.com/PhuPhuoc/verifystruct/common"
	"github.com/PhuPhuoc/verifystruct/utils"
)

type testCase_CheckRequirementField struct {
	name           string
	requestDict    map[string]any
	verifyReqMap   map[string]map[string]string
	expectedErrors []error
}

var tc_CheckRequirementField = []testCase_CheckRequirementField{
	{
		name: "All fields valid",
		requestDict: map[string]any{
			"field1": "value1",
			"field2": 2,
		},
		verifyReqMap: map[string]map[string]string{
			"field1": {
				"required": "true",
				"type":     "string",
			},
			"field2": {
				"required": "true",
				"type":     "int",
			},
		},
		expectedErrors: []error{},
	},
	{
		name: "Missing required field",
		requestDict: map[string]any{
			"field1": "value1",
		},
		verifyReqMap: map[string]map[string]string{
			"field1": {
				"required": "true",
				"type":     "string",
			},
			"field2": {
				"required": "true",
				"type":     "int",
			},
		},
		expectedErrors: []error{
			fmt.Errorf("field 'field2' is required"),
		},
	},
	{
		name: "No required fields",
		requestDict: map[string]any{
			"field1": "value1",
		},
		verifyReqMap: map[string]map[string]string{
			"field1": {
				"type": "string",
			},
			"field2": {
				"type": "int",
			},
		},
		expectedErrors: []error{},
	},
	{
		name: "Multiple missing required fields",
		requestDict: map[string]any{
			"field2": "",
		},
		verifyReqMap: map[string]map[string]string{
			"field1": {
				"required": "true",
				"type":     "string",
			},
			"field2": {
				"required": "true",
				"type":     "int",
			},
		},
		expectedErrors: []error{
			fmt.Errorf("field 'field1' is required"),
			fmt.Errorf("field 'field2' is required and cannot be empty"),
		},
	},
	{
		name: "Required field with empty value",
		requestDict: map[string]any{
			"field1": "",
		},
		verifyReqMap: map[string]map[string]string{
			"field1": {
				"required": "true",
				"type":     "string",
			},
		},
		expectedErrors: []error{
			fmt.Errorf("field 'field1' is required and cannot be empty"),
		},
	},
}

// go test -v ./utils/tests -run TestCheckRequirementField
func TestCheckRequirementField(t *testing.T) {
	for _, tt := range tc_CheckRequirementField {
		t.Run(tt.name, func(t *testing.T) {
			errors := utils.CheckRequirementField(tt.requestDict, tt.verifyReqMap)
			common.LogValidationDetails(tt.requestDict, errors, tt.expectedErrors)
			common.CompareErrorInTestcase(t, tt.name, errors, tt.expectedErrors)
		})
	}
}
