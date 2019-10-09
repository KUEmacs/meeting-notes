package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello, world")
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/", hello)
	fmt.Println("Started listening on 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
