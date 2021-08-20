package main

import (
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("hello pesy\n"))
}

func main() {
	http.HandleFunc("/", Handler)
	log.Println("Start on 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
