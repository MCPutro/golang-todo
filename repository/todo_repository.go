package repository

import (
	"context"
	"database/sql"
	"github.com/MCPutro/golang-todo/model"
)

type TodoRepository interface {
	//untuk menyimpan data Todo baru
	Save(ctx context.Context, tx *sql.Tx, NewTodo *model.Todos) (*model.Todos, error)

	//menampilan seluruh data Todo yang ada
	FindAll(ctx context.Context, tx *sql.Tx) ([]*model.Todos, error)

	//menampilkan Activity berdasarkan Todo ID
	FindByID(ctx context.Context, tx *sql.Tx, Id int) (*model.Todos, error)

	//memperbaharui data Todo
	Update(ctx context.Context, tx *sql.Tx, Activity *model.Todos) error

	//menghapus data Todo berdasarkan Todo Id
	Delete(ctx context.Context, tx *sql.Tx, Id int) error
}
