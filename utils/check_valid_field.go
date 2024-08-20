package utils

import (
	"fmt"
	"strings"
)

// CheckFieldNotExistInStandardModel checks if there are any fields in the request_dict
// that do not exist in the standardModel, as defined in listFieldMap. It returns a slice
// of error messages for any invalid fields found.
func CheckFieldNotExistInStandardModel(request_dict map[string]any, listFieldMap map[string]bool) []string {
	list_err := []string{}
	for key := range request_dict {
		if !listFieldMap[strings.ToLower(key)] {
			list_err = append(list_err, fmt.Sprintf("field '%v' is invalid", key))
		}
	}
	return list_err
}
