package user

import (
	"gorm.io/gorm"
)

type Repository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindByID(ID int) (User, error)
	FindByIDWithRole(ID int) (User, error)
	FindAllUser() ([]User, error)
	Update(user User) (User, error)
	Delete(userID int) (User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserRepository) FindByEmail(email string) (User, error) {
	var user User

	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserRepository) FindByID(ID int) (User, error) {
	var user User

	err := r.db.Where("id = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserRepository) FindByIDWithRole(ID int) (User, error) {
	var user User

	err := r.db.Preload("Role").Where("id = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserRepository) FindAllUser() ([]User, error) {
	var users []User

	err := r.db.Preload("Role").Find(&users).Error
	if err != nil {
		return users, err
	}

	return users, nil
}

func (r *UserRepository) Update(user User) (User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserRepository) Delete(userID int) (User, error) {
	var user User

	err := r.db.Where("id = ?", userID).Delete(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil

}
