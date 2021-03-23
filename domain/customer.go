package domain

import (
	"github.com/PrathameshKesarkar/banking-app/dto"
	"github.com/PrathameshKesarkar/banking-app/errs"
)

type Customer struct {
	Id          string `json:"id" xml:"id" db:"customer_id"`
	Name        string `json:"full_name" xml:"name"`
	City        string `json:"city" xml:"city"`
	Zipcode     string `json:"zip_code" xml:"zipcode"`
	DateofBirth string `json:"date_of_birth" xml:"dateofbirth" db:"date_of_birth"`
	Status      string `json:"status" xml:"status"`
}

type CustomerRepository interface {
	FindAll(string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}

func (cus Customer) ToDto() dto.CustomerResponse {
	return dto.CustomerResponse{Id: cus.Id,
		Name:        cus.Name,
		DateofBirth: cus.DateofBirth,
		City:        cus.City,
		Zipcode:     cus.Zipcode,
		Status:      cus.Status,
	}
}
