package database

import (
	"database/sql"
	"fmt"
	"github.com/MCPutro/golang-todo/config"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func GetInstance() (*sql.DB, error) {
	var db *sql.DB
	var err error
	i := 1
	for ; i <= 5; i++ {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
			config.DbUser, config.DbPass,
			config.DbHost, config.DbPort, config.DbName)
		db, err = sql.Open("mysql", dsn)

		if err = db.Ping(); err != nil {
			log.Printf("error create db connection [rety %d times], message : %s", i, err)
			if i == 5 {
				return nil, err
			}

		} else {
			//set open connection count and time
			db.SetMaxIdleConns(5)
			db.SetMaxOpenConns(100)
			db.SetConnMaxLifetime(60 * time.Minute)
			db.SetConnMaxIdleTime(10 * time.Minute)

			//out from looping
			break
		}
	}

	return db, err
}
