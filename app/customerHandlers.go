package app

import (
	"encoding/json"
	"net/http"

	"github.com/PrathameshKesarkar/banking-app/service"
	"github.com/gorilla/mux"
)

type CustomerHandler struct {
	service service.CustomerService
}

func (ch *CustomerHandler) getAllCustomer(writer http.ResponseWriter, req *http.Request) {

	customers, err := ch.service.GetAllCustomer()

	if err != nil {
		writeResponse(writer, err.Code, err.AsMessage())
	} else {
		writeResponse(writer, 200, customers)
	}

}

func (ch *CustomerHandler) getCustomerById(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	customer, err := ch.service.GetCustomerById(vars["customer_id"])
	if err != nil {
		writeResponse(writer, err.Code, err.AsMessage())
	} else {
		writeResponse(writer, http.StatusOK, customer)
	}
}

func writeResponse(writer http.ResponseWriter, code int, data interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(code)
	if err := json.NewEncoder(writer).Encode(data); err != nil {
		panic(err)
	}

}
