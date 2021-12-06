package repository

import (
	"context"
	"database/sql"
	"tech_task"
)

type UpBalance interface {
	UpBalanceDB(context.Context, int64, float64) error
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

type GetDescriptions interface {
	GetDescriptionsDB(context.Context, int64, string, string) ([]tech_task.Description, error)
}

type Repository struct {
	UpBalance
	BalanceInfo
	WritingOff
	AddDescription
	GetDescriptions
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		UpBalance:       NewUpBalancePostgres(db),
		BalanceInfo:     NewBalanceInfoPostgres(db),
		WritingOff:      NewWritingOffPostgres(db),
		AddDescription:  NewAddDescriptionPostgres(db),
		GetDescriptions: NewGetDescriptionsPostgres(db),
	}
}
