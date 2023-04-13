package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/MCPutro/golang-todo/helper"
	"github.com/MCPutro/golang-todo/model"
	"log"
	"strings"
	"time"
)

type todoRepositoryImpl struct {
}

//var once sync.Once

func NewTodoRepository() TodoRepository {
	return &todoRepositoryImpl{}
}

func (t *todoRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, NewTodo *model.Todo) (*model.Todo, error) {
	SQL := "INSERT INTO todos (activity_group_id, title, priority, is_active, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?);"

	createdTime := NewTodo.Created_at.Format(helper.FORMAT_DATE)
	updatedTime := NewTodo.Updated_at.Format(helper.FORMAT_DATE)

	result, err := tx.ExecContext(ctx, SQL, NewTodo.Activity_group_id, NewTodo.Title, NewTodo.Priority, NewTodo.Is_active, createdTime, updatedTime)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	//update resp data
	NewTodo.Todo_id = int(id)
	createdTime2, _ := time.Parse(helper.FORMAT_DATE, createdTime)
	NewTodo.Created_at = createdTime2
	updatedTime2, _ := time.Parse(helper.FORMAT_DATE, updatedTime)
	NewTodo.Updated_at = updatedTime2

	return NewTodo, nil
}

func (t *todoRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]*model.Todo, error) {
	SQL := "select t.todo_id, t.activity_group_id, t.title, t.priority, t.is_active, t.created_at, t.updated_at from todos t ;"

	rows, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []*model.Todo
	for rows.Next() {
		var todo model.Todo
		if err := rows.Scan(&todo.Todo_id, &todo.Activity_group_id, &todo.Title, &todo.Priority, &todo.Is_active, &todo.Created_at, &todo.Updated_at); err != nil {
			return nil, err
		}
		todos = append(todos, &todo)
	}

	if len(todos) > 0 {
		return todos, nil
	} else {
		return nil, errors.New(helper.NO_DATA_FOUND)
	}
}

func (t *todoRepositoryImpl) FindByID(ctx context.Context, tx *sql.Tx, Id int) (*model.Todo, error) {
	SQL := "select t.todo_id, t.activity_group_id, t.title, t.priority, t.is_active, t.created_at, t.updated_at from todos t where t.todo_id = ?;"

	rows, err := tx.QueryContext(ctx, SQL, Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		var todo model.Todo
		if err := rows.Scan(&todo.Todo_id, &todo.Activity_group_id, &todo.Title, &todo.Priority, &todo.Is_active, &todo.Created_at, &todo.Updated_at); err != nil {
			if err == sql.ErrNoRows {
				return nil, errors.New(helper.NO_DATA_FOUND)
			} else {
				return nil, err
			}
		}
		return &todo, nil
	}
	return nil, errors.New(helper.NO_DATA_FOUND)
}

func (t *todoRepositoryImpl) FindByActivityID(ctx context.Context, tx *sql.Tx, Id int) ([]*model.Todo, error) {
	SQL := "select t.todo_id, t.activity_group_id, t.title, t.priority, t.is_active, t.created_at, t.updated_at from todos t where t.activity_group_id = ?;"

	rows, err := tx.QueryContext(ctx, SQL, Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []*model.Todo
	for rows.Next() {
		var todo model.Todo
		if err := rows.Scan(&todo.Todo_id, &todo.Activity_group_id, &todo.Title, &todo.Priority, &todo.Is_active, &todo.Created_at, &todo.Updated_at); err != nil {
			if err == sql.ErrNoRows {
				return nil, errors.New(helper.NO_DATA_FOUND)
			} else {
				return nil, err
			}
		}
		todos = append(todos, &todo)
	}

	if len(todos) > 0 {
		return todos, nil
	} else {
		return nil, errors.New(helper.NO_DATA_FOUND)
	}
}

func (t *todoRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, Todo *model.Todo) (*model.Todo, error) {

	SQL := "UPDATE todos t SET t.title = ?, t.priority  = ?, t.is_active = ? , t.updated_at = ? WHERE t.todo_id = ?;"

	updatedTime := Todo.Updated_at.Format(helper.FORMAT_DATE)

	_, err := tx.ExecContext(ctx, SQL, Todo.Title, Todo.Priority, Todo.Is_active, updatedTime, Todo.Todo_id)

	if err != nil {
		return nil, err
	}

	//if rowsAffected, err := result.RowsAffected(); err != nil {
	//	return err
	//} else if rowsAffected == 0 {
	//	return errors.New(helper.NO_DATA_FOUND)
	//}
	updatedTime2, _ := time.Parse(helper.FORMAT_DATE, updatedTime)
	Todo.Updated_at = updatedTime2

	return Todo, nil
}

func (t *todoRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, Id int) error {
	SQL := "DELETE FROM todos WHERE todo_id = ? ;"

	result, err := tx.ExecContext(ctx, SQL, Id)
	if err != nil {
		return err
	}

	if rowsAffected, err := result.RowsAffected(); err != nil {
		return err
	} else if rowsAffected == 0 {
		return errors.New(helper.NO_DATA_FOUND)
	}

	return nil
}

func (t *todoRepositoryImpl) GetPriorityList(ctx context.Context, tx *sql.Tx) (map[string]bool, error) {
	SQL := "select p.priority_name from priority p where p.is_active = 1 order by priority_id asc ;"

	rows, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	priorityList := make(map[string]bool, 0)
	for rows.Next() {
		var priority string
		if err := rows.Scan(&priority); err != nil {
			return nil, err
		}
		priorityList[strings.ToUpper(priority)] = true

		//set default priority
		if helper.DEFAULT_PRIORITY == "" {
			helper.DEFAULT_PRIORITY = priority
			log.Println("[INFO] set var DEFAULT_PRIORITY to ", helper.DEFAULT_PRIORITY)
		}
	}

	return priorityList, nil
}
