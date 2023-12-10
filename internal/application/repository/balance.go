package repository

import (
	"balance-service/internal/application/dto"
	"context"
	"database/sql"
)

func (b *balanceRepository) ReplenishBalance(ctx context.Context, replenishBalance dto.ReplenishBalanceRequest) error {
	_, err := b.db.Exec("update service.users set balance=balance+$1 where id=$2;", replenishBalance.Amount, replenishBalance.UserID)
	if err != nil {
		return err
	}

	return nil
}

func (b *balanceRepository) DepleteBalance(ctx context.Context, depleteBalance dto.DepleteBalanceRequest) error {
	_, err := b.db.Exec("update service.users set balance=balance-$1 where id=$2;", depleteBalance.Amount, depleteBalance.UserID)
	if err != nil {
		return err
	}

	return nil
}

type balanceRepository struct {
	db *sql.DB
}

func NewBalanceRepository(db *sql.DB) *balanceRepository {
	return &balanceRepository{
		db: db,
	}
}
