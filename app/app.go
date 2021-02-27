package app

import (
	"log"
	"net/http"
)

func Start() {
	// define the route
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomer)

	err := http.ListenAndServe("localhost:8000", nil)

	if err != nil {
		log.Panic(err)
	}
}
