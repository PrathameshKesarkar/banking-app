package domain

import (
	"database/sql"
	"time"

	"github.com/PrathameshKesarkar/banking-app/errs"
	"github.com/PrathameshKesarkar/banking-app/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {

	var customers = make([]Customer, 0)
	var err error
	if status == "" {
		findAllSQL := "Select customer_id , name, city, zipcode, date_of_birth, status from customers"
		err = d.client.Select(&customers, findAllSQL)
		// rows, err = d.client.Query(findAllSQL)
	} else {
		findAllSQL := "Select customer_id , name, cit y, zipcode, date_of_birth, status from customers where status = ?"
		err = d.client.Select(&customers, findAllSQL, status)
		// rows, err = d.client.Query(findAllSQL, status)
	}

	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("error while fetching")
			return nil, errs.NewNotFoundError("Error while fetching customers")
		} else {
			logger.Error("Error with SQL connection")
			return nil, errs.NewUnexpectedError("Something went wrong")
		}
	}

	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	findCustomer := "Select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ? "

	var customer Customer
	
	err := d.client.Get(&customer,findCustomer,id)
	
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("No Customer found for the id: " + id)
			return nil, errs.NewNotFoundError("Customer Not found")
		} else {
			logger.Error("Sometng went wrong while quering data in sql")
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}
	return &customer, nil
}

func NewCustomerRepositoryDB() CustomerRepository {
	client, err := sqlx.Open("mysql", "root:deadpool@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDb{client: client}
}
