package app

import (
	"log"
	"net/http"

	"github.com/PrathameshKesarkar/banking-app/domain"
	"github.com/PrathameshKesarkar/banking-app/service"
	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	customerHandler := CustomerHandler{service: service.NewCustomerService(domain.NewCustomerRepositoryDB())}

	// define the route
	router.HandleFunc("/customers", customerHandler.getAllCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", customerHandler.getCustomerById).Methods(http.MethodGet)

	err := http.ListenAndServe("localhost:8000", router)

	if err != nil {
		log.Panic(err)
	}
}
