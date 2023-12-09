package create_user_service

import (
	"balance-service/internal/application/dto"
	"context"
	"errors"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user dto.User) error
}

func (s *service) CreateUser(ctx context.Context, user dto.User) error {
	if err := s.userRepository.CreateUser(ctx, user); err != nil {
		return errors.New("cannot create a user")
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
