package database

import (
	"context"
	"database/sql"
	"log"
)

func PrepareDB(db *sql.DB) error {
	SQL1 := `create table if not exists activities (
		activity_id bigint auto_increment 					 primary key,
		title       varchar(100)                         null,
		email       varchar(100)                         null,
		created_at  datetime(3)  null,
		updated_at  datetime(3)  null
	);`

	SQL2 := `create table if not exists todos (
		todo_id           bigint auto_increment         	   primary key,
		activity_group_id bigint                               null,
		title             varchar(200)                         null,
		priority          varchar(50)                          null,
		is_active         tinyint(1)                           null,
		created_at        datetime(3)  null,
		updated_at        datetime(3)  null
	);`

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			log.Println("[INFO] Success init Table and Data")
			tx.Commit()
		}
	}()

	//SQL1
	_, err = tx.ExecContext(context.Background(), SQL1)
	if err != nil {
		return err
	}

	//SQL3
	_, err = tx.ExecContext(context.Background(), SQL2)
	if err != nil {
		return err
	}

	return err
}
