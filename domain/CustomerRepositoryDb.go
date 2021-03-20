package domain

import (
	"database/sql"
	"time"
	"github.com/PrathameshKesarkar/banking-app/errs"
	"github.com/PrathameshKesarkar/banking-app/logger"
	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {

	var rows *sql.Rows
	var err error
	if status == "" {
		findAllSQL := "Select customer_id, name, city, zipcode, date_of_birth, status from customers"
		rows, err = d.client.Query(findAllSQL)
	} else {
		findAllSQL := "Select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		rows, err = d.client.Query(findAllSQL, status)
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

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)

		if err != nil {
			logger.Error("Error while initalizing customer data")
			return nil, errs.NewUnexpectedError("Unexpected error happened")
		}
		customers = append(customers, c)
	}
	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	findCustomer := "Select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ? "

	row := d.client.QueryRow(findCustomer, id)
	var customer Customer
	err := row.Scan(&customer.Id, &customer.Name, &customer.City, &customer.Zipcode, &customer.DateofBirth, &customer.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("No Customer found for the id: "+id)
			return nil, errs.NewNotFoundError("Customer Not found")
		} else {
			logger.Error("Something went wrong while quering data in sql")
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}
	return &customer, nil
}

func NewCustomerRepositoryDB() CustomerRepository {
	client, err := sql.Open("mysql", "root:deadpool@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDb{client: client}
}
