package user

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	GetUserByID(ID int) (User, error)
	GetAllUsers() ([]User, error)
	UpdateUser(inputID GetUserDetailInput, inputData UpdateUserInput) (User, error)
	DeleteUser(inputID GetUserDeleteInput) (User, error)
	SaveAvatar(ID int, fileLocation string) (User, error)
}

type UserService struct {
	repository Repository
}

func NewService(repository Repository) *UserService {
	return &UserService{repository}
}

func (s *UserService) RegisterUser(input RegisterUserInput) (User, error) {
	user := &User{}
	user.Nama = input.Nama
	user.Username = input.Username
	user.Email = input.Email
	user.RoleID = input.RoleID

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return *user, err
	}

	user.Password = string(passwordHash)

	newUser, err := s.repository.Save(*user)
	if err != nil {
		return newUser, err
	}

	currentUser, err := s.repository.FindByIDWithRole(newUser.ID)
	if err != nil {
		return currentUser, err
	}

	return currentUser, nil
}

func (s *UserService) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("User tidak terdaftar dengan email tersebut")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *UserService) GetUserByID(ID int) (User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("No user found with that ID")
	}

	return user, nil
}

func (s *UserService) GetAllUsers() ([]User, error) {
	users, err := s.repository.FindAllUser()
	if err != nil {
		return users, err
	}

	return users, nil
}

func (s *UserService) UpdateUser(inputID GetUserDetailInput, inputData UpdateUserInput) (User, error) {
	user, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("No user found")
	}

	if inputData.Nama != "" {
		user.Nama = inputData.Nama
	}

	if inputData.Username != "" {
		user.Username = inputData.Username
	}

	if inputData.Email != "" {
		user.Email = inputData.Email
	}

	if inputData.RoleID != 0 {
		user.RoleID = inputData.RoleID
	}
	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}
	currentUser, err := s.repository.FindByIDWithRole(updatedUser.ID)
	if err != nil {
		return currentUser, err
	}

	return currentUser, nil
}

func (s *UserService) DeleteUser(inputID GetUserDeleteInput) (User, error) {

	getUser, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return getUser, err
	}

	if getUser.ID == 0 {
		return getUser, errors.New("No user found")
	}

	user, err := s.repository.Delete(inputID.ID)

	if err != nil {
		return user, err
	}
	return user, nil

}

func (s *UserService) SaveAvatar(ID int, fileLocation string) (User, error) {
	user, err := s.repository.FindByID(ID)

	if err != nil {
		return user, err
	}

	user.AvatarFileName = fileLocation
	userUpdate, err := s.repository.Update(user)
	if err != nil {
		return userUpdate, err
	}

	return userUpdate, nil
}
