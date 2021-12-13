package service

import (
	"context"
	"tech_task"
	"tech_task/pkg/repository"
)

type UpBalance interface {
	UpBalanceUser(ctx context.Context, uid int64, amount float64) error
}

type BalanceInfo interface {
	BalanceInfoUser(ctx context.Context, uid int64) (userID int64, balance float64, err error)
}

type WritingOff interface {
	WritingOffUser(ctx context.Context, uid int64, amount float64) (userID int64, amountWritingOff float64, err error)
}

type AddDescription interface {
	AddDescriptionUser(ctx context.Context, uid int64, balanceAtMoment float64, correctAmount float64, refill string, description string, senderReceiver string) error
}

type GetDescriptions interface {
	GetDescriptionsUsers(ctx context.Context, uid int64, sortBy string, orderBy string) (descriptionsList []tech_task.Description, err error)
}

type Service struct {
	UpBalance
	BalanceInfo
	WritingOff
	AddDescription
	GetDescriptions
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		UpBalance:       NewUpBalanceService(repos.UpBalance),
		BalanceInfo:     NewBalanceInfoService(repos.BalanceInfo),
		WritingOff:      NewWritingOffService(repos.WritingOff),
		AddDescription:  NewAddDescriptionService(repos.AddDescription),
		GetDescriptions: NewGetDescriptionsService(repos.GetDescriptions),
	}
}
