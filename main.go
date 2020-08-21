package main

import (
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/verify", verifyHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := net.JoinHostPort("", port)
	log.Printf("Serving on %s...\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
