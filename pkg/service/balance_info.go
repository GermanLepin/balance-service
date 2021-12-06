package service

import (
	"context"
	"tech_task/pkg/repository"
)

type BalanceInfoService struct {
	repo repository.BalanceInfo
}

func NewBalanceInfoService(repo repository.BalanceInfo) *BalanceInfoService {
	return &BalanceInfoService{repo: repo}
}

func (b *BalanceInfoService) BalanceInfoUser(ctx context.Context, id int64) (int64, float64, error) {
	return b.repo.BalanceInfoDB(ctx, id)
}
