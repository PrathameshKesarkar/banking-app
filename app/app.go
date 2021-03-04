package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	// define the route
	router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/customers", getAllCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customer", createCustomer).Methods(http.MethodPost)

	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)

	err := http.ListenAndServe("localhost:8000", router)

	if err != nil {
		log.Panic(err)
	}
}
