package taskservice

import (
	"context"
	"github.com/rezaAmiri123/ormus/destination/entity"
)

type Service struct {
	repo Repository
}

type Repository interface {
	GetTaskByID(ctx context.Context, taskID string) (entity.Task, error)
}

func New(repo Repository) Service {
	return Service{repo: repo}

}
