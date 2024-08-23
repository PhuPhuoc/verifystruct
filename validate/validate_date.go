package validate

import (
	"regexp"
	"strconv"
)

func IsValidDate(date string) bool {
	// YYYY-MM-DD
	dateRegex := `^(?P<Year>\d{4})-(?P<Month>0[1-9]|1[0-2])-(?P<Day>0[1-9]|[12]\d|3[01])$`
	re := regexp.MustCompile(dateRegex)

	if !re.MatchString(date) {
		return false
	}

	matches := re.FindStringSubmatch(date)

	year, err := strconv.Atoi(matches[1])
	if err != nil {
		return false
	}
	month, err := strconv.Atoi(matches[2])
	if err != nil {
		return false
	}
	day, err := strconv.Atoi(matches[3])
	if err != nil {
		return false
	}

	if (month == 4 || month == 6 || month == 9 || month == 11) && day > 30 {
		return false
	}
	if month == 2 {
		if isLeapYear(year) {
			if day > 29 {
				return false
			}
		} else {
			if day > 28 {
				return false
			}
		}
	}

	return true
}

func isLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || (year%400 == 0)
}
