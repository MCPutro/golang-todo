package service

import (
	"context"
	"github.com/MCPutro/golang-todo/model"
)

type TodoService interface {
	Create(ctx context.Context, req *model.Todo) (*model.Todo, error)
	FindAll(ctx context.Context) ([]*model.Todo, error)
	FindById(ctx context.Context, Id int) (*model.Todo, error)
	FindByActivityID(ctx context.Context, Id int) ([]*model.Todo, error)
	Update(ctx context.Context, req *model.Todo) (*model.Todo, error)
	Delete(ctx context.Context, id int) error
}
