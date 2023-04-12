package service

import (
	"context"
	"database/sql"
	"github.com/MCPutro/golang-todo/helper"
	"github.com/MCPutro/golang-todo/model"
	"github.com/MCPutro/golang-todo/repository"
	"time"
)

type activityServiceImpl struct {
	repo repository.ActivityRepository
	db   *sql.DB
}

func NewActivityService(repo repository.ActivityRepository, db *sql.DB) ActivityService {
	return &activityServiceImpl{repo: repo, db: db}
}

func (a *activityServiceImpl) Create(ctx context.Context, req *model.Activities) (*model.Activities, error) {
	//set created and update time to now
	if req.Created_at.IsZero() {
		req.Created_at = time.Now()
	}
	if req.Updated_at.IsZero() {
		req.Updated_at = time.Now()
	}

	//begin db transaction
	tx, err := a.db.Begin()
	if err != nil {
		return nil, err
	}
	//close db transaction
	defer func() {
		helper.CommitOrRollback(err, tx)
	}()

	//call func save from repo
	activitySave, err := a.repo.Save(ctx, tx, req)
	if err != nil {
		return nil, err
	}

	return activitySave, nil
}

func (a *activityServiceImpl) Update(ctx context.Context, req *model.Activities) (*model.Activities, error) {
	//begin db transaction
	tx, err := a.db.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		helper.CommitOrRollback(err, tx)
	}()

	//check activity id is exists or not
	existing, err := a.repo.FindByID(ctx, tx, req.Activity_id)
	if err != nil {
		return nil, err
	}

	//if activity id is exists
	existing.Title = req.Title
	existing.Updated_at = time.Now()

	//call repo to update
	err = a.repo.Update(ctx, tx, existing)
	if err != nil {
		return nil, err
	}

	req = existing

	return req, nil

}

func (a *activityServiceImpl) FindAll(ctx context.Context) ([]*model.Activities, error) {
	//begin db transaction
	tx, err := a.db.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		helper.CommitOrRollback(err, tx)
	}()

	//call repo
	activities, err := a.repo.FindAll(ctx, tx)
	if err != nil {
		return nil, err
	}

	return activities, nil
}

func (a *activityServiceImpl) FindById(ctx context.Context, Id int) (*model.Activities, error) {
	//begin db transaction
	tx, err := a.db.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		helper.CommitOrRollback(err, tx)
	}()

	//call repo
	activities, err := a.repo.FindByID(ctx, tx, Id)
	if err != nil {
		return nil, err
	}

	return activities, nil
}

func (a *activityServiceImpl) Delete(ctx context.Context, id int) error {
	//begin db transaction
	tx, err := a.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		helper.CommitOrRollback(err, tx)
	}()

	err = a.repo.Delete(ctx, tx, id)
	if err != nil {
		return err
	}

	return nil
}
