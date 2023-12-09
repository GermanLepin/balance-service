package fetch_balance_info_service

import (
	"balance-service/internal/application/dto"
	"context"
	"errors"

	"github.com/google/uuid"
)

type UserRepository interface {
	FetchUserById(ctx context.Context, userID uuid.UUID) (dto.User, error)
}

func (s *service) FetchBalanceInfo(ctx context.Context, userID uuid.UUID) (dto.User, error) {
	user, err := s.userRepository.FetchUserById(ctx, userID)
	if err != nil {
		return user, errors.New("cannot fetch a user")
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
