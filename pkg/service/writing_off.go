package service

import (
	"context"
	"tech_task/pkg/repository"
)

type WritingOffService struct {
	repo repository.WritingOff
}

func NewWritingOffService(repo repository.WritingOff) *WritingOffService {
	return &WritingOffService{repo: repo}
}

func (w *WritingOffService) WritingOffUser(ctx context.Context, id int64, amount float64) (int64, float64, error) {
	return w.repo.WritingOffDB(ctx, id, amount)
}
