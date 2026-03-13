package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/siwakasen/go-with-TDD/httpserver"
	vrdc "github.com/siwakasen/go-with-TDD/variadic"
)

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func main() {
	numbers := []int{1, 2, 3, 4}

	// using other package's function
	fmt.Printf("%d\n", vrdc.Calculate(numbers...))

	// httpserver
	server := &httpserver.PlayerServer{&InMemoryPlayerStore{}}
	log.Fatal(http.ListenAndServe(":5001", server))
}
