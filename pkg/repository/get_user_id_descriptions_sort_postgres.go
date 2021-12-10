package repository

import (
	"context"
	"fmt"
	"tech_task"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

type GetUserIdDescriptionsPostgres struct {
	db *pgxpool.Pool
}

func NewGetUserIdDescriptionsPostgres(db *pgxpool.Pool) *GetUserIdDescriptionsPostgres {
	return &GetUserIdDescriptionsPostgres{db: db}
}

func (g *GetUserIdDescriptionsPostgres) GetUserIdDescriptionsSortDB(ctx context.Context, userId int64, sortParams, orderBy, sqlOrderBy string) ([]tech_task.Description, error) {
	var description []tech_task.Description
	rows, err := g.db.Query(ctx, fmt.Sprintf("SELECT id_description, sender_receiver, amount, description, balance_at_moment, user_id, created_at, refill FROM description WHERE user_id=$1 %s %s %s;", sqlOrderBy, sortParams, orderBy), userId)
	if err != nil {
		logrus.WithError(err).Errorf("Syntax error SQL")
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		descript := tech_task.Description{}
		rows.Scan(&descript.Id,
			&descript.SenderReceiver,
			&descript.Amount,
			&descript.Description,
			&descript.BalanceAtMoment,
			&descript.UserId,
			&descript.CreatedAt,
			&descript.Refil)
		description = append(description, descript)
	}

	if description == nil {
		logrus.Errorf("user not found")
		return nil, err
	}

	return description, nil
}
