package database

import (
	"context"
	"database/sql"
	"log"
)

func PrepareDB(db *sql.DB) error {
	SQL1 := `create table if not exists activities (
		activity_id int auto_increment 					 primary key,
		title       varchar(100)                         not null,
		email       varchar(100)                         not null,
		created_at  datetime default current_timestamp() not null,
		updated_at  datetime default current_timestamp() not null
	);`

	SQL2 := `create table if not exists priority (
			priority_id   int auto_increment   primary key,
			priority_name varchar(50)          not null,
			is_active     tinyint(1) default 1 not null,
			constraint priority_name
				unique (priority_name)
	);`

	SQL3 := `create table if not exists todos (
		todo_id           int auto_increment         		   primary key,
		activity_group_id int                                  not null,
		title             varchar(200)                         null,
		priority          varchar(50)                          not null,
		is_active         tinyint(1)                           null,
		created_at        datetime default current_timestamp() null,
		updated_at        datetime default current_timestamp() null,
		constraint todos_activities_activity_id_fk
			foreign key (activity_group_id) references activities (activity_id),
		constraint todos_priority_priority_name_fk
			foreign key (priority) references priority (priority_name)
	);`

	SQL4 := `insert IGNORE  into priority (priority_name, is_active)
				values  ('very-high', 1), ('high', 1), ('medium', 1), ('low', 1),('very-low', 1);` //# ON DUPLICATE KEY UPDATE is_active = 1

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

	//SQL2
	_, err = tx.ExecContext(context.Background(), SQL2)
	if err != nil {
		return err
	}

	//SQL3
	_, err = tx.ExecContext(context.Background(), SQL3)
	if err != nil {
		return err
	}

	//SQL4
	_, err = tx.ExecContext(context.Background(), SQL4)
	if err != nil {
		return err
	}

	return err
}
