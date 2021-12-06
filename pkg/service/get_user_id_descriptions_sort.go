package service

import (
	"context"
	"tech_task"
	"tech_task/pkg/repository"
)

type GetUserIdDescriptionsSortService struct {
	repo repository.GetUserIdDescriptionsSort
}

func NewGetUserIdDescriptionsSortServiceService(repo repository.GetUserIdDescriptionsSort) *GetUserIdDescriptionsSortService {
	return &GetUserIdDescriptionsSortService{repo: repo}
}

func (g *GetUserIdDescriptionsSortService) GetUserIdDescriptionsSort(ctx context.Context, id int64, orderBy string) ([]tech_task.Description, error) {
	return g.repo.GetUserIdDescriptionsSortDB(ctx, id, orderBy)
}
