package utils

import (
	"fmt"
	"regexp"
	"strconv"
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
		var minStr, maxStr *string
		tagMinVal, tagMinExist := mapTarget["min"]
		tagMaxVal, tagMaxExist := mapTarget["max"]
		if tagMinExist {
			minStr = &tagMinVal
		} else {
			minStr = nil
		}
		if tagMaxExist {
			maxStr = &tagMaxVal
		} else {
			maxStr = nil
		}
		if errValidType := compareTypeVerifyTagWithReqField(targetType, curfield, curValue, minStr, maxStr); errValidType != nil {
			list_err = append(list_err, errValidType)
		}
	}
	return list_err
}

// type: string, bool, number, date, time, mail, enum
func compareTypeVerifyTagWithReqField(targetType, curfield string, currentVal any, minStr, maxStr *string) error {
	switch targetType {
	case "string":
		min, max := convertToInt(minStr, maxStr)
		if ok := validate.IsString(currentVal, min, max); !ok {
			if min != nil && max != nil {
				return fmt.Errorf("%s must be a string type and have a string length of %d - %d characters", curfield, *min, *max)
			} else if min != nil {
				return fmt.Errorf("%s must be a string type and have a minimum string length of %d characters", curfield, *min)
			} else if max != nil {
				return fmt.Errorf("%s must be a string type and have a maximum string length of %d characters", curfield, *max)
			} else {
				return fmt.Errorf("%s must be a string type", curfield)
			}
		}
	case "number":
		min, max := convertToInt(minStr, maxStr)
		if ok := validate.IsNumber(currentVal, min, max); !ok {
			if min != nil && max != nil {
				return fmt.Errorf("%s must be a numeric type and have a value from %d to %d", curfield, *min, *max)
			} else if min != nil {
				return fmt.Errorf("%s must be a numeric type and have a minimum value of %d", curfield, *min)
			} else if max != nil {
				return fmt.Errorf("%s must be a numeric type and have a maximum value of %d", curfield, *max)
			} else {
				return fmt.Errorf("%s must be a numeric type", curfield)
			}
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
					messErr := validate.ExtractEnum(targetType)
					re := regexp.MustCompile(`-`)
					messErr = re.ReplaceAllString(messErr, " or ")
					return fmt.Errorf("%s must be a string belonging to %s", curfield, messErr)
				}
			}
		} else {
			return fmt.Errorf("unsupported data type: %s", targetType)
		}
	}
	return nil
}

func convertToInt(s1, s2 *string) (*int, *int) {
	var i1, i2 *int

	if s1 != nil {
		if val, err := strconv.Atoi(*s1); err == nil {
			i1 = &val
		}
	}

	if s2 != nil {
		if val, err := strconv.Atoi(*s2); err == nil {
			i2 = &val
		}
	}
	return i1, i2
}
