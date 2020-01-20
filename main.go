package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const (
	// PORT to run the HTTP server
	PORT = 5000
)

func main() {

	r := mux.NewRouter()
	r.Path("/url").HandlerFunc(createHandler).Methods("POST")
	r.Path("/stats").HandlerFunc(statsHandler).Methods("GET")
	r.Path("/track/{id}").HandlerFunc(redirectHandler).Methods("GET")

	addr := fmt.Sprintf(":%d", PORT)

	log.Println("Listening on", addr)
	err := http.ListenAndServe(addr, r)
	if err != nil {
		log.Fatalln(err)
	}
}
