package utils

import (
	"fmt"
	"strings"
)

// CheckFieldNotExistInStandardModel checks if there are any fields in the request_dict
// that do not exist in the standardModel, as defined in listFieldMap. It returns a slice
// of error messages for any invalid fields found.
func CheckFieldNotExistInStandardModel(request_dict map[string]any, StandardFieldMap map[string]bool) []error {
	list_err := []error{}
	for key := range request_dict {
		if !StandardFieldMap[strings.ToLower(key)] {
			list_err = append(list_err, fmt.Errorf("field '%v' is invalid", key))
		}
	}
	return list_err
}
