package repository

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type WritingOffPostgres struct {
	db *pgxpool.Pool
}

func NewWritingOffPostgres(db *pgxpool.Pool) *WritingOffPostgres {
	return &WritingOffPostgres{db: db}
}

func (w *WritingOffPostgres) WritingOffDB(ctx context.Context, id int64, amount float64) (int64, float64, error) {
	_, err := w.db.Exec(ctx, "UPDATE users SET balance=balance-$1 WHERE id=$2;", amount, id)
	if err != nil {
		return 0, 0, err
	}

	return id, amount, nil
}
