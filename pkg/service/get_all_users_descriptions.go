package service

import (
	"context"
	"tech_task"
	"tech_task/pkg/repository"
)

type GetAllUsersDescriptionsService struct {
	repo repository.GetAllUsersDescriptions
}

func NewGetAllUsersDescriptionsService(repo repository.GetAllUsersDescriptions) *GetAllUsersDescriptionsService {
	return &GetAllUsersDescriptionsService{repo: repo}
}

func (g *GetAllUsersDescriptionsService) GetAllDescriptions(ctx context.Context) ([]tech_task.Description, error) {
	return g.repo.GetAllUsersDescriptionsDB(ctx)
}
