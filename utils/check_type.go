package utils

import (
	"fmt"
	"strings"

	"github.com/PhuPhuoc/verifystruct/validate"
)

func CheckValidType(request_dict map[string]any, verifyReqMap map[string]map[string]string) []error {
	list_err := []error{}
	for curfield, curValue := range request_dict {
		mapTarget := verifyReqMap[curfield]
		targetType, typeExist := mapTarget["type"]
		if !typeExist {
			continue
		}
		if errValidType := compareTypeVerifyTagWithReqField(targetType, curfield, curValue); errValidType != nil {
			list_err = append(list_err, errValidType)
		}
	}
	return list_err
}

// type: string, bool, number, date, time, mail, enum
func compareTypeVerifyTagWithReqField(targetType, curfield string, currentVal any) error {
	switch targetType {
	case "string":
		fmt.Println("string")
		if _, ok := currentVal.(string); !ok {
			return fmt.Errorf("%s must be formatted as a string", curfield)
		}
	case "number":
		fmt.Println("number")
		if ok := validate.IsNumber(currentVal); !ok {
			return fmt.Errorf("%s must be formatted as a number", curfield)
		}
	case "bool":
		if _, ok := currentVal.(bool); !ok {
			return fmt.Errorf("%s must be formatted as a boolean", curfield)
		}
	case "date":
		if str, ok := currentVal.(string); ok {
			if flag := validate.IsValidDate(str); !flag {
				return fmt.Errorf("%s must be formatted as a date (YYYY-MM-DD)", curfield)
			}
		}
	case "time":
		if str, ok := currentVal.(string); ok {
			if flag := validate.IsValidTime(str); !flag {
				return fmt.Errorf("%s must be formatted as a time (HH:MM)", curfield)
			}
		}
	case "email":
		if str, ok := currentVal.(string); ok {
			if flag := validate.IsValidEmail(str); !flag {
				return fmt.Errorf("%s must be formatted as an email (___@gmail.com)", curfield)
			}
		}
	default:
		if strings.Contains(targetType, "enum") {
			if str, ok := currentVal.(string); ok {
				if flag := validate.IsValidEnum(str, targetType); !flag {
					return fmt.Errorf("%s must be a string belonging to %s", curfield, validate.ExtractEnum(targetType))
				}
			}
		} else {
			return fmt.Errorf("unsupported data type: %s", targetType)
		}
	}
	return nil
}
