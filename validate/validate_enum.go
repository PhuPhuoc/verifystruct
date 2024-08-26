package validate

import "strings"

func ExtractEnum(enumString string) string {
	trimmed := strings.TrimPrefix(enumString, "enum[")
	trimmed = strings.TrimSuffix(trimmed, "]")
	return trimmed
}

func IsValidEnum(value, enumStr string) bool {
	enum := ExtractEnum(enumStr)

	list_enum := strings.Split(enum, "-")
	for i := 0; i < len(list_enum); i++ {
		if list_enum[i] == value {
			return true
		}
	}
	return false
}
