package godb

import (
	"context"
	"fmt"
	"net/http"

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
	user_err := i.Db.QueryRow(ctx, "SELECT id, balance FROM users WHERE id=$1;", id).Scan(&user.Id, &user.Balance)
	if user_err != nil {
		_, err := i.Db.Exec(ctx, "INSERT INTO users (id, balance) VALUES ($1, $2)", id, amount)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		_, user_err := i.Db.Exec(ctx, "UPDATE users SET balance=balance+$1 WHERE id=$2;", amount, id)
		if user_err != nil {
			fmt.Println(user_err)
		}
	}

	return user.Id
}

func (i *Instance) BalanceInfoDB(ctx context.Context, w http.ResponseWriter, id int64) (user_id int64, balance float64) {
	user := &entities.User{}
	user_err := i.Db.QueryRow(ctx, "SELECT id, balance FROM users WHERE id=$1;", id).Scan(&user.Id, &user.Balance)
	if user_err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.WithError(user_err).Errorf("User not found in database")
		jsonenc.JSONError(w, "User not found in database")
		return
	}

	return user.Id, user.Balance
}

func (i *Instance) WritingOffDB(ctx context.Context, id int64, amount float64) {
	_, user_err := i.Db.Exec(ctx, "UPDATE users SET balance=balance-$1 WHERE id=$2;", amount, id)
	if user_err != nil {
		fmt.Println(user_err)
		return
	}
}

func (i *Instance) DeleteUserDB(ctx context.Context, id int64) {
	_, user_err := i.Db.Exec(ctx, "DELETE FROM users WHERE id=$1;", id)
	if user_err != nil {
		fmt.Println(user_err)
		return
	}
}
