package service

import (
	"context"
	"tech_task"
	"tech_task/pkg/repository"
)

type UpBalance interface {
	UpBalanceUser(context.Context, int64, float64) error
}

type BalanceInfo interface {
	BalanceInfoUser(context.Context, int64) (int64, float64, error)
}

type WritingOff interface {
	WritingOffUser(context.Context, int64, float64) (int64, float64, error)
}

type AddDescription interface {
	AddDescriptionUser(context.Context, int64, float64, float64, string, string, string) error
}

type GetDescriptions interface {
	GetDescriptionsUsers(context.Context, int64, string, string) ([]tech_task.Description, error)
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
