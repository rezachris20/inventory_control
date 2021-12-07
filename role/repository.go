package role

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Role, error)
}

type RoleRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{db}
}

func (r *RoleRepository) FindAll() ([]Role, error) {
	var roles []Role

	err := r.db.Find(&roles).Error
	if err != nil {
		return roles, err
	}

	return roles, nil
}
