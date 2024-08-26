package utils

import (
	"fmt"
	"testing"

	"github.com/PhuPhuoc/verifystruct/common"
	"github.com/PhuPhuoc/verifystruct/utils"
)

type testCase_ValidField struct {
	name             string
	requestDict      map[string]any
	standardFieldMap map[string]bool
	expectedErrors   []error
}

var testCases = []testCase_ValidField{
	{
		name: "All fields valid",
		requestDict: map[string]any{
			"field1": "value1",
			"field2": 2,
		},
		standardFieldMap: map[string]bool{
			"field1": true,
			"field2": true,
		},
		expectedErrors: []error{},
	},
	{
		name: "One invalid field",
		requestDict: map[string]any{
			"field1": "value1",
			"field2": 2,
			"field3": true,
		},
		standardFieldMap: map[string]bool{
			"field1": true,
			"field2": true,
		},
		expectedErrors: []error{fmt.Errorf("field3 is invalid field")},
	},
	{
		name: "Multiple invalid fields",
		requestDict: map[string]any{
			"field1": "value1",
			"field2": 2,
			"field3": true,
			"Field4": 4,
		},
		standardFieldMap: map[string]bool{
			"field1": true,
			"field2": true,
		},
		expectedErrors: []error{fmt.Errorf("field3 is invalid field"), fmt.Errorf("Field4 is invalid field")},
	},
	{
		name: "Case insensitive check",
		requestDict: map[string]any{
			"Field1": "value1",
			"FIELD2": 2,
		},
		standardFieldMap: map[string]bool{
			"field1": true,
			"field2": true,
		},
		expectedErrors: []error{},
	},
}

func TestCheckFieldNotExistInStandardModel(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			errors := utils.CheckFieldNotExistInStandardModel(tt.requestDict, tt.standardFieldMap)
			common.LogValidationDetails(tt.requestDict, errors, tt.expectedErrors)
			common.CompareErrorInTestcase(t, tt.name, errors, tt.expectedErrors)
		})
	}
}
