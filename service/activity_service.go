package service

import (
	"context"
	"github.com/MCPutro/golang-todo/model"
)

type ActivityService interface {
	Create(ctx context.Context, req *model.Activities) (*model.Activities, error)
	Update(ctx context.Context, req *model.Activities) (*model.Activities, error)
	FindAll(ctx context.Context) ([]*model.Activities, error)
	FindById(ctx context.Context, Id int) (*model.Activities, error)
	Delete(ctx context.Context, id int) error
}
