package main

import "testing"

func TestIsValid(t *testing.T) {
	emails := []string{
		"foo@bar.com",
		"##gh@hello",
		"b.Ryy.hooli@example.com",
	}

	vals := []bool{
		true,
		false,
		true,
	}

	for i, email := range emails {
		ret := isValid(email)
		if ret != vals[i] {
			t.Errorf("Unexpected return for email %s, got %t", email, ret)
		}
	}

}
