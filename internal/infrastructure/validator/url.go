package validator

import (
	"net/url"
	"regexp"
)

func IsValidURL(urlString string) bool {
	_, err := url.ParseRequestURI(urlString)
	return err == nil
}

func IsValidEmail(email string) bool {
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}
