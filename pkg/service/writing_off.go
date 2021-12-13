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

func (w *WritingOffService) WritingOffUser(ctx context.Context, id int64, amount float64) (userID int64, amountWritingOff float64, err error) {
	return w.repo.WritingOffDB(ctx, id, amount)
}
