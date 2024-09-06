package userservice

import (
	"errors"
	"fmt"
	"github.com/salehborhani/todo-list/entity"
	"github.com/salehborhani/todo-list/pkg/hash"
	"github.com/salehborhani/todo-list/pkg/phonenumber"
	"github.com/salehborhani/todo-list/server/httpserver/jwt"
)

type Repository interface {
	RepoRegister(u entity.User) (entity.User, error)
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	IsAuthenticated(userName, password string) (bool, error)
}

type Service struct {
	repo Repository
}

// Register data format

type RegisterRequest struct {
	UserName    string `json:"user_name"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type RegisterResponse struct {
	User entity.User
}

// Login data format

type LoginRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string
}

func New(repo Repository) Service {
	return Service{repo: repo}
}

func (s Service) Register(req RegisterRequest) (RegisterResponse, error) {
	// Validate Phone Number
	ok := phonenumber.IsValid(req.PhoneNumber)
	if !ok {
		return RegisterResponse{}, errors.New("phone number is not valid")
	}

	// TODO - verify the Phone Number with OTP
	// Check if the Phone Number is unique or not
	if isUnique, err := s.repo.IsPhoneNumberUnique(req.PhoneNumber); err != nil || !isUnique {
		if err != nil {
			return RegisterResponse{}, fmt.Errorf("unexpected error: %w", err)
		}

		if !isUnique {
			return RegisterResponse{}, fmt.Errorf(`{"message": "phone number is not unique."}`)
		}
	}
	// Validate password
	if len(req.Password) < 8 {
		return RegisterResponse{}, fmt.Errorf(`{"message": "length of the password should not be less than 8."}`)
	}

	// Validate user name
	if len(req.UserName) < 3 {
		return RegisterResponse{}, fmt.Errorf(`{"message": "length of the username should not be less than 3."}`)
	}

	// Save the Hashed password in database
	hashedPass, err := hash.HashPassword(req.Password)
	if err != nil {
		return RegisterResponse{}, err
	}

	// Create user in the repository
	user := entity.User{
		ID:          0,
		UserName:    req.UserName,
		Password:    hashedPass,
		PhoneNumber: req.PhoneNumber,
	}

	createdUser, err := s.repo.RepoRegister(user)
	if err != nil {
		return RegisterResponse{}, err
	}

	return RegisterResponse{User: createdUser}, nil

}

func (s Service) Login(l LoginRequest) (LoginResponse, error) {
	// Connect to the database the check the creds
	ok, err := s.repo.IsAuthenticated(l.UserName, l.Password)

	if err != nil {
		return LoginResponse{}, err
	}

	if !ok {
		return LoginResponse{}, err
	}

	token, err := jwt.CreateToken(l.UserName)

	if err != nil {
		return LoginResponse{}, err
	}

	return LoginResponse{Token: token}, nil
}
