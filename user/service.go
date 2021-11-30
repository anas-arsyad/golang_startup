package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	LoginUser(input LoginUserInput) (User, error)
	IsEmailAvailable(input EmailUserInput) (bool, error)
	UploadAvatar(id int, fileLocation string) (User, error)
	GetUserById(id int) (User, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) *service {
	return &service{r}
}

func (s *service) IsEmailAvailable(input EmailUserInput) (bool, error) {
	email := input.Email

	user, err := s.repository.FindByEMail(email)
	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, err
	}
	return false, err
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.PasswordHash = string(passwordHash)
	user.Role = "user"

	newUser, err := s.repository.Save(user)
	if err != nil {
		return user, err
	}

	return newUser, err

}

func (s *service) LoginUser(input LoginUserInput) (User, error) {
	var user User
	userLogin, err := s.repository.FindByEMail(input.Email)
	if err != nil {
		return user, err
	}

	if userLogin.ID == 0 {
		return user, errors.New("User Not Found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userLogin.PasswordHash), []byte(input.Password))

	if err != nil {
		return user, err
	}

	return userLogin, nil
}

func (s *service) UploadAvatar(id int, fileLocation string) (User, error) {

	getUserById, err := s.repository.FindByID(id)
	if err != nil {
		return getUserById, err
	}
	getUserById.AvatarFileName = fileLocation

	user, err := s.repository.Update(getUserById)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) GetUserById(id int) (User, error) {

	userById, err := s.repository.FindByID(id)
	if err != nil {
		return userById, err

	}

	if userById.ID == 0 {
		return userById, errors.New("User not Found By that Id")

	}
	return userById, nil

}
