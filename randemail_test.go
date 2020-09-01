package main

import (
	"testing"
	"time"
)

func TestGetRandomEmail(t *testing.T) {
	s1 := time.Now().UnixNano()
	e1 := getRandomEmail(s1)

	s2 := time.Now().UnixNano()
	e2 := getRandomEmail(s2)

	if !isValid(e1) {
		t.Errorf("getRandomEmail returns invalid email, got '%s'", e1)
		return
	}
	if !isValid(e2) {
		t.Errorf("getRandomEmail returns invalid email, got '%s'", e2)
		return
	}
	if e1 == e2 {
		t.Errorf("getRandomEmail returns the same email '%s'", e1)
	}
}
