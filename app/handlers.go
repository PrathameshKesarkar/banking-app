package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/PrathameshKesarkar/banking-app/service"
	"github.com/gorilla/mux"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zipcode"`
}

type CustomerHandler struct {
	service service.CustomerService
}

func (ch *CustomerHandler) getAllCustomer(writer http.ResponseWriter, req *http.Request) {

	contentType := req.Header.Get("Content-Type")

	customers, _ := ch.service.GetAllCustomer()

	if contentType == "application/xml" {
		writer.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(writer).Encode(customers)
	} else {
		writer.Header().Add("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(customers)
	}

}

func (ch *CustomerHandler) getCustomerById(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	customer, err := ch.service.GetCustomerById(vars["customer_id"])
	writer.Header().Add("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(err.Code)
		json.NewEncoder(writer).Encode(err)
		return
	}
	json.NewEncoder(writer).Encode(customer)
}
