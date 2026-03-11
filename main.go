package main

import (
	"fmt"
	"log"
	"net/http"

	httpserver "github.com/siwakasen/go-with-TDD/httpserver"
	vrdc "github.com/siwakasen/go-with-TDD/variadic"
)

func main() {
	numbers := []int{1, 2, 3, 4}

	// using other package's function
	fmt.Printf("%d\n", vrdc.Calculate(numbers...))

	// httpserver
	handler := http.HandlerFunc(httpserver.PlayerServer)
	log.Fatal(http.ListenAndServe(":5001", handler))
}
