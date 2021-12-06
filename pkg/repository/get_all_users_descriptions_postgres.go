package repository

import (
	"context"
	"errors"
	"tech_task"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

type GetAllUsersDescriptionsPostgres struct {
	db *pgxpool.Pool
}

func NewGetAllUsersDescriptionsPostgres(db *pgxpool.Pool) *GetAllUsersDescriptionsPostgres {
	return &GetAllUsersDescriptionsPostgres{db: db}
}

func (g *GetAllUsersDescriptionsPostgres) GetAllUsersDescriptionsDB(ctx context.Context) ([]tech_task.Description, error) {
	var descriptions []tech_task.Description

	rows, err := g.db.Query(ctx, "SELECT id_description, sender_receiver, amount, description, balance_at_moment, user_id, created_at, refill FROM description;")
	if err != nil {
		logrus.WithError(err).Errorf("Syntax error SQL")
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
		logrus.Errorf("Users not found in databas")
		return nil, errors.New("users not found in databas")
	}

	return descriptions, nil
}
