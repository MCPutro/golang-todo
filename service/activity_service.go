package service

import (
	"context"
	"github.com/MCPutro/golang-todo/model"
)

type ActivityService interface {
	Create(ctx context.Context, req *model.Activity) (*model.Activity, error)
	Update(ctx context.Context, req *model.Activity) (*model.Activity, error)
	FindAll(ctx context.Context) ([]*model.Activity, error)
	FindById(ctx context.Context, Id int) (*model.Activity, error)
	Delete(ctx context.Context, id int) error
}
