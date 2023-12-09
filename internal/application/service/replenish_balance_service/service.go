package replenish_balance_service

import (
	"balance-service/internal/application/dto"
	"fmt"

	"context"
)

type BalanceRepository interface {
	ReplenishBalance(ctx context.Context, replenishBalance dto.ReplenishBalanceRequest) (dto.User, error)
}

func (s *service) ReplenishBalance(ctx context.Context, replenishBalanceRequest dto.ReplenishBalanceRequest) (user dto.User, err error) {
	if err = validateAmount(replenishBalanceRequest.Amount); err != nil {
		return user, err
	}

	user, err = s.balanceRepository.ReplenishBalance(ctx, replenishBalanceRequest)
	if err != nil {
		return user, fmt.Errorf("cannot replenish a balance of user: %s", replenishBalanceRequest.UserID.String())
	}

	return user, nil
}

func validateAmount(amount float64) error {
	if amount < 0 {
		return fmt.Errorf("incorrect balance: %x", amount)
	}

	return nil
}

func New(
	balanceRepository BalanceRepository,
) *service {
	return &service{
		balanceRepository: balanceRepository,
	}
}

type service struct {
	balanceRepository BalanceRepository
}
