package repository

import (
	"context"
	"database/sql"
)

type WritingOffPostgres struct {
	db *sql.DB
}

func NewWritingOffPostgres(db *sql.DB) *WritingOffPostgres {
	return &WritingOffPostgres{db: db}
}

func (w *WritingOffPostgres) WritingOffDB(ctx context.Context, id int64, amount float64) (int64, float64, error) {
	_, err := w.db.Exec("UPDATE users SET balance=balance-$1 WHERE id=$2;", amount, id)
	if err != nil {
		return 0, 0, err
	}

	return id, amount, nil
}
