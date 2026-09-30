package main

import (
	"log"
	"net/http"
)

// 1. validasi ?
// 2. nentukan url ?
// 3. post get ?

func main() {

	// server

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "internal server error", http.StatusInternalServerError)

	})

	log.Println("running")
	err := http.ListenAndServe("localhost:8080", mux)

	if err != nil {
		panic(err)
	}
}
