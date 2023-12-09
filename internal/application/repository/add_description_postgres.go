package repository

// import (
// 	"context"
// 	"database/sql"
// 	"time"

// 	"github.com/sirupsen/logrus"
// )

// func (a *AddDescriptionPostgres) AddDescriptionDB(ctx context.Context, id int64, balanceAtMoment, corectAmount float64, refill, description, senderReceiver string) error {
// 	_, err := a.db.Exec("insert into descriptions (created_at, description, sender_receiver, balance_at_moment, amount, refill, user_id) VALUES ($1, $2, $3, $4, $5, $6, $7)",
// 		time.Now(),
// 		description,
// 		senderReceiver,
// 		balanceAtMoment,
// 		corectAmount,
// 		refill,
// 		id,
// 	)
// 	if err != nil {
// 		logrus.WithError(err).Errorf("user not found")
// 		return err
// 	}

// 	return nil
// }

// type AddDescriptionPostgres struct {
// 	db *sql.DB
// }

// func NewAddDescriptionPostgres(db *sql.DB) *AddDescriptionPostgres {
// 	return &AddDescriptionPostgres{db: db}
// }
