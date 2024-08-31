package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // or _ "github.com/mattn/go-sqlite3" for SQLite
)

var DB *sql.DB

func Init() {
	var err error
	DB, err = sql.Open("postgres", "user=postgres password=m dbname=blog sslmode=require")
	if err != nil {
		log.Fatal(err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}

	createTable()
}

func createTable() {
	query := `
	CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
	
	CREATE TABLE IF NOT EXISTS posts(
		id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		title VARCHAR(255) NOT NULL,
		content TEXT NOT NULL,
		date VARCHAR(10) NOT NULL
	);
	`

	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
