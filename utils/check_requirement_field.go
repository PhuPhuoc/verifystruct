package utils

import (
	"fmt"
)

func CheckRequirementField(request_dict map[string]any, verifyReqMap map[string]map[string]string) []error {
	list_err := []error{}
	for field := range verifyReqMap {
		fieldKeyMap := verifyReqMap[field]
		tagRequiredVal, tagRequiredExist := fieldKeyMap["required"]
		if !tagRequiredExist {
			continue
		}
		if tagRequiredVal == "true" {
			if _, fieldExist := request_dict[field]; !fieldExist {
				list_err = append(list_err, fmt.Errorf("field '%v' is required", field))
			}
		}
	}
	return list_err
}
