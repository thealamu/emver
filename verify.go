package main

import (
	"errors"
	"net"
	"net/smtp"
	"regexp"
	"strings"
)

var errInvalidEmail = errors.New("Email is invalid")

// isValid returns true if the email is valid
func isValid(email string) bool {
	matcher := regexp.MustCompile(`^[_a-zA-Z0-9-]+(\.[_a-zA-Z0-9-]+)*@[a-zA-Z0-9-]+(\.[a-z0-9-]+)*(\.[a-z]{2,4})$`)
	return matcher.MatchString(email)
}

func getMXAddr(email string) (string, error) {
	// split on @
	parts := strings.Split(email, "@")
	// pick the last split item, that's the domain
	domain := parts[len(parts)-1]
	// do the lookup
	mxs, err := net.LookupMX(domain)
	if err != nil {
		return "", err
	}
	// return the first, most prefered address
	prefed := mxs[0].Host
	return prefed, nil
}

func verify(email string) (bool, error) {
	if !isValid(email) {
		return false, errInvalidEmail
	}

	host, err := getMXAddr(email)
	if err != nil {
		return false, err
	}

	addr := net.JoinHostPort(host, "25")
	conn, err := smtp.Dial(addr)
	if err != nil {
		return false, err
	}

	err = conn.Verify(email)
	if err != nil {
		return false, err
	}

	return true, nil
}
