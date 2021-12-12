package service

import (
	"context"
	"sync"
	"tech_task/pkg/repository"
)

type WritingOffService struct {
	repo repository.WritingOff
	mu   sync.Mutex
}

func NewWritingOffService(repo repository.WritingOff) *WritingOffService {
	return &WritingOffService{repo: repo}
}

func (w *WritingOffService) WritingOffUser(ctx context.Context, id int64, amount float64) (int64, float64, error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.repo.WritingOffDB(ctx, id, amount)
}
