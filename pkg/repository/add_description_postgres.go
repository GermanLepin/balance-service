package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

type AddDescriptionPostgres struct {
	db *pgxpool.Pool
}

func NewAddDescriptionPostgres(db *pgxpool.Pool) *AddDescriptionPostgres {
	return &AddDescriptionPostgres{db: db}
}

func (a *AddDescriptionPostgres) AddDescriptionDB(ctx context.Context, id int64, balanceAtMoment, corectAmount float64, refill, description, senderReceiver string) error {
	_, err := a.db.Exec(ctx, "INSERT INTO description (created_at, description, sender_receiver, balance_at_moment, amount, refill, user_id) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		time.Now(), description, senderReceiver, balanceAtMoment, corectAmount, refill, id)
	if err != nil {
		logrus.WithError(err).Errorf("user not found")
		return err
	}

	return nil
}
