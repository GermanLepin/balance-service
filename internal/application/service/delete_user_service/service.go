package delete_user_service

import (
	"balance-service/internal/application/dto"
	"context"
	"errors"

	"github.com/google/uuid"
)

type UserRepository interface {
	FetchUser(ctx context.Context, userID uuid.UUID) (dto.User, error)
	DeleteUser(ctx context.Context, userID uuid.UUID) error
}

func (s *service) DeleteUser(ctx context.Context, userID uuid.UUID) (dto.User, error) {
	user, err := s.userRepository.FetchUser(ctx, userID)
	if err != nil {
		return user, errors.New("cannot fetch a user")
	}

	if err := s.userRepository.DeleteUser(ctx, userID); err != nil {
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
