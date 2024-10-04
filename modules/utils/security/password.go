package security

import "regexp"

func CheckPassword(password string) bool {
	/*
		Rules:
		Because of Bcrypt:
		length < 72
		length >= 12
		must contain:
		lower case
		upper case
		special character
		number
	*/
	lower := regexp.MustCompile(`[a-z]`)
	upper := regexp.MustCompile(`[A-Z]`)
	nonWord := regexp.MustCompile(`[\W]`)
	digit := regexp.MustCompile(`[\d]`)
	if len(password) < 12 || len(password) > 72 {
		return false
	}
	if lower.Match([]byte(password)) == false {
		return false
	}
	if upper.Match([]byte(password)) == false {
		return false
	}
	if nonWord.Match([]byte(password)) == false {
		return false
	}
	if digit.Match([]byte(password)) == false {
		return false
	}
	return true
}
