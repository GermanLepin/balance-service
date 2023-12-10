package deplete_balance_service

import (
	"balance-service/internal/application/dto"
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type (
	UserRepository interface {
		FetchUserById(ctx context.Context, userID uuid.UUID) (dto.User, error)
	}

	BalanceRepository interface {
		DepleteBalance(ctx context.Context, depleteBalance dto.DepleteBalanceRequest) error
	}
)

func (s *service) DepleteBalance(ctx context.Context, depleteBalance dto.DepleteBalanceRequest) (user dto.User, err error) {
	if err := s.validateLevel(ctx, depleteBalance.Amount, depleteBalance.UserID); err != nil {
		return user, err
	}

	err = s.balanceRepository.DepleteBalance(ctx, depleteBalance)
	if err != nil {
		return user, fmt.Errorf("cannot write off funds from the balance of the user: %s", depleteBalance.UserID.String())
	}

	user, err = s.userRepository.FetchUserById(ctx, depleteBalance.UserID)
	if err != nil {
		return user, fmt.Errorf("cannot fetch the user: %s", depleteBalance.UserID)
	}

	return user, nil
}

func (s *service) validateLevel(ctx context.Context, amount float32, userID uuid.UUID) error {
	user, err := s.userRepository.FetchUserById(ctx, userID)
	if err != nil {
		return fmt.Errorf("cannot fetch the user: %s", userID)
	}

	if amount < 0 {
		return fmt.Errorf("incorrect amount: %x", amount)
	}

	if user.Balance-amount < 0 {
		return errors.New("the writte-off amount cannot be more than the user's balance")
	}

	return nil
}

func New(
	balanceRepository BalanceRepository,
	userRepository UserRepository,
) *service {
	return &service{
		balanceRepository: balanceRepository,
		userRepository:    userRepository,
	}
}

type service struct {
	balanceRepository BalanceRepository
	userRepository    UserRepository
}
