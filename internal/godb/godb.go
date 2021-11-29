package godb

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"tech_task/internal/entities"
	"tech_task/pkg/helpers/jsonenc"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
)

type Instance struct {
	Db *pgxpool.Pool
}

func (i *Instance) UpBalanceDB(ctx context.Context, id int64, amount float64) int64 {
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

func (i *Instance) AddDescriptionDB(ctx context.Context, id int64, balanceAtMoment, corectAmount float64, refill, description, senderReceiver string) {
	_, err := i.Db.Exec(ctx, "INSERT INTO description (created_at, description, sender_receiver, balance_at_moment, amount, refill, user_id) VALUES ($1, $2, $3, $4, $5, $6, $7)", time.Now(), description, senderReceiver, balanceAtMoment, corectAmount, refill, id)
	if err != nil {
		fmt.Println(err)
	}
}

func (i *Instance) GetUserIdDB(ctx context.Context, w http.ResponseWriter, userId int64) (result []entities.Description) {
	var descript []entities.Description

	rows, err := i.Db.Query(ctx, "SELECT id_description, sender_receiver, amount, description, balance_at_moment, user_id, created_at, refill FROM description WHERE user_id=$1;", userId)
	if err == pgx.ErrNoRows {
		w.WriteHeader(http.StatusBadRequest)
		log.WithError(err).Errorf("No rows")
		jsonenc.JSONError(w, "No rows")
		return
	} else if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.WithError(err).Errorf("Syntax error SQL")
		jsonenc.JSONError(w, "Syntax error SQL")
		return
	}
	defer rows.Close()

	for rows.Next() {
		description := entities.Description{}
		rows.Scan(&description.Id, &description.SenderReceiver, &description.Amount, &description.Description, &description.BalanceAtMoment, &description.UserId, &description.CreatedAt, &description.Refil)
		descript = append(descript, description)
	}

	if descript == nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Errorf("User not found in databas")
		jsonenc.JSONError(w, "User not found in databas")
		return
	}

	return descript
}

func (i *Instance) GetUserIdSortDB(ctx context.Context, w http.ResponseWriter, userId int64, paramsSort string) (result []entities.Description) {
	var descript []entities.Description

	rows, err := i.Db.Query(ctx, fmt.Sprintf("SELECT id_description, sender_receiver, amount, description, balance_at_moment, user_id, created_at, refill FROM description WHERE user_id=$1 ORDER BY %s;", paramsSort), userId)
	if err == pgx.ErrNoRows {
		w.WriteHeader(http.StatusBadRequest)
		log.WithError(err).Errorf("No rows")
		jsonenc.JSONError(w, "No rows")
		return
	} else if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.WithError(err).Errorf("Syntax error SQL")
		jsonenc.JSONError(w, "Syntax error SQL")
		return
	}
	defer rows.Close()

	for rows.Next() {
		description := entities.Description{}
		rows.Scan(&description.Id, &description.SenderReceiver, &description.Amount, &description.Description, &description.BalanceAtMoment, &description.UserId, &description.CreatedAt, &description.Refil)
		descript = append(descript, description)
	}

	if descript == nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Errorf("User not found in databas")
		jsonenc.JSONError(w, "User not found in databas")
		return
	}

	return descript
}

func (i *Instance) GetUserIdSortDescDB(ctx context.Context, w http.ResponseWriter, userId int64, paramsSort string) (result []entities.Description) {
	var descript []entities.Description

	rows, err := i.Db.Query(ctx, fmt.Sprintf("SELECT id_description, sender_receiver, amount, description, balance_at_moment, user_id, created_at, refill FROM description WHERE user_id=$1 ORDER BY %s DESC;", paramsSort), userId)
	if err == pgx.ErrNoRows {
		w.WriteHeader(http.StatusBadRequest)
		log.WithError(err).Errorf("No rows")
		jsonenc.JSONError(w, "No rows")
		return
	} else if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.WithError(err).Errorf("Syntax error SQL")
		jsonenc.JSONError(w, "Syntax error SQL")
		return
	}
	defer rows.Close()

	for rows.Next() {
		description := entities.Description{}
		rows.Scan(&description.Id, &description.SenderReceiver, &description.Amount, &description.Description, &description.BalanceAtMoment, &description.UserId, &description.CreatedAt, &description.Refil)
		descript = append(descript, description)
	}

	if descript == nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Errorf("User not found in databas")
		jsonenc.JSONError(w, "User not found in databas")
		return
	}

	return descript
}

func (i *Instance) GetAllUsersDB(ctx context.Context, w http.ResponseWriter) (result []entities.Description) {
	var descript []entities.Description

	rows, err := i.Db.Query(ctx, "SELECT id_description, sender_receiver, amount, description, balance_at_moment, user_id, created_at, refill FROM description;")
	if err == pgx.ErrNoRows {
		w.WriteHeader(http.StatusBadRequest)
		log.WithError(err).Errorf("No rows")
		jsonenc.JSONError(w, "No rows")
		return
	} else if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.WithError(err).Errorf("Syntax error SQL")
		jsonenc.JSONError(w, "Syntax error SQL")
		return
	}

	for rows.Next() {
		description := entities.Description{}
		rows.Scan(&description.Id, &description.SenderReceiver, &description.Amount, &description.Description, &description.BalanceAtMoment, &description.UserId, &description.CreatedAt, &description.Refil)
		descript = append(descript, description)
	}

	defer rows.Close()

	if descript == nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Errorf("User not found in databas")
		jsonenc.JSONError(w, "User not found in databas")
		return
	}

	return descript
}

func (i *Instance) GetAllUsersSortDB(ctx context.Context, w http.ResponseWriter, params string) (result []entities.Description) {
	var descript []entities.Description

	rows, err := i.Db.Query(ctx, fmt.Sprintf("SELECT id_description, sender_receiver, amount, description, balance_at_moment, user_id, created_at, refill FROM description ORDER BY %s;", params))
	if err == pgx.ErrNoRows {
		w.WriteHeader(http.StatusBadRequest)
		log.WithError(err).Errorf("No rows")
		jsonenc.JSONError(w, "No rows")
		return
	} else if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.WithError(err).Errorf("Syntax error SQL")
		jsonenc.JSONError(w, "Syntax error SQL")
		return
	}

	for rows.Next() {
		description := entities.Description{}
		rows.Scan(&description.Id, &description.SenderReceiver, &description.Amount, &description.Description, &description.BalanceAtMoment, &description.UserId, &description.CreatedAt, &description.Refil)
		descript = append(descript, description)
	}

	defer rows.Close()

	if descript == nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Errorf("User not found in databas")
		jsonenc.JSONError(w, "User not found in databas")
		return
	}

	return descript
}

func (i *Instance) GetAllUsersSortDescDB(ctx context.Context, w http.ResponseWriter, params string) (result []entities.Description) {
	var descript []entities.Description

	rows, err := i.Db.Query(ctx, fmt.Sprintf("SELECT id_description, sender_receiver, amount, description, balance_at_moment, user_id, created_at, refill FROM description ORDER BY %s DESC;", params))
	if err == pgx.ErrNoRows {
		w.WriteHeader(http.StatusBadRequest)
		log.WithError(err).Errorf("No rows")
		jsonenc.JSONError(w, "No rows")
		return
	} else if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.WithError(err).Errorf("Syntax error SQL")
		jsonenc.JSONError(w, "Syntax error SQL")
		return
	}
	defer rows.Close()

	for rows.Next() {
		description := entities.Description{}
		rows.Scan(&description.Id, &description.SenderReceiver, &description.Amount, &description.Description, &description.BalanceAtMoment, &description.UserId, &description.CreatedAt, &description.Refil)
		descript = append(descript, description)
	}

	if descript == nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Errorf("User not found in databas")
		jsonenc.JSONError(w, "User not found in databas")
		return
	}

	return descript
}
