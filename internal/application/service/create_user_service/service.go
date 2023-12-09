package create_user_service

import (
	"balance-service/internal/application/dto"
	"context"
	"errors"
)

type UserRepository interface {
	CreateUserById(ctx context.Context, user dto.User) error
}

func (s *service) CreateUser(ctx context.Context, user dto.User) error {
	if err := validateBalance(user.Balance); err != nil {
		return err
	}

	if err := s.userRepository.CreateUserById(ctx, user); err != nil {
		return errors.New("cannot create a user")
	}

	return nil
}

func validateBalance(balance float64) error {
	if balance != 0 {
		return errors.New("incorrect balance")
	}

	return nil
}

func New(userRepository UserRepository) *service {
	return &service{
		userRepository: userRepository,
	}
}

type service struct {
	userRepository UserRepository
}
