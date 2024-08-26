package utils

import (
	"fmt"
	"reflect"
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
			fieldValue, fieldExist := request_dict[field]
			if !fieldExist {
				list_err = append(list_err, fmt.Errorf("%v is required but missing", field))
			} else {
				if isEmpty(fieldValue) {
					list_err = append(list_err, fmt.Errorf("%v is required and cannot be empty", field))
				}
			}
		}
	}
	return list_err
}

func isEmpty(value any) bool {
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.String, reflect.Slice, reflect.Array, reflect.Map, reflect.Chan:
		return v.Len() == 0
	case reflect.Ptr, reflect.Interface:
		return v.IsNil()
	}
	return false
}
