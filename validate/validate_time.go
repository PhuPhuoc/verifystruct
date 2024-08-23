package validate

import "regexp"

func IsValidTime(value string) bool {
	re := regexp.MustCompile(`^(?:[01]?\d|2[0-3]):[0-5]\d$`)
	return re.MatchString(value)
}
