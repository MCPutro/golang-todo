package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/MCPutro/golang-todo/helper"
	"github.com/MCPutro/golang-todo/model"
)

type activityRepositoryImpl struct {
}

func NewActivityRepository() ActivityRepository {
	return &activityRepositoryImpl{}
}

func (a *activityRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, NewActivity *model.Activities) (*model.Activities, error) {

	SQL := "INSERT INTO activities (title, email, created_at, updated_at) VALUES (?, ?, ?, ?);"

	result, err := tx.ExecContext(ctx, SQL,
		NewActivity.Title, NewActivity.Email, NewActivity.Created_at.Format(helper.FORMAT_DATE), NewActivity.Updated_at.Format(helper.FORMAT_DATE))
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	NewActivity.Activity_id = int(id)

	return NewActivity, nil
}

func (a *activityRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]*model.Activities, error) {
	SQL := "select a.activity_id, a.title, a.email, a.created_at, a.updated_at from activities a"
	rows, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var activities []*model.Activities
	for rows.Next() {
		var act model.Activities
		if err := rows.Scan(&act.Activity_id, &act.Title, &act.Email, &act.Created_at, &act.Updated_at); err != nil {
			return nil, err
		}
		activities = append(activities, &act)
	}

	return activities, nil
}

func (a *activityRepositoryImpl) FindByID(ctx context.Context, tx *sql.Tx, Id int) (*model.Activities, error) {
	SQL := "select a.activity_id, a.title, a.email, a.created_at, a.updated_at from activities a where a.activity_id = ?"
	rows, err := tx.QueryContext(ctx, SQL, Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		var act model.Activities
		if err := rows.Scan(&act.Activity_id, &act.Title, &act.Email, &act.Created_at, &act.Updated_at); err != nil {
			if err == sql.ErrNoRows {
				return nil, errors.New("no data found")
			} else {
				return nil, err
			}
		}
		return &act, nil
	}

	return nil, errors.New("no data found")
}

func (a *activityRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, Activity *model.Activities) error {
	SQL := "UPDATE activities t SET t.title = ?, t.updated_at = ? WHERE t.activity_id = ?;"

	result, err := tx.ExecContext(ctx, SQL, Activity.Title, Activity.Updated_at.Format(helper.FORMAT_DATE), Activity.Activity_id)
	if err != nil {
		return err
	}

	if rowsAffected, err := result.RowsAffected(); err != nil {
		return err
	} else if rowsAffected == 0 {
		return errors.New("no data updated")
	}

	return nil
}

func (a *activityRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, Id int) error {
	SQL := "DELETE FROM activities WHERE activity_id = ?"

	result, err := tx.ExecContext(ctx, SQL, Id)
	if err != nil {
		return err
	}

	if rowsAffected, err := result.RowsAffected(); err != nil {
		return err
	} else if rowsAffected == 0 {
		return errors.New("no data deleted")
	}

	return nil
}
