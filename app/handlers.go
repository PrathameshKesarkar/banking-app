package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zipcode"`
}

func greet(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(writer, "Hello World")
}

func getAllCustomer(writer http.ResponseWriter, req *http.Request) {
	customers := []Customer{
		{Name: "Prathamesh Kesarkar", City: "Mumbai", Zipcode: "480012"},
		{Name: "Shoubik Gosh", City: "Pune", Zipcode: "650013"},
	}

	contentType := req.Header.Get("Content-Type")

	if contentType == "application/xml" {
		writer.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(writer).Encode(customers)
	} else {
		writer.Header().Add("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(customers)
	}

}

func getCustomer(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	fmt.Fprint(writer, vars["customer_id"])
}

func createCustomer(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprint(writer, "Post request")
}
