package app

import (
	"log"
	"net/http"
)

func Start() {
	mux := http.NewServeMux()
	// define the route
	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/customers", getAllCustomer)

	err := http.ListenAndServe("localhost:8000", mux)

	if err != nil {
		log.Panic(err)
	}
}
