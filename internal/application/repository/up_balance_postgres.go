package repository

// import (
// 	"context"
// 	"database/sql"
// 	"tech_task"
// )

// type UpBalancePostgres struct {
// 	db *sql.DB
// }

// func NewUpBalancePostgres(db *sql.DB) *UpBalancePostgres {
// 	return &UpBalancePostgres{db: db}
// }

// func (u *UpBalancePostgres) UpBalanceDB(ctx context.Context, id int64, amount float64) error {
// 	user := &tech_task.User{}
// 	userErr := u.db.QueryRow("SELECT id, balance FROM users WHERE id=$1;", id).Scan(&user.ID, &user.Balance)
// 	if userErr != nil {
// 		_, err := u.db.Exec("INSERT INTO users (id, balance) VALUES ($1, $2)", id, amount)
// 		if err != nil {
// 			return err
// 		}
// 	} else {
// 		_, err := u.db.Exec("UPDATE users SET balance=balance+$1 WHERE id=$2;", amount, id)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }
