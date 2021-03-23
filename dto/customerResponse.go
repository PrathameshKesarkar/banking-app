package dto

type CustomerResponse struct {
	Id          string `json:"id" xml:"id"`
	Name        string `json:"full_name" xml:"name"`
	City        string `json:"city" xml:"city"`
	Zipcode     string `json:"zip_code" xml:"zipcode"`
	DateofBirth string `json:"date_of_birth" xml:"dateofbirth"`
	Status      string `json:"status" xml:"status"`
}
