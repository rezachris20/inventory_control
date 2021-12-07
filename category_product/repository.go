package category_product

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]CategoriProduct, error)
	FindByID(ID int) (CategoriProduct, error)
	Save(categoryProduct CategoriProduct) (CategoriProduct, error)
	Update(categoryProduct CategoriProduct) (CategoriProduct,error)
	Delete(ID int) (bool,error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(categoryProduct CategoriProduct) (CategoriProduct, error) {
	err := r.db.Create(&categoryProduct).Error
	if err != nil {
		return categoryProduct,err
	}

	return categoryProduct,nil
}

func (r *repository) FindAll() ([]CategoriProduct, error) {
	var categoryProduct []CategoriProduct

	err := r.db.Find(&categoryProduct).Error
	if err != nil {
		return categoryProduct,err
	}

	return categoryProduct,nil
}

func (r *repository) FindByID(ID int) (CategoriProduct, error) {
	var categoryProduct CategoriProduct

	err := r.db.Where("id = ? ", ID).Find(&categoryProduct).Error
	if err != nil {
		return categoryProduct,err
	}

	return categoryProduct,nil
}

func (r *repository) Update(categoryProduct CategoriProduct) (CategoriProduct,error) {
	err := r.db.Save(&categoryProduct).Error
	if err != nil {
		return categoryProduct,err
	}

	return categoryProduct,nil
}

func (r*repository) Delete(ID int) (bool,error) {
	var categoryProduct CategoriProduct

	err := r.db.Where("id = ? ",ID).Delete(&categoryProduct).Error
	if err != nil {
		return false, err
	}

	return true,err
}