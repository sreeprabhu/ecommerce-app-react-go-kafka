package service

import (
	"go-react-ecommerce-app/internal/domain"
	"go-react-ecommerce-app/internal/dto"
	"log"
)

type UserService struct {
}

func (s UserService) SignUp(input dto.UserSignup) (string, error) {
	log.Println(input)

	// call db to create users

	return "this is my token", nil
}

func (s UserService) Login(input any) (string, error) {
	return "", nil
}

/*
*
camel case: Accessible inside this package only
*/
func (s UserService) findUserByEmail(email string) (*domain.User, error) {
	// perform some db operation
	// business logic
	return nil, nil
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
