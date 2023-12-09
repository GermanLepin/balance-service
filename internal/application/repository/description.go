package repository

import (
	"database/sql"
)

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

// func (gp *GetDescriptionsPostgres) GetDescriptionsDB(ctx context.Context, uid int64, sortBy, orderBy string) (descriptionsList []tech_task.Description, err error) {
// 	var descriptions []tech_task.Description

// 	baseQuery := sq.Select(`
// 		id_description,
// 		sender_receiver,
// 		amount,
// 		description,
// 		balance_at_moment,
// 		user_id,
// 		created_at,
// 		refill
// 		`).
// 		From("descriptions")

// 	if uid != 0 {
// 		baseQuery = baseQuery.Where("user_id = $1", uid)
// 	}

// 	switch {
// 	case sortBy != "" && orderBy == "desc":
// 		baseQuery = baseQuery.OrderBy(sortBy + " " + orderBy)
// 	case sortBy != "":
// 		baseQuery = baseQuery.OrderBy(sortBy + " ASC")
// 	default:
// 		baseQuery = baseQuery.OrderBy("created_at ASC")
// 	}

// 	rows, err := baseQuery.RunWith(gp.db).QueryContext(ctx)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		description := tech_task.Description{}
// 		rows.Scan(&description.Id,
// 			&description.SenderReceiver,
// 			&description.Amount,
// 			&description.Description,
// 			&description.BalanceAtMoment,
// 			&description.UserID,
// 			&description.CreatedAt,
// 			&description.Refill)
// 		descriptions = append(descriptions, description)
// 	}

// 	if descriptions == nil {
// 		logrus.Errorf("user not found")
// 		return nil, errors.New("user not found")
// 	}

// 	return descriptions, nil
// }

type descriptionRepository struct {
	db *sql.DB
}

func NewDescriptionRepository(db *sql.DB) *descriptionRepository {
	return &descriptionRepository{
		db: db,
	}
}
