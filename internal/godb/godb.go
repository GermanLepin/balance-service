package godb

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"tech_task/internal/entities"
	"tech_task/pkg/helpers/jsonenc.go"

	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
)

type Instance struct {
	Db *pgxpool.Pool
}

func (i *Instance) UpBalanceDB(ctx context.Context, w http.ResponseWriter, id int64, amount float64) int64 {
	user := &entities.User{}
	userErr := i.Db.QueryRow(ctx, "SELECT id, balance FROM users WHERE id=$1;", id).Scan(&user.Id, &user.Balance)
	if userErr != nil {
		_, err := i.Db.Exec(ctx, "INSERT INTO users (id, balance) VALUES ($1, $2)", id, amount)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		_, userErr := i.Db.Exec(ctx, "UPDATE users SET balance=balance+$1 WHERE id=$2;", amount, id)
		if userErr != nil {
			fmt.Println(userErr)
		}
	}

	return user.Id
}

func (i *Instance) BalanceInfoDB(ctx context.Context, w http.ResponseWriter, id int64) (userId int64, balance float64) {
	user := &entities.User{}
	userErr := i.Db.QueryRow(ctx, "SELECT id, balance FROM users WHERE id=$1;", id).Scan(&user.Id, &user.Balance)
	if userErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.WithError(userErr).Errorf("User not found in database")
		jsonenc.JSONError(w, "User not found in database")
		return
	}

	return user.Id, user.Balance
}

func (i *Instance) WritingOffDB(ctx context.Context, id int64, amount float64) {
	_, userErr := i.Db.Exec(ctx, "UPDATE users SET balance=balance-$1 WHERE id=$2;", amount, id)
	if userErr != nil {
		fmt.Println(userErr)
	}
}

func (i *Instance) DeleteUserDB(ctx context.Context, id int64) {
	_, userErr := i.Db.Exec(ctx, "DELETE FROM users WHERE id=$1;", id)
	if userErr != nil {
		fmt.Println(userErr)
	}
}

func (i *Instance) AddDiscriptionDB(ctx context.Context, w http.ResponseWriter, id int64, balanceAtMoment, corectAmount float64, refill, description, senderReceiver string) {
	_, err := i.Db.Exec(ctx, "INSERT INTO description (created_at, description, sender_receiver, balance_at_moment, amount, refill, userId) VALUES ($1, $2, $3, $4, $5, $6, $7)", time.Now(), description, senderReceiver, balanceAtMoment, corectAmount, refill, id)
	if err != nil {
		fmt.Println(err)
	}
}
