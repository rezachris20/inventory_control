package product

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Product, error)
	FindByID(ID int) (Product, error)
	FindByIDWithPreload(ID int) (Product, error)
	Save(product Product) (Product, error)
	Update(product Product) (Product, error)
	Delete(ID int) (bool, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Product, error) {
	var products []Product

	err := r.db.Preload("Supplier.CategoriProduct").Find(&products).Error
	if err != nil {
		return products, err
	}

	return products, nil
}

func (r *repository) FindByID(ID int) (Product, error) {
	var product Product

	err := r.db.Where("id = ? ", ID).Find(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) FindByIDWithPreload(ID int) (Product, error) {
	var product Product

	err := r.db.Preload("Supplier.CategoriProduct").Where("id = ?", ID).Find(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) Save(product Product) (Product, error) {
	err := r.db.Create(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) Update(product Product) (Product, error) {
	err := r.db.Save(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) Delete(ID int) (bool, error) {
	var product Product

	err := r.db.Where("id=?", ID).Delete(&product).Error
	if err != nil {
		return false, err
	}

	return true, nil
}
