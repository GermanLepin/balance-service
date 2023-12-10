package delete_user_service

import (
	"balance-service/internal/application/dto"
	"context"
	"fmt"

	"github.com/google/uuid"
)

type UserRepository interface {
	FetchUserById(ctx context.Context, userID uuid.UUID) (dto.User, error)
	DeleteUserById(ctx context.Context, userID uuid.UUID) error
}

func (s *service) DeleteUser(ctx context.Context, userID uuid.UUID) (dto.User, error) {
	user, err := s.validateLevel(ctx, userID)
	if err != nil {
		return user, err
	}

	if err := s.userRepository.DeleteUserById(ctx, userID); err != nil {
		return user, fmt.Errorf("cannot delete the user: %s", userID)
	}

	return user, nil
}

func (s *service) validateLevel(ctx context.Context, userID uuid.UUID) (dto.User, error) {
	user, err := s.userRepository.FetchUserById(ctx, userID)
	if err != nil {
		return user, fmt.Errorf("cannot fetch the user: %s", userID)
	}

	if user.Balance < 0 {
		return user, fmt.Errorf("cannot delete a user with a negative balance: %f", user.Balance)
	}

	if user.Balance > 0 {
		return user, fmt.Errorf("cannot delete a user with a possitive balance: %f", user.Balance)
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
