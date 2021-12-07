package product

import (
	"backend-simple-pos/supplier"
	"errors"
	"os"
)

type Service interface {
	GetAllProducts() ([]Product, error)
	Save(inputProduct InputProduct) (Product, error)
	Update(inputDetailProduct InputDetailProduct, inputProduct InputProduct) (Product, error)
	Delete(inputDeleteProduct InputDeleteProduct) (bool, error)
}

type service struct {
	repository         Repository
	supplierRepository supplier.Repository
}

func NewService(repository Repository, supplierRepository supplier.Repository) *service {
	return &service{repository, supplierRepository}
}

func (s *service) GetAllProducts() ([]Product, error) {
	products, err := s.repository.FindAll()
	if err != nil {
		return products, err
	}

	return products, nil
}

func (s *service) Save(inputProduct InputProduct) (Product, error) {

	product := Product{
		Nama:       inputProduct.Nama,
		SupplierID: inputProduct.SupplierID,
		Satuan:     inputProduct.Satuan,
		Hargabeli:  inputProduct.Hargabeli,
		Hargajual:  inputProduct.Hargajual,
		Image:      inputProduct.ImagePath,
	}

	supplier, err := s.supplierRepository.FindByID(inputProduct.SupplierID)
	if supplier.ID == 0 {
		return product, errors.New("No found supplier ID")
	}

	saveProduct, err := s.repository.Save(product)
	if err != nil {
		return saveProduct, nil
	}

	newProduct, err := s.repository.FindByIDWithPreload(saveProduct.ID)
	if err != nil {
		return newProduct, err
	}

	return newProduct, nil
}

func (s *service) Update(inputDetailProduct InputDetailProduct, inputProduct InputProduct) (Product, error) {
	product, err := s.repository.FindByID(inputDetailProduct.ID)
	if err != nil {
		return product, err
	}

	// cek jika gambar tidak sama maka lakukan update dan hapus gambar
	if inputProduct.ImagePath != "" {
		if product.Image != "" {
			err := os.Remove(product.Image)
			if err != nil {
				return product, err
			}
		}
		product.Image = inputProduct.ImagePath
	}

	if inputProduct.Nama != "" {
		product.Nama = inputProduct.Nama
	}

	if inputProduct.SupplierID != 0 {
		product.SupplierID = inputProduct.SupplierID
	}

	if inputProduct.Satuan != "" {
		product.Satuan = inputProduct.Satuan
	}

	if inputProduct.Hargajual != 0 {
		product.Hargajual = inputProduct.Hargajual
	}

	if inputProduct.Hargabeli != 0 {
		product.Hargabeli = inputProduct.Hargabeli
	}

	updateProduct, err := s.repository.Update(product)
	if err != nil {
		return updateProduct, err
	}

	return updateProduct, nil
}

func (s *service) Delete(inputDeleteProduct InputDeleteProduct) (bool, error) {
	product, err := s.repository.FindByID(inputDeleteProduct.ID)
	if err != nil {
		return false, err
	}

	if product.ID == 0 {
		return false, errors.New("No product ID found")
	}

	if product.Image != "" {
		err := os.Remove(product.Image)
		if err != nil {
			return false, err
		}
	}

	deleteProduct, err := s.repository.Delete(inputDeleteProduct.ID)
	if err != nil {
		return deleteProduct, nil
	}

	return deleteProduct, err
}
