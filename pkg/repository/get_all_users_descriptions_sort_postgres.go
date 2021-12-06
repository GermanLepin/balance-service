package repository

import (
	"context"
	"errors"
	"fmt"
	"tech_task"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

type GetAllUsersDescriptionsSortPostgres struct {
	db *pgxpool.Pool
}

func NewGetAllUsersDescriptionsSortPostgres(db *pgxpool.Pool) *GetAllUsersDescriptionsSortPostgres {
	return &GetAllUsersDescriptionsSortPostgres{db: db}
}

func (g *GetAllUsersDescriptionsSortPostgres) GetAllUsersDescriptionsSortDB(ctx context.Context, sortParams, orderBy, sqlOrderBy string) ([]tech_task.Description, error) {
	var descriptions []tech_task.Description
	rows, err := g.db.Query(ctx, fmt.Sprintf("SELECT id_description, sender_receiver, amount, description, balance_at_moment, user_id, created_at, refill FROM description %s %s %s;", sqlOrderBy, sortParams, orderBy))
	if err != nil {
		logrus.WithError(err).Errorf(err.Error())
		return nil, err
	}

	for rows.Next() {
		description := tech_task.Description{}
		rows.Scan(&description.Id,
			&description.SenderReceiver,
			&description.Amount,
			&description.Description,
			&description.BalanceAtMoment,
			&description.UserId,
			&description.CreatedAt,
			&description.Refil)
		descriptions = append(descriptions, description)
	}
	defer rows.Close()

	if descriptions == nil {
		logrus.Errorf("user not found in database")
		return nil, errors.New("user not found in database")
	}

	return descriptions, nil
}
