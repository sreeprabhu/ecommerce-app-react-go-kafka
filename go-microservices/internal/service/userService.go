package service

import (
	"errors"
	"go-react-ecommerce-app/internal/domain"
	"go-react-ecommerce-app/internal/dto"
	"go-react-ecommerce-app/internal/helper"
	"go-react-ecommerce-app/internal/repository"
)

type UserService struct {
	Repo repository.UserRepository
	Auth helper.Auth
}

func (s UserService) SignUp(input dto.UserSignup) (string, error) {

	hashPassword, err := s.Auth.CreateHashedPassword(input.Password)

	if err != nil {
		return "", err
	}

	// call db to create user
	user, err := s.Repo.CreateUser(domain.User{
		Email:    input.Email,
		Password: hashPassword,
		Phone:    input.Phone,
	})

	// generate and return token
	return s.Auth.GenerateToken(user.ID, user.Email, user.UserType)
}

func (s UserService) Login(email string, password string) (string, error) {

	user, err := s.findUserByEmail(email)
	if err != nil {
		return "", errors.New("user does not exist with the provided email id")
	}

	// verify password
	err = s.Auth.VerifyPassword(password, user.Password)

	if err != nil {
		return "", err
	}

	// generate and return token
	return s.Auth.GenerateToken(user.ID, user.Email, user.UserType)
}

/*
*
camel case: Accessible inside this package only
*/
func (s UserService) findUserByEmail(email string) (*domain.User, error) {
	// perform some db operation
	// business logic

	user, err := s.Repo.FindUser(email)

	return &user, err
}

func (s UserService) GetVerificationCode(u domain.User) (int, error) {
	return 0, nil
}

func (s UserService) VerifyCode(id uint, code int) (string, error) {
	return "", nil
}

func (s UserService) CreateProfile(id uint, input any) (string, error) {
	return "", nil
}

/*
*
Returns a pointer of User type, as at nany point of time we might need to edit the profile
*/
func (s UserService) GetProfile(id uint) (*domain.User, error) {
	return nil, nil
}

func (s UserService) UpdateProfile(id uint, input any) (string, error) {
	return "", nil
}

func (s UserService) FindCart(id uint) ([]interface{}, error) {
	return nil, nil
}

func (s UserService) CreateCart(input any, u domain.User) ([]interface{}, error) {
	return nil, nil
}

func (s UserService) CreateOrder(u domain.User) (int, error) {
	return 0, nil
}

func (s UserService) GetOrders(u domain.User) ([]interface{}, error) {
	return nil, nil
}

func (s UserService) GetOrderById(id uint, userId uint) ([]interface{}, error) {
	return nil, nil
}

func (s UserService) BecomeSeller(id uint, input any) (string, error) {
	return "", nil
}
