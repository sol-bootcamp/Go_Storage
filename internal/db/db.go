package db

import (
	"clases/internal/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB(cfg *config.Config) *sql.DB {
	connStr := fmt.Sprintf("%s@tcp(%s:%s)/%s",
		cfg.DBUser, cfg.DBHost, cfg.DBPort, cfg.DBName)
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database!")
	return db
}
