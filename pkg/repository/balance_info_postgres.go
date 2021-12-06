package repository

import (
	"context"
	"tech_task"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

type BalanceInfoPostgres struct {
	db *pgxpool.Pool
}

func NewBalanceInfoPostgres(db *pgxpool.Pool) *BalanceInfoPostgres {
	return &BalanceInfoPostgres{db: db}
}

func (u *BalanceInfoPostgres) BalanceInfoDB(ctx context.Context, id int64) (int64, float64, error) {
	user := &tech_task.User{}
	err := u.db.QueryRow(ctx, "SELECT id, balance FROM users WHERE id=$1;", id).Scan(&user.Id, &user.Balance)
	if err != nil {
		logrus.WithError(err).Errorf("User not found in database")
		return 0, 0, err
	}

	return user.Id, user.Balance, nil
}
