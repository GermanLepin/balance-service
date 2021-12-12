package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"tech_task"

	sq "github.com/Masterminds/squirrel"
	"github.com/sirupsen/logrus"
)

type GetDescriptionsPostgres struct {
	db *sql.DB
}

func NewGetDescriptionsPostgres(db *sql.DB) *GetDescriptionsPostgres {
	return &GetDescriptionsPostgres{db: db}
}

func (gp *GetDescriptionsPostgres) GetDescriptionsDB(ctx context.Context, uid int64, sortBy, orderBy string) ([]tech_task.Description, error) {
	var descriptions []tech_task.Description

	baseQuery := sq.Select(`
		id_description,
		sender_receiver,
		amount,
		description,
		balance_at_moment,
		user_id,
		created_at,
		refill
		`).
		From("description")

	if uid != 0 {
		baseQuery = baseQuery.Where(sq.Eq{"user_id": uid})
	}

	switch {
	case sortBy != "" && orderBy != "":
		params := fmt.Sprintf(sortBy + " " + orderBy)
		baseQuery = baseQuery.OrderBy(params)
	case sortBy != "":
		params := fmt.Sprintf(sortBy + " ASC")
		baseQuery = baseQuery.OrderBy(params)
	default:
		baseQuery = baseQuery.OrderBy("created_at ASC")
	}



	rows, err := baseQuery.RunWith(gp.db).QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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
		logrus.Errorf("user not found")
		return nil, errors.New("user not found")
	}

	return descriptions, nil
}
