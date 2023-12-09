package repository

import (
	"balance-service/internal/application/dto"
	"context"
	"database/sql"
)

func (b *balanceRepository) ReplenishBalance(ctx context.Context, replenishBalance dto.ReplenishBalanceRequest) (dto.User, error) {
	var user dto.User

	_, err := b.db.Exec("update service.users set balance=balance+$1 where id=$2;", replenishBalance.Amount, replenishBalance.UserID)
	if err != nil {
		return user, err
	}

	err = b.db.QueryRow("select * from service.users where id = $1;", replenishBalance.UserID).Scan(&user.ID, &user.Name, &user.Balance)
	if err != nil {
		return user, err
	}

	return user, nil
}

// func NewWritingOffPostgres(db *sql.DB) *WritingOffPostgres {
// 	return &WritingOffPostgres{db: db}
// }

// func (w *WritingOffPostgres) WritingOffDB(ctx context.Context, id int64, amount float64) (userID int64, amountWritingOff float64, err error) {
// 	_, err = w.db.Exec("UPDATE users SET balance=balance-$1 WHERE id=$2;", amount, id)
// 	if err != nil {
// 		return 0, 0, err
// 	}

// 	return id, amount, nil
// }

type balanceRepository struct {
	db *sql.DB
}

func NewBalanceRepository(db *sql.DB) *balanceRepository {
	return &balanceRepository{
		db: db,
	}
}
