package replenish_balance_service

import (
	"balance-service/internal/application/dto"
	"fmt"

	"context"

	"github.com/google/uuid"
)

type (
	UserRepository interface {
		FetchUserById(ctx context.Context, userID uuid.UUID) (dto.User, error)
	}

	BalanceRepository interface {
		ReplenishBalance(ctx context.Context, replenishBalance dto.ReplenishBalanceRequest) error
	}
)

func (s *service) ReplenishBalance(ctx context.Context, replenishBalance dto.ReplenishBalanceRequest) (user dto.User, err error) {
	if err = validateAmount(replenishBalance.Amount); err != nil {
		return user, err
	}

	err = s.balanceRepository.ReplenishBalance(ctx, replenishBalance)
	if err != nil {
		return user, fmt.Errorf("cannot top up the balance of the user: %s", replenishBalance.UserID.String())
	}

	user, err = s.userRepository.FetchUserById(ctx, replenishBalance.UserID)
	if err != nil {
		return user, fmt.Errorf("cannot fetch the user: %s", replenishBalance.UserID)
	}

	return user, nil
}

func validateAmount(amount float32) error {
	if amount < 0 {
		return fmt.Errorf("incorrect amount: %f", amount)
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
