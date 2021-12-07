package supplier

import (
	"backend-simple-pos/category_product"
	"errors"
)

type Service interface {
	GetAllSupplier() ([]Supplier, error)
	RegisterSupplier(inputData InputRegisterSupplier) (Supplier, error)
	UpdateSupplier(inputID InputDetailSupplier, inputData InputRegisterSupplier) (Supplier, error)
	DeleteSupplier(inputID InputDeleteSupplier) (bool, error)
}

type service struct {
	repository                Repository
	repositoryCategoryProduct category_product.Repository
}

func NewService(repository Repository, repositoryCategoryProduct category_product.Repository) *service {
	return &service{repository, repositoryCategoryProduct}
}

func (s *service) GetAllSupplier() ([]Supplier, error) {
	suppliers, err := s.repository.FindAll()

	if err != nil {
		return suppliers, err
	}

	return suppliers, nil
}

func (s *service) RegisterSupplier(inputData InputRegisterSupplier) (Supplier, error) {

	supplier := Supplier{
		Nama:              inputData.Nama,
		NoHp:              inputData.NoHp,
		Alamat:            inputData.Alamat,
		Email:             inputData.Email,
		CategoriProductID: inputData.CategoriProductID,
	}

	cek, err := s.repositoryCategoryProduct.FindByID(inputData.CategoriProductID)
	if err != nil {
		return supplier, err
	}

	if cek.ID == 0 {
		return supplier, errors.New("No category product found")
	}

	newSupplier, err := s.repository.Save(supplier)
	if err != nil {
		return newSupplier, err
	}

	return newSupplier, nil
}

func (s *service) UpdateSupplier(inputID InputDetailSupplier, inputData InputRegisterSupplier) (Supplier, error) {
	supplier, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return supplier, err
	}

	if supplier.ID == 0 {
		return supplier, errors.New("No found supplier that ID")
	}

	categoryProduct, _ := s.repositoryCategoryProduct.FindByID(inputData.CategoriProductID)
	if categoryProduct.ID == 0 {
		return supplier, errors.New("No found category product that ID")
	}

	if inputData.Nama != "" {
		supplier.Nama = inputData.Nama
	}

	if inputData.NoHp != "" {
		supplier.NoHp = inputData.NoHp
	}

	if inputData.Alamat != "" {
		supplier.Alamat = inputData.Alamat
	}

	if inputData.Email != "" {
		supplier.Email = inputData.Email
	}

	if inputData.CategoriProductID != 0 {
		supplier.CategoriProductID = inputData.CategoriProductID
	}

	updated, err := s.repository.Update(supplier)
	if err != nil {
		return updated, err
	}

	updatedSupplier, err := s.repository.FindByIDWithCategoryProduct(updated.ID)
	if err != nil {
		return updatedSupplier, err
	}

	return updatedSupplier, nil
}

func (s *service) DeleteSupplier(inputID InputDeleteSupplier) (bool, error) {

	cek, err := s.repository.FindByID(inputID.ID)

	if cek.ID == 0 {
		return false, errors.New("No supplier found that ID")
	}

	deleteSupplier, err := s.repository.Delete(inputID.ID)
	if err != nil {
		return deleteSupplier, err
	}

	return deleteSupplier, nil
}
