package validate

import "regexp"

func IsValidEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@(gmail\.com)$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}
