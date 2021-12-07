package role

type Service interface {
	GetAllRoles() ([]Role, error)
}

type RoleService struct {
	repository Repository
}

func NewService(repository Repository) *RoleService {
	return &RoleService{repository}
}

func (s *RoleService) GetAllRoles() ([]Role, error) {
	roles, err := s.repository.FindAll()
	if err != nil {
		return roles, err
	}
	return roles, nil
}
