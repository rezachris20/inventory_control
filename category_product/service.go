package category_product

import "errors"

type Service interface {
	GetAllCategoryProduct() ([]CategoriProduct, error)
	RegisterCategoryProduct(input InputCategoryProduct) (CategoriProduct, error)
	UpdateCategoryProduct(inputID InputDetailCategoryProduct, inputData InputCategoryProduct) (CategoriProduct, error)
	DeleteCategoryProduct(inputID InputDeleteCategoryProduct) (bool,error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAllCategoryProduct() ([]CategoriProduct, error) {
	categoryProduct, err := s.repository.FindAll()
	if err != nil {
		return categoryProduct, err
	}

	return categoryProduct, nil
}

func (s *service) RegisterCategoryProduct(input InputCategoryProduct) (CategoriProduct, error) {
	categoryProduct := CategoriProduct{
		Nama: input.Nama,
	}

	newCategoryProduct, err := s.repository.Save(categoryProduct)
	if err != nil {
		return newCategoryProduct, err
	}

	return newCategoryProduct, nil
}

func (s *service) UpdateCategoryProduct(inputID InputDetailCategoryProduct, inputData InputCategoryProduct) (CategoriProduct, error) {
	categoryProduct, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return categoryProduct, err
	}

	if categoryProduct.ID == 0 {
		return categoryProduct, errors.New("No category found")
	}

	categoryProduct.Nama = inputData.Nama

	updateCategoryProduct, err := s.repository.Update(categoryProduct)
	if err != nil {
		return updateCategoryProduct, err
	}

	return updateCategoryProduct, nil
}

func (s *service) DeleteCategoryProduct(inputID InputDeleteCategoryProduct) (bool,error) {

	categoryProduct, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return false, err
	}

	if categoryProduct.ID == 0 {
		return true, errors.New("No category product found")
	}

	deleteProductCategory, err := s.repository.Delete(inputID.ID)
	if err != nil {
		return deleteProductCategory, err
	}

	return deleteProductCategory,nil
}