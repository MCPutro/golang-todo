package repository

import (
	"context"
	"database/sql"
	"github.com/MCPutro/golang-todo/model"
)

type ActivityRepository interface {
	//untuk menyimpan data Activity baru
	Save(ctx context.Context, tx *sql.Tx, NewActivity *model.Activities) (*model.Activities, error)

	//menampilan seluruh data Activity yang sudah terdaftar
	FindAll(ctx context.Context, tx *sql.Tx) ([]*model.Activities, error)

	//menampilkan Activity berdasarkan Activity ID
	FindByID(ctx context.Context, tx *sql.Tx, Id int) (*model.Activities, error)

	//memperbaharui data Activity
	Update(ctx context.Context, tx *sql.Tx, Activity *model.Activities) error

	//menghapus data Activity berdasarkan Activity Id
	Delete(ctx context.Context, tx *sql.Tx, Id int) error
}
