package repository

import (
	"context"
	"database/sql"
	"github.com/MCPutro/golang-todo/model"
)

type ActivityRepository interface {
	//untuk menyimpan data Activity baru
	Save(ctx context.Context, tx *sql.Tx, NewActivity *model.Activity) (*model.Activity, error)

	//menampilan seluruh data Activity yang sudah terdaftar
	FindAll(ctx context.Context, tx *sql.Tx) ([]*model.Activity, error)

	//menampilkan Activity berdasarkan Activity ID
	FindByID(ctx context.Context, tx *sql.Tx, Id int) (*model.Activity, error)

	//memperbaharui data Activity
	Update(ctx context.Context, tx *sql.Tx, Activity *model.Activity) (*model.Activity, error)

	//menghapus data Activity berdasarkan Activity Id
	Delete(ctx context.Context, tx *sql.Tx, Id int) error
}
