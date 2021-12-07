package customer

import "errors"

type Service interface {
	RegisterCustomer(input RegisterCustomerInput) (Customer, error)
	UpdateCustomer(input InputDetailCustomer, inputData RegisterCustomerInput) (Customer, error)
	DeleteCustomer(input InputDeleteCustomer) (Customer, error)
	GetAllCustomer() ([]Customer, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterCustomer(input RegisterCustomerInput) (Customer, error) {
	customer := &Customer{
		Nama:  input.Nama,
		NoHp:  input.NoHp,
		Email: input.Email,
	}

	newCustomer, err := s.repository.Save(*customer)
	if err != nil {
		return newCustomer, err
	}

	return newCustomer, nil
}

func (s *service) UpdateCustomer(input InputDetailCustomer, inputData RegisterCustomerInput) (Customer, error) {

	customer, err := s.repository.FindByID(input.ID)
	if err != nil {
		return customer, err
	}

	if customer.ID == 0 {
		return customer, errors.New("No customer found")
	}

	if inputData.Email != "" {
		customer.Email = inputData.Email
	}

	updateCustomer, err := s.repository.Update(customer)
	if err != nil {
		return updateCustomer, err
	}

	return updateCustomer, nil
}

func (s *service) DeleteCustomer(input InputDeleteCustomer) (Customer, error) {

	customer, err := s.repository.FindByID(input.ID)
	if err != nil {
		return customer, err
	}

	if customer.ID == 0 {
		return customer, errors.New("No customer found")
	}

	deleteCustomer, err := s.repository.Delete(input.ID)
	if err != nil {
		return deleteCustomer, nil
	}

	return deleteCustomer, err
}

func (s *service) GetAllCustomer() ([]Customer, error) {
	customers, err := s.repository.FindAll()
	if err != nil {
		return customers, err
	}

	return customers, nil
}
