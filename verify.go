package main

import "regexp"

// isValid returns true if the email is valid
func isValid(email string) bool {
	matcher := regexp.MustCompile(`^[_a-zA-Z0-9-]+(\.[_a-zA-Z0-9-]+)*@[a-zA-Z0-9-]+(\.[a-z0-9-]+)*(\.[a-z]{2,4})$`)
	return matcher.MatchString(email)
}
