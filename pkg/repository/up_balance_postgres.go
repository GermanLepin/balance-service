package repository

import (
	"context"
	"tech_task"

	"github.com/jackc/pgx/v4/pgxpool"
)

type UpBalancePostgres struct {
	db *pgxpool.Pool
}

func NewUpBalancePostgres(db *pgxpool.Pool) *UpBalancePostgres {
	return &UpBalancePostgres{db: db}
}

func (u *UpBalancePostgres) UpBalanceDB(ctx context.Context, id int64, amount float64) (int64, float64, error) {
	user := &tech_task.User{}
	userErr := u.db.QueryRow(ctx, "SELECT id, balance FROM users WHERE id=$1;", id).Scan(&user.Id, &user.Balance)
	if userErr != nil {
		_, err := u.db.Exec(ctx, "INSERT INTO users (id, balance) VALUES ($1, $2)", id, amount)
		if err != nil {
			return 0, 0, err
		}
	} else {
		_, err := u.db.Exec(ctx, "UPDATE users SET balance=balance+$1 WHERE id=$2;", amount, id)
		if err != nil {
			return 0, 0, err
		}
	}

	return user.Id, amount, nil
}
