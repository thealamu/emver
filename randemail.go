package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var chars = []string{"a", "b", "c", "d", "e", "f", "g",
	"h", "i", "j", "k", "l", "m", "n", "o", "p", "q",
	"r", "s", "t", "u", "v", "w", "x", "y", "z",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J",
	"K", "L", "M", "N", "O", "P", "Q", "R", "S", "T",
	"U", "V", "W", "X", "Y", "Z"}

var domains = []string{"gmail.com", "outlook.com", "hey.com", "yahoo.com", "tempmail.com"}

func getRandomEmail(seed int64) string {
	rand.Seed(seed)
	// Build random username
	charLen := rand.Intn(11)
	overIndx := len(chars)
	var email strings.Builder
	for i := 0; i < charLen; i++ {
		indx := rand.Intn(overIndx)
		email.WriteString(chars[indx])
	}
	// Add random domain
	lenDom := len(domains)
	indx := rand.Intn(lenDom)
	email.WriteString(fmt.Sprintf("@%s", domains[indx]))

	return email.String()
}
