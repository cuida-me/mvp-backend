package commons

import (
	"regexp"
	"unicode"
)

func IsValidEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	regex, err := regexp.Compile(emailRegex)
	if err != nil {
		return false
	}

	return regex.MatchString(email)
}

func IsValidUsername(username string) bool {
	usernameRegex := `^[a-zA-Z0-9]{3,13}$`

	regex, err := regexp.Compile(usernameRegex)
	if err != nil {
		return false
	}

	return regex.MatchString(username)
}

func IsValidPassword(s string) bool {
	var number, upper, special, sevenOrMore bool
	for _, c := range s {
		switch {
		case unicode.IsNumber(c):
			number = true
		case unicode.IsUpper(c):
			upper = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
		}
	}
	sevenOrMore = len(s) >= 7
	return number && upper && special && sevenOrMore
}
