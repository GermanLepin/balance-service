package service

import (
	"context"
	"tech_task/pkg/repository"
)

type AddDescriptionService struct {
	repo repository.AddDescription
}

func NewAddDescriptionService(repo repository.AddDescription) *AddDescriptionService {
	return &AddDescriptionService{repo: repo}
}

func (a *AddDescriptionService) AddDescriptionUser(ctx context.Context, id int64, balanceAtMoment, corectAmount float64, refill, description, senderReceiver string) error {
	return a.repo.AddDescriptionDB(ctx, id, balanceAtMoment, corectAmount, refill, description, senderReceiver)
}
