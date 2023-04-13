package repository

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/MCPutro/golang-todo/helper"
	"github.com/MCPutro/golang-todo/model"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetAllActivity(t *testing.T) {

	ctx := context.Background()

	repository := NewActivityRepository()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock db: %v", err)
	}
	defer db.Close()

	activities := []model.Activity{
		{Activity_id: 1, Title: "1", Email: "1", Created_at: time.Now(), Updated_at: time.Now()},
		{Activity_id: 2, Title: "2", Email: "2", Created_at: time.Now(), Updated_at: time.Now()},
		{Activity_id: 3, Title: "3", Email: "3", Created_at: time.Now(), Updated_at: time.Now()},
	}

	//set expect data
	rows := sqlmock.NewRows([]string{"Activity_id", "Title", "Email", "Created_at", "Updated_at"})
	for _, act := range activities {
		rows.AddRow(act.Activity_id, act.Title, act.Email, act.Created_at, act.Updated_at)
	}

	mock.ExpectBegin()

	mock.ExpectQuery(`select a.activity_id, a.title, a.email, a.created_at, a.updated_at from activities a order by a.activity_id desc ;`).WillReturnRows(rows)

	mock.ExpectCommit()

	//membuat Database transaction
	tx, err := db.Begin()

	//call save func from UserRepository
	findAll, err := repository.FindAll(ctx, tx)

	if err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.NoError(t, err)
	assert.Len(t, findAll, 3, "invalid length data")
}

func TestGetOneActivity(t *testing.T) {

	ctx := context.Background()

	repository := NewActivityRepository()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock db: %v", err)
	}
	defer db.Close()

	activities := []model.Activity{
		{Activity_id: 1, Title: "sekolah", Email: "1", Created_at: time.Now(), Updated_at: time.Now()},
	}

	//set expect data
	rows := sqlmock.NewRows([]string{"Activity_id", "Title", "Email", "Created_at", "Updated_at"})
	for _, act := range activities {
		rows.AddRow(act.Activity_id, act.Title, act.Email, act.Created_at, act.Updated_at)
	}

	mock.ExpectBegin()

	mock.ExpectQuery(`select a.activity_id, a.title, a.email, a.created_at, a.updated_at from activities a where a.activity_id = ? `).WillReturnRows(rows)

	mock.ExpectCommit()

	//membuat Database transaction
	tx, err := db.Begin()

	//call save func from UserRepository
	findAll, err := repository.FindByID(ctx, tx, 1)

	if err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.NoError(t, err)
	assert.Equal(t, findAll.Activity_id, 1)
	assert.Equal(t, findAll.Title, "sekolah")
}

func TestSaveActivity(t *testing.T) {

	ctx := context.Background()

	repository := NewActivityRepository()

	newAct := model.Activity{
		Title:      "1",
		Email:      "1",
		Created_at: time.Now(),
		Updated_at: time.Now(),
	}

	createdTime := newAct.Created_at.Format(helper.FORMAT_DATE)
	updatedTime := newAct.Updated_at.Format(helper.FORMAT_DATE)

	//create database mock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock db: %v", err)
	}
	defer db.Close()

	mock.ExpectBegin()

	// expect insert query
	mock.ExpectExec(`INSERT INTO activities `).
		WithArgs(newAct.Title, newAct.Email, createdTime, updatedTime).
		WillReturnResult(sqlmock.NewResult(2, 1))

	mock.ExpectCommit()

	//membuat Database transaction
	tx, err := db.Begin()

	//call save func from UserRepository
	save, err1 := repository.Save(ctx, tx, &newAct)

	if err1 != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Nil(t, err1)
	assert.Equal(t, save.Activity_id, 2)
}
