package service

import (
	"context"
	"tech_task"
	"tech_task/pkg/repository"
)

type GetAllUsersDescriptionsSortService struct {
	repo repository.GetAllUsersDescriptionsSort
}

func NewGetAllUsersDescriptionsSortService(repo repository.GetAllUsersDescriptionsSort) *GetAllUsersDescriptionsSortService {
	return &GetAllUsersDescriptionsSortService{repo: repo}
}

func (g *GetAllUsersDescriptionsSortService) GetAllDescriptionsSort(ctx context.Context, params string, orderBy string) ([]tech_task.Description, error) {
	return g.repo.GetAllUsersDescriptionsSortDB(ctx, params, orderBy)
}
