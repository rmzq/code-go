package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var (
		dbUsername = "root"
		dbPassword = "asdf123"
		dbHost     = "localhost:3306"
		dbName     = "test_db"
	)
	connstStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true",
		dbUsername,
		dbPassword,
		dbHost,
		dbName,
	)
	connDB, err := sql.Open("mysql", connstStr)

	if err != nil {
		log.Fatalf("failed to connect server db, error: %s", err.Error())
	}

	connDB.SetConnMaxLifetime(time.Minute * 3)
	connDB.SetConnMaxIdleTime(time.Minute * 3)
	connDB.SetMaxOpenConns(10)
	connDB.SetMaxIdleConns(10)

	if err := connDB.Ping(); err != nil {
		log.Fatalf("failed to connect server db, error: %s", err.Error())
	}

	fmt.Println("mantap dh konek")
}
