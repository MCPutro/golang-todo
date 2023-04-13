package repository

import (
	"context"
	"database/sql"
	"github.com/MCPutro/golang-todo/model"
)

type TodoRepository interface {
	// Save untuk menyimpan data Todo baru
	Save(ctx context.Context, tx *sql.Tx, NewTodo *model.Todo) (*model.Todo, error)

	// FindAll menampilan seluruh data Todo yang ada
	FindAll(ctx context.Context, tx *sql.Tx) ([]*model.Todo, error)

	// FindByID menampilkan Activity berdasarkan Todo ID
	FindByID(ctx context.Context, tx *sql.Tx, Id int) (*model.Todo, error)

	// FindByActivityID menampilkan Activity berdasarkan Activity ID
	FindByActivityID(ctx context.Context, tx *sql.Tx, Id int) ([]*model.Todo, error)

	// Update memperbaharui data Todo
	Update(ctx context.Context, tx *sql.Tx, Todo *model.Todo) (*model.Todo, error)

	// Delete menghapus data Todo berdasarkan Todo Id
	Delete(ctx context.Context, tx *sql.Tx, Id int) error

	// GetPriorityList get list priority
	GetPriorityList(ctx context.Context, tx *sql.Tx) (map[string]bool, error)
}
