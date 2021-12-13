package service

import (
	"context"
	"sync"
	"tech_task/pkg/repository"
)

type BalanceInfoService struct {
	mu   sync.Mutex
	repo repository.BalanceInfo
}

func NewBalanceInfoService(repo repository.BalanceInfo) *BalanceInfoService {
	return &BalanceInfoService{repo: repo}
}

func (b *BalanceInfoService) BalanceInfoUser(ctx context.Context, id int64) (userID int64, balance float64, err error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.repo.BalanceInfoDB(ctx, id)
}
