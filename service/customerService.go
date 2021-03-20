package service

import (
	"github.com/PrathameshKesarkar/banking-app/domain"
	"github.com/PrathameshKesarkar/banking-app/errs"
)

type CustomerService interface {
	GetAllCustomer(string) ([]domain.Customer, *errs.AppError)
	GetCustomerById(string) (*domain.Customer, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]domain.Customer, *errs.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}

	return s.repo.FindAll(status)
}

func (s DefaultCustomerService) GetCustomerById(id string) (*domain.Customer, *errs.AppError) {
	return s.repo.ById(id)
}

func NewCustomerService(repo domain.CustomerRepository) CustomerService {
	return DefaultCustomerService{repo: repo}
}
