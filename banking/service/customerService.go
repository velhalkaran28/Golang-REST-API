package service

import (
	"banking-restapi/domain"
	"banking-restapi/dto"
	"banking-restapi/errs"
)

type CustomerService interface {
	GetAllCustomer(string) ([]dto.CustomerResponse, *errs.AppErr)
	GetCustomer(id string) (*dto.CustomerResponse, *errs.AppErr)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]dto.CustomerResponse, *errs.AppErr) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}
	c, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}
	customers := make([]dto.CustomerResponse, 0)
	for _, cust := range c {
		customers = append(customers, cust.ToDto())
	}
	return customers, nil
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppErr) {
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}
	response := c.ToDto()
	return &response, nil
}

func NewCustomerService(repo domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo}
}
