package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/PrathameshKesarkar/banking-app/service"
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
