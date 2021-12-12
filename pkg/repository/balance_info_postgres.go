package repository

import (
	"context"
	"database/sql"
	"tech_task"

	"github.com/sirupsen/logrus"
)

type BalanceInfoPostgres struct {
	db *sql.DB
}

func NewBalanceInfoPostgres(db *sql.DB) *BalanceInfoPostgres {
	return &BalanceInfoPostgres{db: db}
}

func (u *BalanceInfoPostgres) BalanceInfoDB(ctx context.Context, id int64) (int64, float64, error) {
	user := &tech_task.User{}
	err := u.db.QueryRow("SELECT id, balance FROM users WHERE id=$1;", id).Scan(&user.Id, &user.Balance)
	if err != nil {
		logrus.WithError(err).Errorf("user not found")
		return 0, 0, err
	}

	return user.Id, user.Balance, nil
}
