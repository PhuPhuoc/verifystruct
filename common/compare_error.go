package common

import "testing"

func CompareErrorInTestcase(t *testing.T, testcaseName string, actualErrors, expectedErrors []error) {
	if len(actualErrors) != len(expectedErrors) {
		t.Errorf("\n\t + Testcase: %v - Expected error: %s ~ but got: %s \n", testcaseName, convertSliceErrToAString(expectedErrors), convertSliceErrToAString(actualErrors))
	} else {
		for _, actErr := range actualErrors {
			exist := false
			for _, expErr := range expectedErrors {
				if actErr.Error() == expErr.Error() {
					exist = true
					break
				}
			}
			if !exist {
				t.Errorf("\n\t + Testcase: %v - Unexpected error: %v \n", testcaseName, actErr)
			}
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
	if mess == "" {
		return "_nil_"
	}
	return mess
}
