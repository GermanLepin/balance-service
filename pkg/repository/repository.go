package repository

import (
	"context"
	"tech_task"

	"github.com/jackc/pgx/v4/pgxpool"
)

type UpBalance interface {
	UpBalanceDB(context.Context, int64, float64) (int64, float64, error)
}

type BalanceInfo interface {
	BalanceInfoDB(context.Context, int64) (int64, float64, error)
}

type WritingOff interface {
	WritingOffDB(context.Context, int64, float64) (int64, float64, error)
}

type AddDescription interface {
	AddDescriptionDB(context.Context, int64, float64, float64, string, string, string) error
}

type GetAllUsersDescriptionsSort interface {
	GetAllUsersDescriptionsSortDB(context.Context, string, string, string) ([]tech_task.Description, error)
}

type GetUserIdDescriptionsSort interface {
	GetUserIdDescriptionsSortDB(context.Context, int64, string, string, string) ([]tech_task.Description, error)
}

type Repository struct {
	UpBalance
	BalanceInfo
	WritingOff
	AddDescription
	GetAllUsersDescriptionsSort
	GetUserIdDescriptionsSort
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		UpBalance:                   NewUpBalancePostgres(db),
		BalanceInfo:                 NewBalanceInfoPostgres(db),
		WritingOff:                  NewWritingOffPostgres(db),
		AddDescription:              NewAddDescriptionPostgres(db),
		GetAllUsersDescriptionsSort: NewGetAllUsersDescriptionsSortPostgres(db),
		GetUserIdDescriptionsSort:   NewGetUserIdDescriptionsPostgres(db),
	}
}
