package delete_user_service

import (
	"balance-service/internal/application/dto"
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type UserRepository interface {
	FetchUserById(ctx context.Context, userID uuid.UUID) (dto.User, error)
	DeleteUserById(ctx context.Context, userID uuid.UUID) error
}

func (s *service) DeleteUser(ctx context.Context, userID uuid.UUID) (dto.User, error) {
	user, err := s.userRepository.FetchUserById(ctx, userID)
	if err != nil {
		return user, fmt.Errorf("cannot fetch a user: user id %s", userID)
	}

	if user.Balance < 0 {
		return user, fmt.Errorf("cannot delete a user with negative balance: %x", user.Balance)
	}

	if err := s.userRepository.DeleteUserById(ctx, userID); err != nil {
		return user, errors.New("cannot delete a user")
	}

	return user, nil
}

func New(userRepository UserRepository) *service {
	return &service{
		userRepository: userRepository,
	}
}

type service struct {
	userRepository UserRepository
}
