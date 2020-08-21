package main

import "testing"

func TestVerify(t *testing.T) {
	emails := []string{
		"human@outlook.com",
		"foo.bar@gmail.com",
	}

	vals := []bool{
		true,
		true,
	}

	for i, email := range emails {
		ret, err := verify(email)
		if err != nil {
			t.Error(err)
		}
		if ret != vals[i] {
			t.Errorf("verify returned %t for email %s", ret, email)
		}
	}
}

func TestGetMxAddr(t *testing.T) {
	emails := []string{
		"foo.bar@gmail.com",
	}

	vals := []string{
		"gmail-smtp-in.l.google.com.",
	}

	for i, email := range emails {
		ret, err := getMXAddr(email)
		if err != nil {
			t.Error(err)
		}
		if ret != vals[i] {
			t.Errorf("getMXAddr returned %s for email %s, expected %s", ret, email, vals[i])
		}
	}
}

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
