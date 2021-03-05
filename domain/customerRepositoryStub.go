package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepository {
	customers := []Customer{
		{Id: "1001", Name: "Prathamesh Kesarkar", City: "Mumbai", Zipcode: "480012", DateofBirth: "1995-07-15", Status: "1"},
		{Id: "1002", Name: "Shoubik Gosh", City: "Pune", Zipcode: "650013", DateofBirth: "1996-06-14", Status: "0"},
	}
	return CustomerRepositoryStub{customers: customers}
}
