package domain

import "banking-restapi/errs"

type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateofBirth string
	Status      string
}

type CustomerRepository interface {
	FindAll(string) ([]Customer, *errs.AppErr)
	ById(string) (*Customer, *errs.AppErr)
}
