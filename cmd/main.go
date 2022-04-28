package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	// r.HandleFunc("/", r)
	log.Fatal(http.ListenAndServe("localhost:9090", r))
}
