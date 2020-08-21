package main

import (
	"net"
	"regexp"
	"strings"
)

// isValid returns true if the email is valid
func isValid(email string) bool {
	matcher := regexp.MustCompile(`^[_a-zA-Z0-9-]+(\.[_a-zA-Z0-9-]+)*@[a-zA-Z0-9-]+(\.[a-z0-9-]+)*(\.[a-z]{2,4})$`)
	return matcher.MatchString(email)
}

func getMXAddr(email string) (string, error) {
	parts := strings.Split(email, "@")
	domain := parts[len(parts)-1]
	mxs, err := net.LookupMX(domain)
	if err != nil {
		return "", err
	}
	prefed := mxs[0].Host
	return prefed, nil
}
