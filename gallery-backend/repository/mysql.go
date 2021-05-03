package repository

import (
	"database/sql"
	. "gallery-backend/utils"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var galleryDB *sql.DB

func init() {
	var err error
	galleryDB, err = sql.Open("mysql", "root:root@tcp(mysql)/gallery")

	if err != nil {
		ErrorLog.Printf("Connect to db error: %v\n", err)
		panic(err)
	}

	galleryDB.SetConnMaxLifetime(time.Minute * 3)
	galleryDB.SetMaxOpenConns(10)
	galleryDB.SetMaxIdleConns(10)
}
