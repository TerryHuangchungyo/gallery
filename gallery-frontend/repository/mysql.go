package repository

import (
	"database/sql"
	"fmt"
	"gallery-frontend/config"
	. "gallery-frontend/utils"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var galleryDB *sql.DB

func init() {
	var err error

	dbType := config.GalleryDB.DatabaseType
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		config.GalleryDB.User,
		config.GalleryDB.Password,
		config.GalleryDB.Host,
		config.GalleryDB.DatabaseName,
	)

	galleryDB, err = sql.Open(dbType, dsn)

	if err != nil {
		ErrorLog.Printf("Connect to db error: %v\n", err)
		panic(err)
	}

	galleryDB.SetConnMaxLifetime(time.Minute * 3)
	galleryDB.SetMaxOpenConns(10)
	galleryDB.SetMaxIdleConns(10)
}
