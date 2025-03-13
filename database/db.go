package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() *sql.DB {
	dsn := "root:yourpassword@tcp(127.0.0.1:3306)/task_manager?parseTime=true"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Error pinging database:", err)
	}

	fmt.Println("Connected to database successfully!")
	DB = db
	return db
}
