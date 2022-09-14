package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "ABC", "Delhi", "110011", "01-01-2000", "1"},
		{"1002", "XYZ", "Delhi", "110011", "01-01-2000", "1"},
	}
	return CustomerRepositoryStub{customers}
}
