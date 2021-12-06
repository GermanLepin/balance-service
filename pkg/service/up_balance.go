package service

import (
	"context"
	"tech_task/pkg/repository"
)

type UpBalanceService struct {
	repo repository.UpBalance
}

func NewUpBalanceService(repo repository.UpBalance) *UpBalanceService {
	return &UpBalanceService{repo: repo}
}

func (u *UpBalanceService) UpBalanceUser(ctx context.Context, id int64, amount float64) (int64, float64, error) {
	return u.repo.UpBalanceDB(ctx, id, amount)
}
