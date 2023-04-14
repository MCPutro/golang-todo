package service

import (
	"context"
	"database/sql"
	"github.com/MCPutro/golang-todo/helper"
	"github.com/MCPutro/golang-todo/model"
	"github.com/MCPutro/golang-todo/repository"
	"sync"
	"time"
)

var once sync.Once

type todoServiceImpl struct {
	repo    repository.TodoRepository
	actRepo repository.ActivityRepository
	db      *sql.DB
}

func NewTodoService(repo repository.TodoRepository, actRepo repository.ActivityRepository, db *sql.DB) TodoService {
	return &todoServiceImpl{repo: repo, actRepo: actRepo, db: db}
}

func (t *todoServiceImpl) Create(ctx context.Context, req *model.Todo) (*model.Todo, error) {
	//begin db transaction
	tx, err := t.db.Begin()
	if err != nil {
		return nil, err
	}

	//run commit or rollback in last func
	defer func() { helper.CommitOrRollback(err, tx) }()

	//conplete data in req
	now := time.Now().UTC()
	req.Created_at = now
	req.Updated_at = now
	if req.Priority == "" {
		req.Priority = "very-high"
	}

	//call repo
	save, err := t.repo.Save(ctx, tx, req)
	if err != nil {
		return nil, err
	}

	return save, nil
}

func (t *todoServiceImpl) FindAll(ctx context.Context) ([]*model.Todo, error) {
	//begin db transaction
	tx, err := t.db.Begin()
	if err != nil {
		return nil, err
	}

	//run commit or rollback in last func
	defer func() { helper.CommitOrRollback(err, tx) }()

	//call func from repo
	todos, err := t.repo.FindAll(ctx, tx)

	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (t *todoServiceImpl) FindById(ctx context.Context, Id int) (*model.Todo, error) {
	//begin db transaction
	tx, err := t.db.Begin()
	if err != nil {
		return nil, err
	}

	//run commit or rollback in last func
	defer func() { helper.CommitOrRollback(err, tx) }()

	//call func from repo
	todo, err := t.repo.FindByID(ctx, tx, Id)

	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (t *todoServiceImpl) FindByActivityID(ctx context.Context, Id int) ([]*model.Todo, error) {
	//begin db transaction
	tx, err := t.db.Begin()
	if err != nil {
		return nil, err
	}

	//run commit or rollback in last func
	defer func() { helper.CommitOrRollback(err, tx) }()

	//call func from repo
	todos, err := t.repo.FindByActivityID(ctx, tx, Id)

	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (t *todoServiceImpl) Update(ctx context.Context, req *model.Todo) (*model.Todo, error) {
	//begin db transaction
	tx, err := t.db.Begin()
	if err != nil {
		return nil, err
	}

	//run commit or rollback in last func
	defer func() { helper.CommitOrRollback(err, tx) }()

	//check todo id
	existing, err := t.repo.FindByID(ctx, tx, req.Todo_id)
	if err != nil {
		return nil, err
	}

	existing.Todo_id = req.Todo_id
	existing.Updated_at = time.Now().UTC()

	//validasi data req
	if req.Title != "" {
		existing.Title = req.Title
	}
	if req.Priority != "" {
		existing.Priority = req.Priority
	}
	//if req.Is_active != existing.Is_active {
	existing.Is_active = req.Is_active
	//}

	//call func repo to update
	update, err := t.repo.Update(ctx, tx, existing)
	if err != nil {
		return nil, err
	}

	return update, nil

}

func (t *todoServiceImpl) Delete(ctx context.Context, id int) error {
	//begin db transaction
	tx, err := t.db.Begin()
	if err != nil {
		return err
	}

	//run commit or rollback in last func
	defer func() { helper.CommitOrRollback(err, tx) }()

	//call repo
	err = t.repo.Delete(ctx, tx, id)

	return err
}
