// Package main
package main

import (
	"fmt"
	"log"
	"net/http"

	. "github.com/siwakasen/go-with-TDD/httpserver"
)

func main() {
	// httpserver
	port := 5001
	fmt.Printf("running go http server... on port %d", port)
	server := &PlayerServer{NewInMemoryPlayerStore()}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), server))
}
