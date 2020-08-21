package main

import (
	"fmt"
	"log"
	"net/http"
)

func verifyHandler(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	valid, err := verify(email)
	log.Printf("Verifying %s", email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "%t", valid)
}
