package customer

import "gorm.io/gorm"

type Repository interface {
	Save(customer Customer) (Customer, error)
	Update(customer Customer) (Customer, error)
	FindByID(ID int) (Customer, error)
	Delete(ID int) (Customer, error)
	FindAll() ([]Customer, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(customer Customer) (Customer, error) {
	err := r.db.Create(&customer).Error
	if err != nil {
		return customer, err
	}

	return customer, nil
}

func (r *repository) Update(customer Customer) (Customer, error) {
	err := r.db.Save(&customer).Error
	if err != nil {
		return customer, err
	}

	return customer, nil
}

func (r *repository) FindByID(ID int) (Customer, error) {
	var customer Customer

	err := r.db.Where("id = ?", ID).Find(&customer).Error
	if err != nil {
		return customer, err
	}

	return customer, nil
}

func (r *repository) Delete(ID int) (Customer, error) {
	var customer Customer

	err := r.db.Where("id = ? ", ID).Delete(&customer).Error
	if err != nil {
		return customer, err
	}

	return customer, nil
}

func (r *repository) FindAll() ([]Customer, error) {
	var customer []Customer

	err := r.db.Find(&customer).Error
	if err != nil {
		return customer, err
	}

	return customer, nil
}
