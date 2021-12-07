package supplier

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Supplier, error)
	FindByID(ID int) (Supplier, error)
	FindByIDWithCategoryProduct(ID int) (Supplier, error)
	Save(supplier Supplier) (Supplier, error)
	Update(supplier Supplier) (Supplier, error)
	Delete(supplierID int) (bool, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Supplier, error) {
	var supplier []Supplier

	err := r.db.Preload("CategoriProduct").Find(&supplier).Error
	if err != nil {
		return supplier, err
	}

	return supplier, nil
}

func (r *repository) FindByID(ID int) (Supplier, error) {
	var supplier Supplier
	err := r.db.Where("id = ?", ID).Find(&supplier).Error

	if err != nil {
		return supplier, err
	}

	return supplier, nil
}

func (r *repository) FindByIDWithCategoryProduct(ID int) (Supplier, error) {
	var supplier Supplier
	err := r.db.Preload("CategoriProduct").Where("id = ?", ID).Find(&supplier).Error

	if err != nil {
		return supplier, err
	}

	return supplier, nil
}

func (r *repository) Save(supplier Supplier) (Supplier, error) {
	err := r.db.Create(&supplier).Error
	if err != nil {
		return supplier, err
	}

	newSupplier, err := r.FindByIDWithCategoryProduct(supplier.ID)
	if err != nil {
		return newSupplier, err
	}

	return newSupplier, nil
}

func (r *repository) Update(supplier Supplier) (Supplier, error) {
	err := r.db.Save(&supplier).Error
	if err != nil {
		return supplier, err
	}

	updateSupplier, err := r.FindByIDWithCategoryProduct(supplier.ID)
	if err != nil {
		return updateSupplier, err
	}

	return updateSupplier, nil
}

func (r *repository) Delete(supplierID int) (bool, error) {
	var supplier Supplier

	err := r.db.Where("id = ?", supplierID).Delete(&supplier).Error
	if err != nil {
		return false, err
	}

	return true, nil
}
