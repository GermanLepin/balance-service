package service

import (
	"context"
	"tech_task"
	"tech_task/pkg/repository"
)

type GetDescriptionsService struct {
	repo repository.GetDescriptions
}

func NewGetDescriptionsService(repo repository.GetDescriptions) *GetDescriptionsService {
	return &GetDescriptionsService{repo: repo}
}

func (g *GetDescriptionsService) GetDescriptionsUsers(ctx context.Context, id int64, sortBy, orderBy string) (descriptionsList []tech_task.Description, err error) {
	return g.repo.GetDescriptionsDB(ctx, id, sortBy, orderBy)
}
